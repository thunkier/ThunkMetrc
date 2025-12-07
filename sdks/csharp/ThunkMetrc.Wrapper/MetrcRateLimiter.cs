using System;
using System.Collections.Concurrent;
using System.Net.Http;
using System.Threading;
using System.Threading.RateLimiting;
using System.Threading.Tasks;

namespace ThunkMetrc.Wrapper
{
    public class RateLimiterConfig
    {
        public bool Enabled { get; set; } = false;
        public int MaxGetPerSecondPerFacility { get; set; } = 50;
        public int MaxGetPerSecondIntegrator { get; set; } = 150;
        public int MaxConcurrentGetPerFacility { get; set; } = 10;
        public int MaxConcurrentGetIntegrator { get; set; } = 30;
    }

    public class MetrcRateLimiter
    {
        private readonly RateLimiterConfig _config;
        
        // Integrator Limits
        private readonly TokenBucketRateLimiter _integratorRateLimiter;
        private readonly ConcurrencyLimiter _integratorConcurrencyLimiter;

        // Facility Limits
        private readonly ConcurrentDictionary<string, TokenBucketRateLimiter> _facilityRateLimiters = new();
        private readonly ConcurrentDictionary<string, ConcurrencyLimiter> _facilityConcurrencyLimiters = new();

        public MetrcRateLimiter(RateLimiterConfig? config = null)
        {
            _config = config ?? new RateLimiterConfig();

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

        public async Task<T> ExecuteAsync<T>(string? facility, bool isGet, Func<Task<T>> op)
        {
            if (!_config.Enabled || !isGet)
            {
                return await op();
            }

            // 1. Acquire Integrator Concurrency
            using var integratorConcurrencyLease = await _integratorConcurrencyLimiter.AcquireAsync(1);
            if (!integratorConcurrencyLease.IsAcquired) throw new InvalidOperationException("Failed to acquire integrator concurrency permit");

            // 2. Acquire Facility Concurrency
            IDisposable? facilityConcurrencyLease = null;
            if (!string.IsNullOrEmpty(facility))
            {
                var fl = GetFacilityConcurrency(facility);
                facilityConcurrencyLease = await fl.AcquireAsync(1);
                if (facilityConcurrencyLease is RateLimitLease lease && !lease.IsAcquired)
                     throw new InvalidOperationException($"Failed to acquire facility concurrency permit for {facility}");
            }

            try
            {
                // 3. Acquire Integrator Rate
                using var integratorRateLease = await _integratorRateLimiter.AcquireAsync(1);
                if (!integratorRateLease.IsAcquired) 
                {
                     // Should wait in queue, but if rejected:
                     await Task.Delay(100); // Simple fallback check
                }

                // 4. Acquire Facility Rate
                if (!string.IsNullOrEmpty(facility))
                {
                    var fl = GetFacilityRate(facility);
                    using var facilityRateLease = await fl.AcquireAsync(1);
                    if (!facilityRateLease.IsAcquired)
                    {
                         // queue logic handled by limiter usually
                    }
                }

                // 5. Retry Loop for 429
                while (true)
                {
                    try
                    {
                        return await op();
                    }
                    catch (Exception ex)
                    {
                        // Check for 429
                        // The generated client currently throws generic exceptions or returns object
                        // We need to check exception message if Client isn't strongly typed error
                        // Assuming Client throws HttpRequestException or similar with status code info
                        if (ex.Message.Contains("429") || (ex is HttpRequestException hre && hre.StatusCode == System.Net.HttpStatusCode.TooManyRequests))
                        {
                            // Backoff
                            await Task.Delay(1000);
                            continue;
                        }
                        throw;
                    }
                }
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
