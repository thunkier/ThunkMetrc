package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class SandboxClient(private val client: MetrcClient) {
    /**
     * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
     * POST CreateIntegratorSetup
     * @param userKey Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createSandboxIntegratorSetup(userKey: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sandbox/v2/integrator/setup")
        val query = ArrayList<String>()
        if (userKey != null) {
            query.add("userKey=${URLEncoder.encode(userKey, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
    }

}

