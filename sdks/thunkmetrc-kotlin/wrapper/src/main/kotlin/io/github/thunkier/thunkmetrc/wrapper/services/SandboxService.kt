package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class SandboxService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
     * POST /sandbox/v2/integrator/setup
     * @param userKey Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createIntegratorSetup(userKey: String? = null, body: Any? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sandbox.createSandboxIntegratorSetup(userKey, body) 
        } as? Any
    }
}

