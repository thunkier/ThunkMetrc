package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class UnitsOfMeasureService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Retrieves all active units of measure.
     * GET /unitsofmeasure/v2/active
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveUnitsOfMeasure(): PaginatedResponse<UnitOfMeasure>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.unitsOfMeasure.getActiveUnitsOfMeasure() 
        } as? PaginatedResponse<UnitOfMeasure>
    }

    /**
     * Retrieves all inactive units of measure.
     * GET /unitsofmeasure/v2/inactive
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveUnitsOfMeasure(): PaginatedResponse<UnitOfMeasure>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.unitsOfMeasure.getInactiveUnitsOfMeasure() 
        } as? PaginatedResponse<UnitOfMeasure>
    }
}

