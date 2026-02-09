package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class PlantBatchesService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * POST /plantbatches/v2/adjust
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createAdjustPlantBatches(licenseNumber: String? = null, body: List<PlantBatchesCreateAdjustPlantBatchesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createAdjustPlantBatches(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Updates the growth phase of plants at a specified Facility based on tracking information.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * POST /plantbatches/v2/growthphase
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createGrowthPhase(licenseNumber: String? = null, body: List<PlantBatchesCreateGrowthPhaseRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createPlantBatchesGrowthPhase(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates packages from mother plants at the specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST /plantbatches/v2/packages/frommotherplant
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPackagesFromMotherPlant(licenseNumber: String? = null, body: List<PlantBatchesCreatePackagesFromMotherPlantRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createPlantBatchesPackagesFromMotherPlant(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records Additive usage details for plant batches at a specific Facility.
     * Permissions Required:
     * - Manage Plants Additives
     * POST /plantbatches/v2/additives
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantBatchesAdditives(licenseNumber: String? = null, body: List<PlantBatchesCreatePlantBatchesAdditivesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createPlantBatchesAdditives(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records Additive usage for plant batches at a Facility using predefined additive templates.
     * Permissions Required:
     * - Manage Plants Additives
     * POST /plantbatches/v2/additives/usingtemplate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantBatchesAdditivesUsingTemplate(licenseNumber: String? = null, body: List<PlantBatchesCreatePlantBatchesAdditivesUsingTemplateRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createPlantBatchesAdditivesUsingTemplate(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST /plantbatches/v2/packages
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantBatchesPackages(licenseNumber: String? = null, body: List<PlantBatchesCreatePlantBatchesPackagesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createPlantBatchesPackages(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates new plantings for a Facility by generating plant batches based on provided planting details.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * POST /plantbatches/v2/plantings
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantBatchesPlantings(licenseNumber: String? = null, body: List<PlantBatchesCreatePlantBatchesPlantingsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createPlantBatchesPlantings(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records waste information for plant batches based on the submitted data for the specified Facility.
     * Permissions Required:
     * - Manage Plants Waste
     * POST /plantbatches/v2/waste
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPlantBatchesWaste(licenseNumber: String? = null, body: List<PlantBatchesCreatePlantBatchesWasteRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createPlantBatchesWaste(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Splits an existing Plant Batch into multiple groups at the specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * POST /plantbatches/v2/split
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createSplit(licenseNumber: String? = null, body: List<PlantBatchesCreateSplitRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.createPlantBatchesSplit(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Completes the destruction of plant batches based on the provided input data.
     * Permissions Required:
     * - View Immature Plants
     * - Destroy Immature Plants
     * DELETE /plantbatches/v2
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deletePlantBatches(licenseNumber: String? = null): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.deletePlantBatches(licenseNumber, ) 
        } as? WriteResponse
    }

    /**
     * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
     * Permissions Required:
     * - View Immature Plants
     * GET /plantbatches/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActivePlantBatches(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PlantBatch>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plantBatches.getActivePlantBatches(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PlantBatch>
    }

    /**
     * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
     * Permissions Required:
     * - View Immature Plants
     * GET /plantbatches/v2/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactivePlantBatches(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PlantBatch>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plantBatches.getInactivePlantBatches(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PlantBatch>
    }

    /**
     * Retrieves a Plant Batch by Id.
     * Permissions Required:
     * - View Immature Plants
     * GET /plantbatches/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantBatchesById(id: String, licenseNumber: String? = null): PlantBatch? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plantBatches.getPlantBatchesById(id, licenseNumber, ) 
        } as? PlantBatch
    }

    /**
     * Retrieves a list of plant batch types.
     * GET /plantbatches/v2/types
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantBatchesTypes(pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PlantBatchesType>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plantBatches.getPlantBatchesTypes(pageNumber, pageSize, ) 
        } as? PaginatedResponse<PlantBatchesType>
    }

    /**
     * Retrieves waste details associated with plant batches at a specified Facility.
     * Permissions Required:
     * - View Plants Waste
     * GET /plantbatches/v2/waste
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantBatchesWaste(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PlantBatchesWaste>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plantBatches.getPlantBatchesWaste(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PlantBatchesWaste>
    }

    /**
     * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
     * GET /plantbatches/v2/waste/reasons
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPlantBatchesWasteReasons(licenseNumber: String? = null): List<PlantBatchesWasteReason>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.plantBatches.getPlantBatchesWasteReasons(licenseNumber, ) 
        } as? List<PlantBatchesWasteReason>
    }

    /**
     * Renames plant batches at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT /plantbatches/v2/name
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateName(licenseNumber: String? = null, body: List<PlantBatchesUpdateNameRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.updatePlantBatchesName(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Moves one or more plant batches to new locations with in a specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants
     * PUT /plantbatches/v2/location
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updatePlantBatchesLocation(licenseNumber: String? = null, body: List<PlantBatchesUpdatePlantBatchesLocationRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.updatePlantBatchesLocation(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Changes the strain of plant batches at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT /plantbatches/v2/strain
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updatePlantBatchesStrain(licenseNumber: String? = null, body: List<PlantBatchesUpdatePlantBatchesStrainRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.updatePlantBatchesStrain(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Replaces tags for plant batches at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT /plantbatches/v2/tag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updatePlantBatchesTag(licenseNumber: String? = null, body: List<PlantBatchesUpdatePlantBatchesTagRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.plantBatches.updatePlantBatchesTag(licenseNumber, body) 
        } as? WriteResponse
    }
}

