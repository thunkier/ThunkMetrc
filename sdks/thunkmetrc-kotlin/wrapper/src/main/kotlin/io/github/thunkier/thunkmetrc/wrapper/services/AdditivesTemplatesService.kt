package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class AdditivesTemplatesService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Creates new additive templates for a specified Facility.
     * Permissions Required:
     * - Manage Additives
     * POST /additivestemplates/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createAdditivesTemplates(licenseNumber: String? = null, body: List<AdditivesTemplatesCreateAdditivesTemplatesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.additivesTemplates.createAdditivesTemplates(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Retrieves a list of active additive templates for a specified Facility.
     * Permissions Required:
     * - Manage Additives
     * GET /additivestemplates/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveAdditivesTemplates(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<AdditivesTemplate>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.additivesTemplates.getActiveAdditivesTemplates(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<AdditivesTemplate>
    }

    /**
     * Retrieves an Additive Template by its Id.
     * Permissions Required:
     * - Manage Additives
     * GET /additivestemplates/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getAdditivesTemplatesById(id: String, licenseNumber: String? = null): AdditivesTemplate? {
        return rateLimiter.execute(null,true,
        ) { 
            client.additivesTemplates.getAdditivesTemplatesById(id, licenseNumber, ) 
        } as? AdditivesTemplate
    }

    /**
     * Retrieves a list of inactive additive templates for a specified Facility.
     * Permissions Required:
     * - Manage Additives
     * GET /additivestemplates/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveAdditivesTemplates(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<AdditivesTemplate>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.additivesTemplates.getInactiveAdditivesTemplates(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<AdditivesTemplate>
    }

    /**
     * Updates existing additive templates for a specified Facility.
     * Permissions Required:
     * - Manage Additives
     * PUT /additivestemplates/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateAdditivesTemplates(licenseNumber: String? = null, body: List<AdditivesTemplatesUpdateAdditivesTemplatesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.additivesTemplates.updateAdditivesTemplates(licenseNumber, body) 
        } as? WriteResponse
    }
}

