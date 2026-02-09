package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class TagsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Returns a list of available package tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * GET /tags/v2/package/available
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getAvailablePackage(licenseNumber: String? = null): List<Tag>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.tags.getTagsAvailablePackage(licenseNumber, ) 
        } as? List<Tag>
    }

    /**
     * Returns a list of available plant tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * GET /tags/v2/plant/available
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getAvailablePlant(licenseNumber: String? = null): List<Tag>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.tags.getTagsAvailablePlant(licenseNumber, ) 
        } as? List<Tag>
    }

    /**
     * Returns a list of staged tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * - RetailId.AllowPackageStaging Key Value enabled
     * GET /tags/v2/staged
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getStagedTags(licenseNumber: String? = null): List<Staged>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.tags.getStagedTags(licenseNumber, ) 
        } as? List<Staged>
    }
}

