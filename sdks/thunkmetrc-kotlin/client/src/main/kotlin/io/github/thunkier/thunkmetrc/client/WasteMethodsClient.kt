package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class WasteMethodsClient(private val client: MetrcClient) {
    /**
     * Retrieves all available waste methods.
     * GET GetWasteMethodsWasteMethods
     * @throws IOException If request fails
     * @return Response object
     */
    fun getWasteMethodsWasteMethods(): Any? {
        val path = StringBuilder("/wastemethods/v2")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

}

