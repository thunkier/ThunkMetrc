package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class FacilitiesService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * GET /facilities/v2
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getFacilities(): List<Facility>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.facilities.getFacilities() 
        } as? List<Facility>
    }
}

