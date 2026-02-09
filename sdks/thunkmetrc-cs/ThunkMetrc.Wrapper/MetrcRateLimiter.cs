using System;
using System.Collections.Concurrent;
using System.Net.Http;
using System.Threading;
using System.Threading.RateLimiting;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;
using Polly;
using Polly.Retry;
using Polly.CircuitBreaker;
using System.Diagnostics.Metrics;
using ThunkMetrc.Client;

namespace ThunkMetrc.Wrapper
{
    public class RateLimiterConfig
    {
        public bool Enabled { get; set; } = false;
        public int MaxGetPerSecondPerFacility { get; set; } = 50;
        public int MaxGetPerSecondIntegrator { get; set; } = 150;
        public int MaxConcurrentGetPerFacility { get; set; } = 10;
        public int MaxConcurrentGetIntegrator { get; set; } = 30;

        // Resilience Settings
        public int? MaxRetries { get; set; } = null; // null = unlimited
        public TimeSpan InitialBackoff { get; set; } = TimeSpan.FromSeconds(1);
        public TimeSpan MaxBackoff { get; set; } = TimeSpan.FromSeconds(60);
        public double BackoffExponent { get; set; } = 2.0;

        // Circuit Breaker
        public double CircuitBreakerFailureRatio { get; set; } = 0.5;
        public int CircuitBreakerMinimumThroughput { get; set; } = 10;
        public TimeSpan CircuitBreakerDuration { get; set; } = TimeSpan.FromSeconds(30);
    }

    public class MetrcRateLimiter : IMetrcRateLimiter
    {
        private readonly RateLimiterConfig _config;
        private readonly ILogger? _logger;
        private readonly ResiliencePipeline _pipeline;
        
        // Metrics
        private static readonly Meter _meter = new("ThunkMetrc.Wrapper", "1.0.0");
        private static readonly Counter<long> _requestCounter = _meter.CreateCounter<long>("metrc.wrapper.requests");
        private static readonly Counter<long> _retryCounter = _meter.CreateCounter<long>("metrc.wrapper.retries");
        private static readonly Counter<long> _rateLimitCounter = _meter.CreateCounter<long>("metrc.wrapper.ratelimits");

        // Integrator Limits
        private readonly TokenBucketRateLimiter _integratorRateLimiter;
        private readonly ConcurrencyLimiter _integratorConcurrencyLimiter;

        // Facility Limits
        private readonly ConcurrentDictionary<string, TokenBucketRateLimiter> _facilityRateLimiters = new();
        private readonly ConcurrentDictionary<string, ConcurrencyLimiter> _facilityConcurrencyLimiters = new();

        public MetrcRateLimiter(RateLimiterConfig? config = null, ILogger? logger = null)
        {
            _logger = logger;
            _config = config ?? new RateLimiterConfig();

            // Initialize Resilience Pipeline
            _pipeline = new ResiliencePipelineBuilder()
                .AddRetry(new RetryStrategyOptions
                {
                    ShouldHandle = new PredicateBuilder().Handle<Exception>(ex => 
                    {
                        var isMetrc429 = ex is MetrcException mex429 && (int)mex429.StatusCode == 429;
                        var isHttp429 = ex is HttpRequestException hre429 && hre429.StatusCode == System.Net.HttpStatusCode.TooManyRequests;
                        var isMetrc5xx = ex is MetrcException mex5xx && (int)mex5xx.StatusCode >= 500;
                        var isHttp5xx = ex is HttpRequestException hre5xx && (int?)hre5xx.StatusCode >= 500;
                        return isMetrc429 || isHttp429 || isMetrc5xx || isHttp5xx;
                    }),
                    MaxRetryAttempts = _config.MaxRetries ?? int.MaxValue,
                    DelayGenerator = args =>
                    {
                        _retryCounter.Add(1);
                        var ex = args.Outcome.Exception;
                        
                        // Check for Retry-After
                        if (ex is MetrcException mex && (int)mex.StatusCode == 429)
                        {
                            if (mex.Headers?.RetryAfter?.Delta != null)
                                return new ValueTask<TimeSpan?>(mex.Headers.RetryAfter.Delta.Value);
                            if (mex.Headers?.RetryAfter?.Date != null)
                                return new ValueTask<TimeSpan?>(mex.Headers.RetryAfter.Date.Value - DateTimeOffset.Now);
                        }

                        // Exponential Backoff with Jitter
                        double exponentialMs = _config.InitialBackoff.TotalMilliseconds * Math.Pow(_config.BackoffExponent, args.AttemptNumber);
                        double cappedMs = Math.Min(exponentialMs, _config.MaxBackoff.TotalMilliseconds);
                        // Add up to 20% Jitter
                        double jitter = cappedMs * (Random.Shared.NextDouble() * 0.2); 
                        
                        return new ValueTask<TimeSpan?>(TimeSpan.FromMilliseconds(cappedMs + jitter));
                    },
                    OnRetry = args =>
                    {
                        if (args.Outcome.Exception is Exception ex)
                        {
                            _logger?.LogWarning(ex, "Retry {RetryCount} after {Delay}ms due to {Message}", 
                                args.AttemptNumber, args.RetryDelay.TotalMilliseconds, ex.Message);
                        }
                        return default;
                    }
                })
                .AddCircuitBreaker(new CircuitBreakerStrategyOptions
                {
                    FailureRatio = _config.CircuitBreakerFailureRatio,
                    SamplingDuration = TimeSpan.FromSeconds(30),
                    MinimumThroughput = _config.CircuitBreakerMinimumThroughput,
                    BreakDuration = _config.CircuitBreakerDuration,
                    ShouldHandle = new PredicateBuilder().Handle<Exception>(ex => 
                    {
                        var isMetrc5xx = ex is MetrcException mex5xx && (int)mex5xx.StatusCode >= 500;
                        var isHttp5xx = ex is HttpRequestException hre5xx && (int?)hre5xx.StatusCode >= 500;
                        return isMetrc5xx || isHttp5xx;
                    }),
                    OnOpened = args =>
                    {
                        _logger?.LogError("Circuit Breaker OPENED for {Duration}s due to high failure rate.", args.BreakDuration.TotalSeconds);
                        return default;
                    },
                    OnClosed = args =>
                    {
                        _logger?.LogInformation("Circuit Breaker CLOSED. Resuming normal operation.");
                        return default;
                    }
                })
                .Build();


            // Setup global integrator limits
            _integratorRateLimiter = new TokenBucketRateLimiter(new TokenBucketRateLimiterOptions
            {
                TokenLimit = _config.MaxGetPerSecondIntegrator,
                QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                QueueLimit = int.MaxValue,
                ReplenishmentPeriod = TimeSpan.FromSeconds(1),
                TokensPerPeriod = _config.MaxGetPerSecondIntegrator,
                AutoReplenishment = true
            });

            _integratorConcurrencyLimiter = new ConcurrencyLimiter(new ConcurrencyLimiterOptions
            {
                PermitLimit = _config.MaxConcurrentGetIntegrator,
                QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                QueueLimit = int.MaxValue
            });
        }

