using System;
using System.Net.Http;
using Microsoft.Extensions.Logging;
using ThunkMetrc.Client;

namespace ThunkMetrc.Wrapper
{
    /// <summary>
    /// A Factory for creating MetrcWrappers that share underlying resources
    /// (RateLimiters and HttpClients) to ensure safe high-throughput usage
    /// across multiple licenses or facilities.
    /// </summary>
    public class MetrcFactory : IDisposable
    {
        private readonly MetrcRateLimiter _sharedLimiter;
        private readonly HttpClient _sharedHttpClient;
        private readonly bool _disposeClient;
        private readonly ILogger? _logger;

        /// <summary>
        /// Initializes a new instance of the MetrcFactory.
        /// </summary>
        /// <param name="config">Optional rate limiter configuration.</param>
        /// <param name="logger">Optional logger.</param>
        /// <param name="httpClient">
        /// Optional shared HttpClient. If provided, the factory will use this client for all wrappers.
        /// If null, the factory will create and manage its own HttpClient.
        /// </param>
        public MetrcFactory(RateLimiterConfig? config = null, ILogger? logger = null, HttpClient? httpClient = null)
        {
            _logger = logger;
            _sharedLimiter = new MetrcRateLimiter(config, logger);
            if (httpClient != null)
            {
                _sharedHttpClient = httpClient;
                _disposeClient = false;
            }
            else
            {
                _sharedHttpClient = new HttpClient();
                _disposeClient = true;
            }
        }

        /// <summary>
        /// Creates a new MetrcWrapper configured for a specific facility/user.
        /// The wrapper will use the factory's shared RateLimiter and HttpClient.
        /// </summary>
        /// <param name="baseUrl">The base URL (e.g. https://api-ca.metrc.com)</param>
        /// <param name="vendorKey">The Vendor Key</param>
        /// <param name="userKey">The User Key</param>
        /// <returns>A configured MetrcWrapper instance.</returns>
        public MetrcWrapper Create(string baseUrl, string vendorKey, string userKey)
        {
            var client = new MetrcClient(baseUrl, vendorKey, userKey, _sharedHttpClient, _logger);
            return new MetrcWrapper(client, _sharedLimiter);
        }

        public void Dispose()
        {
            if (_disposeClient)
            {
                _sharedHttpClient.Dispose();
            }
            GC.SuppressFinalize(this);
        }
    }
}
