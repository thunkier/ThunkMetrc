package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class StrainsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Creates new strain records for a specified Facility.
     * Permissions Required:
     * - Manage Strains
     * POST /strains/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createStrains(licenseNumber: String? = null, body: List<StrainsCreateStrainsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.strains.createStrains(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Archives an existing strain record for a Facility
     * Permissions Required:
     * - Manage Strains
     * DELETE /strains/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteStrainsById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.strains.deleteStrainsById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Strains
     * GET /strains/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveStrains(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Strain>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.strains.getActiveStrains(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Strain>
    }

    /**
     * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Strains
     * GET /strains/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveStrains(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Strain>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.strains.getInactiveStrains(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Strain>
    }

    /**
     * Retrieves a Strain record by its Id, with an optional license number.
     * Permissions Required:
     * - Manage Strains
     * GET /strains/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getStrainsById(id: String, licenseNumber: String? = null): Strain? {
        return rateLimiter.execute(null,true,
        ) { 
            client.strains.getStrainsById(id, licenseNumber, ) 
        } as? Strain
    }

    /**
     * Updates existing strain records for a specified Facility.
     * Permissions Required:
     * - Manage Strains
     * PUT /strains/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateStrains(licenseNumber: String? = null, body: List<StrainsUpdateStrainsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.strains.updateStrains(licenseNumber, body) 
        } as? WriteResponse
    }
}

