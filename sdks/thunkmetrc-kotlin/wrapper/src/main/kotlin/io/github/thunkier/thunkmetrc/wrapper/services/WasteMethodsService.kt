package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class WasteMethodsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Retrieves all available waste methods.
     * GET /wastemethods/v2
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getWasteMethodsWasteMethods(): List<WasteMethod>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.wasteMethods.getWasteMethodsWasteMethods() 
        } as? List<WasteMethod>
    }
}

