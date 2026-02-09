package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;

public class WebhooksService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public WebhooksService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * 
     * DELETE /webhooks/v2/{subscriptionId}
     * @param subscriptionId Path parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteWebhooksBySubscriptionId(
        String subscriptionId
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.webhooks.deleteWebhooksBySubscriptionId(
                subscriptionId
            )
        );
    }

    /**
     * 
     * GET /webhooks/v2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object getWebhooks(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.webhooks.getWebhooks(
                
            )
        );
    }

    /**
     * 
     * PUT /webhooks/v2/disable/{subscriptionId}
     * @param subscriptionId Path parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public Object updateWebhooksDisableBySubscriptionId(
        String subscriptionId, Object body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.webhooks.updateWebhooksDisableBySubscriptionId(
                subscriptionId, body
            )
        );
    }

    /**
     * 
     * PUT /webhooks/v2/enable/{subscriptionId}
     * @param subscriptionId Path parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public Object updateWebhooksEnableBySubscriptionId(
        String subscriptionId, Object body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.webhooks.updateWebhooksEnableBySubscriptionId(
                subscriptionId, body
            )
        );
    }

    /**
     * 
     * PUT /webhooks/v2
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateWebhooks(
        WebhooksUpdateWebhooksRequest body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.webhooks.updateWebhooks(
                body
            )
        );
    }

}

