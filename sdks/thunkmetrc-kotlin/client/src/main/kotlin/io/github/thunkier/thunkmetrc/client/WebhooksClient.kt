package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class WebhooksClient(private val client: MetrcClient) {
    /**
     * 
     * DELETE DeleteWebhooksBySubscriptionId
     * @param subscriptionId Path parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun deleteWebhooksBySubscriptionId(subscriptionId: String, ): Any? {
        val path = StringBuilder("/webhooks/v2/${URLEncoder.encode(subscriptionId, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("DELETE", path.toString(), null)
    }

    /**
     * 
     * GET GetWebhooks
     * @throws IOException If request fails
     * @return Response object
     */
    fun getWebhooks(): Any? {
        val path = StringBuilder("/webhooks/v2")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * 
     * PUT UpdateDisableBySubscriptionId
     * @param subscriptionId Path parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateWebhooksDisableBySubscriptionId(subscriptionId: String, body: Any? = null): Any? {
        val path = StringBuilder("/webhooks/v2/disable/${URLEncoder.encode(subscriptionId, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

    /**
     * 
     * PUT UpdateEnableBySubscriptionId
     * @param subscriptionId Path parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateWebhooksEnableBySubscriptionId(subscriptionId: String, body: Any? = null): Any? {
        val path = StringBuilder("/webhooks/v2/enable/${URLEncoder.encode(subscriptionId, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

    /**
     * 
     * PUT UpdateWebhooks
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateWebhooks(body: Any? = null): Any? {
        val path = StringBuilder("/webhooks/v2")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

}

