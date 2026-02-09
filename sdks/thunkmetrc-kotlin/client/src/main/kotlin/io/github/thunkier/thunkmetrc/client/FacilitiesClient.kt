package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class FacilitiesClient(private val client: MetrcClient) {
    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * GET GetFacilities
     * @throws IOException If request fails
     * @return Response object
     */
    fun getFacilities(): Any? {
        val path = StringBuilder("/facilities/v2")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

}

