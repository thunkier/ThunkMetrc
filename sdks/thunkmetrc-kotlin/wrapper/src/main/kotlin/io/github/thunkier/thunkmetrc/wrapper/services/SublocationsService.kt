package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class SublocationsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Creates new sublocation records for a Facility.
     * Permissions Required:
     * - Manage Locations
     * POST /sublocations/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createSublocations(licenseNumber: String? = null, body: List<SublocationsCreateSublocationsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sublocations.createSublocations(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Archives an existing Sublocation record for a Facility.
     * Permissions Required:
     * - Manage Locations
     * DELETE /sublocations/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteSublocationsById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sublocations.deleteSublocationsById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Locations
     * GET /sublocations/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveSublocations(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Sublocation>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sublocations.getActiveSublocations(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Sublocation>
    }

    /**
     * Retrieves a list of inactive sublocations for the specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET /sublocations/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveSublocations(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Sublocation>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sublocations.getInactiveSublocations(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Sublocation>
    }

    /**
     * Retrieves a Sublocation by its Id, with an optional license number.
     * Permissions Required:
     * - Manage Locations
     * GET /sublocations/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getSublocationsById(id: String, licenseNumber: String? = null): Sublocation? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sublocations.getSublocationsById(id, licenseNumber, ) 
        } as? Sublocation
    }

    /**
     * Updates existing sublocation records for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * PUT /sublocations/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateSublocations(licenseNumber: String? = null, body: List<SublocationsUpdateSublocationsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sublocations.updateSublocations(licenseNumber, body) 
        } as? WriteResponse
    }
}

