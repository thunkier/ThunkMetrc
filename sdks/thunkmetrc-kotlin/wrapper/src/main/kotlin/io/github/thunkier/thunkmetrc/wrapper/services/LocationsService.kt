package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class LocationsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Creates new locations for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * POST /locations/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createLocations(licenseNumber: String? = null, body: List<LocationsCreateLocationsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.locations.createLocations(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Archives a specified Location, identified by its Id, for a Facility.
     * Permissions Required:
     * - Manage Locations
     * DELETE /locations/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteLocationsById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.locations.deleteLocationsById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Retrieves a list of active locations for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET /locations/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveLocations(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<LocationsLocation>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.locations.getActiveLocations(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<LocationsLocation>
    }

    /**
     * Retrieves a list of inactive locations for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET /locations/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveLocations(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<LocationsLocation>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.locations.getInactiveLocations(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<LocationsLocation>
    }

    /**
     * Retrieves a Location by its Id.
     * Permissions Required:
     * - Manage Locations
     * GET /locations/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getLocationsById(id: String, licenseNumber: String? = null): LocationsLocation? {
        return rateLimiter.execute(null,true,
        ) { 
            client.locations.getLocationsById(id, licenseNumber, ) 
        } as? LocationsLocation
    }

    /**
     * Retrieves a list of active location types for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET /locations/v2/types
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getLocationsTypes(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<LocationsType>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.locations.getLocationsTypes(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<LocationsType>
    }

    /**
     * Updates existing locations for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * PUT /locations/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateLocations(licenseNumber: String? = null, body: List<LocationsUpdateLocationsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.locations.updateLocations(licenseNumber, body) 
        } as? WriteResponse
    }
}

