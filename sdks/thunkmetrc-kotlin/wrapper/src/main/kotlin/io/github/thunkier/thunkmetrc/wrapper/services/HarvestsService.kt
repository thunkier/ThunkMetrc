package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class HarvestsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Creates packages from harvested products for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST /harvests/v2/packages
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createHarvestsPackages(licenseNumber: String? = null, body: List<HarvestsCreateHarvestsPackagesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.createHarvestsPackages(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * POST /harvests/v2/waste
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createHarvestsWaste(licenseNumber: String? = null, body: List<HarvestsCreateHarvestsWasteRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.createHarvestsWaste(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates packages for testing from harvested products for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST /harvests/v2/packages/testing
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPackagesTesting(licenseNumber: String? = null, body: List<HarvestsCreatePackagesTestingRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.createHarvestsPackagesTesting(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Discontinues a specific harvest waste record by Id for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Discontinue Harvest Waste
     * DELETE /harvests/v2/waste/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteWasteById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.deleteHarvestsWasteById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Marks one or more harvests as finished for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Finish/Discontinue Harvests
     * PUT /harvests/v2/finish
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun finishHarvests(licenseNumber: String? = null, body: List<HarvestsFinishHarvestsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.finishHarvestsHarvests(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Retrieves a list of active harvests for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * GET /harvests/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Harvest>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.harvests.getActiveHarvests(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Harvest>
    }

    /**
     * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
     * Permissions Required:
     * - View Harvests
     * GET /harvests/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getHarvestsById(id: String, licenseNumber: String? = null): Harvest? {
        return rateLimiter.execute(null,true,
        ) { 
            client.harvests.getHarvestsById(id, licenseNumber, ) 
        } as? Harvest
    }

    /**
     * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
     * Permissions Required:
     * - View Harvests
     * GET /harvests/v2/waste
     * @param harvestId Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getHarvestsWaste(harvestId: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<HarvestsWaste>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.harvests.getHarvestsWaste(harvestId, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<HarvestsWaste>
    }

    /**
     * Retrieves a list of inactive harvests for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * GET /harvests/v2/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Harvest>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.harvests.getInactiveHarvests(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Harvest>
    }

    /**
     * Retrieves a list of harvests on hold for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * GET /harvests/v2/onhold
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getOnHoldHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Harvest>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.harvests.getOnHoldHarvests(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Harvest>
    }

    /**
     * Retrieves a list of Waste types for harvests.
     * GET /harvests/v2/waste/types
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getWasteTypes(pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<WasteType>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.harvests.getHarvestsWasteTypes(pageNumber, pageSize, ) 
        } as? PaginatedResponse<WasteType>
    }

    /**
     * Reopens one or more previously finished harvests for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Finish/Discontinue Harvests
     * PUT /harvests/v2/unfinish
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun unfinishHarvests(licenseNumber: String? = null, body: List<HarvestsUnfinishHarvestsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.unfinishHarvestsHarvests(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Updates the Location of Harvest for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * PUT /harvests/v2/location
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateHarvestsLocation(licenseNumber: String? = null, body: List<HarvestsUpdateHarvestsLocationRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.updateHarvestsLocation(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Renames one or more harvests for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * PUT /harvests/v2/rename
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateRename(licenseNumber: String? = null, body: List<HarvestsUpdateRenameRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.updateHarvestsRename(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Restores previously harvested plants to their original state for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Finish/Discontinue Harvests
     * PUT /harvests/v2/restore/harvestedplants
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateRestoreHarvestedPlants(licenseNumber: String? = null, body: List<HarvestsUpdateRestoreHarvestedPlantsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.harvests.updateHarvestsRestoreHarvestedPlants(licenseNumber, body) 
        } as? WriteResponse
    }
}

