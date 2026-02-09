package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class WebhooksClient {
    private final MetrcClient client;

    public WebhooksClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * 
     * DELETE DeleteWebhooksBySubscriptionId
     * @param subscriptionId Path parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteWebhooksBySubscriptionId(
        String subscriptionId
    ) throws IOException {
        StringBuilder path = new StringBuilder("/webhooks/v2/"+URLEncoder.encode(subscriptionId, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * 
     * GET GetWebhooks
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getWebhooks(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/webhooks/v2");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * 
     * PUT UpdateDisableBySubscriptionId
     * @param subscriptionId Path parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateWebhooksDisableBySubscriptionId(
        String subscriptionId, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/webhooks/v2/disable/"+URLEncoder.encode(subscriptionId, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * 
     * PUT UpdateEnableBySubscriptionId
     * @param subscriptionId Path parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateWebhooksEnableBySubscriptionId(
        String subscriptionId, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/webhooks/v2/enable/"+URLEncoder.encode(subscriptionId, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * 
     * PUT UpdateWebhooks
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateWebhooks(
        Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/webhooks/v2");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

