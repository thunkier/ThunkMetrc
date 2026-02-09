package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class WebhooksService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * 
     * DELETE /webhooks/v2/{subscriptionId}
     * @param subscriptionId Path parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteWebhooksBySubscriptionId(subscriptionId: String): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.webhooks.deleteWebhooksBySubscriptionId(subscriptionId, ) 
        } as? Any
    }

    /**
     * 
     * GET /webhooks/v2
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getWebhooks(): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.webhooks.getWebhooks() 
        } as? Any
    }

    /**
     * 
     * PUT /webhooks/v2/disable/{subscriptionId}
     * @param subscriptionId Path parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDisableBySubscriptionId(subscriptionId: String, body: Any? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.webhooks.updateWebhooksDisableBySubscriptionId(subscriptionId, body) 
        } as? Any
    }

    /**
     * 
     * PUT /webhooks/v2/enable/{subscriptionId}
     * @param subscriptionId Path parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateEnableBySubscriptionId(subscriptionId: String, body: Any? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.webhooks.updateWebhooksEnableBySubscriptionId(subscriptionId, body) 
        } as? Any
    }

    /**
     * 
     * PUT /webhooks/v2
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateWebhooks(body: WebhooksUpdateWebhooksRequest): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.webhooks.updateWebhooks(body) 
        } as? WriteResponse
    }
}

