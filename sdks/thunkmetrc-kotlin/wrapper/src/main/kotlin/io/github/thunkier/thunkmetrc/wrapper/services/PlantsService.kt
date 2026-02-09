package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class PlantsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Records additive usage for plants based on their location within a specified Facility.
     * Permissions Required:
     * - Manage Plants
     * - Manage Plants Additives
     * POST /plants/v2/additives/bylocation
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createAdditivesByLocation(licenseNumber: String? = null, body: List<PlantsCreateAdditivesByLocationRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.createPlantsAdditivesByLocation(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records additive usage for plants by location using a predefined additive template at a specified Facility.
     * Permissions Required:
     * - Manage Plants Additives
     * POST /plants/v2/additives/bylocation/usingtemplate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createAdditivesByLocationUsingTemplate(licenseNumber: String? = null, body: List<PlantsCreateAdditivesByLocationUsingTemplateRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.createPlantsAdditivesByLocationUsingTemplate(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates harvest product records from plant batches at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manicure/Harvest Veg/Flower Plants
     * POST /plants/v2/manicure
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createManicure(licenseNumber: String? = null, body: List<PlantsCreateManicureRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.createPlantsManicure(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates packages from plant batches at a specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST /plants/v2/plantbatch/packages
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantBatchPackages(licenseNumber: String? = null, body: List<PlantsCreatePlantBatchPackagesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.createPlantsPlantBatchPackages(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records additive usage details applied to specific plants at a Facility.
     * Permissions Required:
     * - Manage Plants Additives
     * POST /plants/v2/additives
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantsAdditives(licenseNumber: String? = null, body: List<PlantsCreatePlantsAdditivesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.createPlantsAdditives(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records additive usage for plants using predefined additive templates at a specified Facility.
     * Permissions Required:
     * - Manage Plants Additives
     * POST /plants/v2/additives/usingtemplate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantsAdditivesUsingTemplate(licenseNumber: String? = null, body: List<PlantsCreatePlantsAdditivesUsingTemplateRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.createPlantsAdditivesUsingTemplate(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates new plant batches at a specified Facility from existing plant data.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * POST /plants/v2/plantings
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantsPlantings(licenseNumber: String? = null, body: List<PlantsCreatePlantsPlantingsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.createPlantsPlantings(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records waste events for plants at a Facility, including method, reason, and location details.
     * Permissions Required:
     * - Manage Plants Waste
     * POST /plants/v2/waste
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantsWaste(licenseNumber: String? = null, body: List<PlantsCreatePlantsWasteRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.createPlantsWaste(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Removes plants from a Facility’s inventory while recording the reason for their disposal.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Destroy Veg/Flower Plants
     * DELETE /plants/v2
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deletePlants(licenseNumber: String? = null): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.deletePlants(licenseNumber, ) 
        } as? WriteResponse
    }

    /**
     * Retrieves additive records applied to plants at a specified Facility.
     * Permissions Required:
     * - View/Manage Plants Additives
     * GET /plants/v2/additives
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getAdditives(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Additive>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsAdditives(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Additive>
    }

    /**
     * Retrieves a list of all plant additive types defined within a Facility.
     * GET /plants/v2/additives/types
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getAdditivesTypes(): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsAdditivesTypes() 
        } as? Any
    }

    /**
     * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET /plants/v2/flowering
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getFloweringPlants(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Plant>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getFloweringPlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Plant>
    }

    /**
     * Retrieves the list of growth phases supported by a specified Facility.
     * GET /plants/v2/growthphases
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getGrowthPhases(licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsGrowthPhases(licenseNumber, ) 
        } as? Any
    }

    /**
     * Retrieves inactive plants at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET /plants/v2/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactivePlants(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Plant>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getInactivePlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Plant>
    }

    /**
     * Retrieves inactive mother-phase plants at a specified Facility.
     * Permissions Required:
     * - View Mother Plants
     * GET /plants/v2/mother/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getMotherInactivePlants(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Mother>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getMotherInactivePlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Mother>
    }

    /**
     * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
     * Permissions Required:
     * - View Mother Plants
     * GET /plants/v2/mother/onhold
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getMotherOnHoldPlants(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Mother>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getMotherOnHoldPlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Mother>
    }

    /**
     * Retrieves mother-phase plants at a specified Facility.
     * Permissions Required:
     * - View Mother Plants
     * GET /plants/v2/mother
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getMotherPlants(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Mother>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getMotherPlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Mother>
    }

    /**
     * Retrieves plants that are currently on hold at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET /plants/v2/onhold
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getOnHoldPlants(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Plant>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getOnHoldPlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Plant>
    }

    /**
     * Retrieves a Plant by Id.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET /plants/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantsById(id: String, licenseNumber: String? = null): Plant? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsById(id, licenseNumber, ) 
        } as? Plant
    }

    /**
     * Retrieves a Plant by label.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET /plants/v2/{label}
     * @param label Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantsByLabel(label: String, licenseNumber: String? = null): List<Plant>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsByLabel(label, licenseNumber, ) 
        } as? List<Plant>
    }

    /**
     * Retrieves a list of recorded plant waste events for a specific Facility.
     * Permissions Required:
     * - View Plants Waste
     * GET /plants/v2/waste
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantsWaste(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PlantsWaste>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsWaste(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PlantsWaste>
    }

    /**
     * Retrieves a list of all available plant waste methods for use within a Facility.
     * GET /plants/v2/waste/methods/all
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantsWasteMethods(pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<WasteMethod>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsWasteMethods(pageNumber, pageSize, ) 
        } as? PaginatedResponse<WasteMethod>
    }

    /**
     * Retriveves available reasons for recording mature plant waste at a specified Facility.
     * GET /plants/v2/waste/reasons
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantsWasteReasons(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<WasteReason>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsWasteReasons(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<WasteReason>
    }

    /**
     * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET /plants/v2/vegetative
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getVegetativePlants(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Plant>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getVegetativePlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Plant>
    }

    /**
     * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
     * Permissions Required:
     * - View Plants Waste
     * GET /plants/v2/waste/{id}/plant
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getWasteById(id: String, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PlantsWaste>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsWasteById(id, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PlantsWaste>
    }

    /**
     * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
     * Permissions Required:
     * - View Plants Waste
     * GET /plants/v2/waste/{id}/package
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getWastePackageById(id: String, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<WastePackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plants.getPlantsWastePackageById(id, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<WastePackage>
    }

    /**
     * Adjusts the recorded count of plants at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT /plants/v2/adjust
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateAdjustPlants(licenseNumber: String? = null, body: List<PlantsUpdateAdjustPlantsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.updateAdjustPlants(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Changes the growth phases of plants within a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT /plants/v2/growthphase
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateGrowthPhase(licenseNumber: String? = null, body: List<PlantsUpdateGrowthPhaseRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.updatePlantsGrowthPhase(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manicure/Harvest Veg/Flower Plants
     * PUT /plants/v2/harvest
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateHarvest(licenseNumber: String? = null, body: List<PlantsUpdateHarvestRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.updatePlantsHarvest(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Merges multiple plant groups into a single group within a Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manicure/Harvest Veg/Flower Plants
     * PUT /plants/v2/merge
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateMerge(licenseNumber: String? = null, body: List<PlantsUpdateMergeRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.updatePlantsMerge(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Moves plant batches to new locations within a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT /plants/v2/location
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updatePlantsLocation(licenseNumber: String? = null, body: List<PlantsUpdatePlantsLocationRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.updatePlantsLocation(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Updates the strain information for plants within a Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT /plants/v2/strain
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updatePlantsStrain(licenseNumber: String? = null, body: List<PlantsUpdatePlantsStrainRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.updatePlantsStrain(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Replaces existing plant tags with new tags for plants within a Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT /plants/v2/tag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updatePlantsTag(licenseNumber: String? = null, body: List<PlantsUpdatePlantsTagRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.updatePlantsTag(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Splits an existing plant group into multiple groups within a Facility.
     * Permissions Required:
     * - View Plant
     * PUT /plants/v2/split
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateSplit(licenseNumber: String? = null, body: List<PlantsUpdateSplitRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plants.updatePlantsSplit(licenseNumber, body) 
        } as? WriteResponse
    }
}