        public async Task<T> ExecuteAsync<T>(string? facility, bool isGet, Func<Task<T>> op, CancellationToken cancellationToken = default)
        {
            if (!_config.Enabled || !isGet)
            {
                return await _pipeline.ExecuteAsync(async ct => await op(), cancellationToken);
            }

            // Acquire Integrator Concurrency
            using var integratorConcurrencyLease = await _integratorConcurrencyLimiter.AcquireAsync(1, cancellationToken);
            if (!integratorConcurrencyLease.IsAcquired) throw new InvalidOperationException("Failed to acquire integrator concurrency permit");

            // Acquire Facility Concurrency
            RateLimitLease? facilityConcurrencyLease = null;
            if (!string.IsNullOrEmpty(facility))
            {
                var fl = GetFacilityConcurrency(facility);
                facilityConcurrencyLease = await fl.AcquireAsync(1, cancellationToken);
                if (!facilityConcurrencyLease.IsAcquired)
                     throw new InvalidOperationException($"Failed to acquire facility concurrency permit for {facility}");
            }

            try
            {
                // Acquire Integrator Rate (Wait if needed)
                using var integratorRateLease = await _integratorRateLimiter.AcquireAsync(1, cancellationToken);
                if (!integratorRateLease.IsAcquired) 
                {
                     _rateLimitCounter.Add(1);
                     await Task.Delay(100, cancellationToken); 
                }

                // Acquire Facility Rate
                if (!string.IsNullOrEmpty(facility))
                {
                    var fl = GetFacilityRate(facility);
                    using var facilityRateLease = await fl.AcquireAsync(1, cancellationToken);
                    if (!facilityRateLease.IsAcquired)
                    {
                         _rateLimitCounter.Add(1);
                    }
                }

                _requestCounter.Add(1);
                return await _pipeline.ExecuteAsync(async ct => await op(), cancellationToken);
            }
            finally
            {
                facilityConcurrencyLease?.Dispose();
            }
        }
        
        private TokenBucketRateLimiter GetFacilityRate(string facility)
        {
            return _facilityRateLimiters.GetOrAdd(facility, _ => new TokenBucketRateLimiter(new TokenBucketRateLimiterOptions
            {
                TokenLimit = _config.MaxGetPerSecondPerFacility,
                QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                QueueLimit = int.MaxValue,
                ReplenishmentPeriod = TimeSpan.FromSeconds(1),
                TokensPerPeriod = _config.MaxGetPerSecondPerFacility,
                AutoReplenishment = true
            }));
        }

        private ConcurrencyLimiter GetFacilityConcurrency(string facility)
        {
             return _facilityConcurrencyLimiters.GetOrAdd(facility, _ => new ConcurrencyLimiter(new ConcurrencyLimiterOptions
            {
                PermitLimit = _config.MaxConcurrentGetPerFacility,
                QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                QueueLimit = int.MaxValue
            }));
        }
    }
}
