using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// 
        /// </summary>
        /// <param name="subscriptionId">subscriptionId</param>
        public async Task DeleteWebhooksBySubscriptionId(string subscriptionId)
        {
            var queryParams = new Dictionary<string, string>();
            await SendAsync(new HttpMethod("DELETE"), $"/webhooks/v2/{Uri.EscapeDataString(subscriptionId)}", null, queryParams);
        }
        /// <summary>
        /// 
        /// </summary>
        public async Task<object> GetWebhooks()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/webhooks/v2", null, queryParams);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="subscriptionId">subscriptionId</param>
        /// <param name="body">Request body.</param>
        public async Task UpdateWebhooksDisableBySubscriptionId(string subscriptionId, object? body = null)
        {
            var queryParams = new Dictionary<string, string>();
            await SendAsync(new HttpMethod("PUT"), $"/webhooks/v2/disable/{Uri.EscapeDataString(subscriptionId)}", body, queryParams);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="subscriptionId">subscriptionId</param>
        /// <param name="body">Request body.</param>
        public async Task UpdateWebhooksEnableBySubscriptionId(string subscriptionId, object? body = null)
        {
            var queryParams = new Dictionary<string, string>();
            await SendAsync(new HttpMethod("PUT"), $"/webhooks/v2/enable/{Uri.EscapeDataString(subscriptionId)}", body, queryParams);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="body">Request body.</param>
        public async Task UpdateWebhooks(object? body = null)
        {
            var queryParams = new Dictionary<string, string>();
            await SendAsync(new HttpMethod("PUT"), $"/webhooks/v2", body, queryParams);
        }
    }
}

