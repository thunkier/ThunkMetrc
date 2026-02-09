using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using ThunkMetrc.Client;
using Microsoft.Extensions.Logging;
using System.Text.Json;
using System.Threading;
using System.Runtime.CompilerServices;
using ThunkMetrc.Wrapper;
using ThunkMetrc.Wrapper.Models;
using File = ThunkMetrc.Wrapper.Models.File;

namespace ThunkMetrc.Wrapper.Services
{
    public class WebhooksService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// 
        /// </summary>
        /// <param name="subscriptionId">subscriptionId</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteWebhooksBySubscriptionId(string subscriptionId, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteWebhooksBySubscriptionId(subscriptionId); return null; }, cancellationToken);
        }
        /// <summary>
        /// 
        /// </summary>
        
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetWebhooks(CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetWebhooks(), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="subscriptionId">subscriptionId</param>
        /// <param name="body">Request body.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateWebhooksDisableBySubscriptionId(string subscriptionId, object? body = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateWebhooksDisableBySubscriptionId(subscriptionId, body); return null; }, cancellationToken);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="subscriptionId">subscriptionId</param>
        /// <param name="body">Request body.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateWebhooksEnableBySubscriptionId(string subscriptionId, object? body = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateWebhooksEnableBySubscriptionId(subscriptionId, body); return null; }, cancellationToken);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateWebhooks(WebhooksUpdateWebhooksRequest? body = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateWebhooks(body); return null; }, cancellationToken);
        }
    }
}

