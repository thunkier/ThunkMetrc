package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class PlantBatchesService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public PlantBatchesService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

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
    public WriteResponse createAdjustPlantBatches(
        String licenseNumber, List<PlantBatchesCreateAdjustPlantBatchesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createAdjustPlantBatches(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantBatchesGrowthPhase(
        String licenseNumber, List<PlantBatchesCreateGrowthPhaseRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createPlantBatchesGrowthPhase(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantBatchesPackagesFromMotherPlant(
        String licenseNumber, List<PlantBatchesCreatePackagesFromMotherPlantRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createPlantBatchesPackagesFromMotherPlant(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantBatchesAdditives(
        String licenseNumber, List<PlantBatchesCreatePlantBatchesAdditivesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createPlantBatchesAdditives(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantBatchesAdditivesUsingTemplate(
        String licenseNumber, List<PlantBatchesCreatePlantBatchesAdditivesUsingTemplateRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createPlantBatchesAdditivesUsingTemplate(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantBatchesPackages(
        String licenseNumber, List<PlantBatchesCreatePlantBatchesPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createPlantBatchesPackages(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantBatchesPlantings(
        String licenseNumber, List<PlantBatchesCreatePlantBatchesPlantingsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createPlantBatchesPlantings(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantBatchesWaste(
        String licenseNumber, List<PlantBatchesCreatePlantBatchesWasteRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createPlantBatchesWaste(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantBatchesSplit(
        String licenseNumber, List<PlantBatchesCreateSplitRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.createPlantBatchesSplit(
                licenseNumber, body
            )
        );
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
    public WriteResponse deletePlantBatches(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.deletePlantBatches(
                licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PlantBatch> getActivePlantBatches(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PlantBatch>) client.plantBatches.getActivePlantBatches(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PlantBatch> getInactivePlantBatches(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PlantBatch>) client.plantBatches.getInactivePlantBatches(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    public PlantBatch getPlantBatchesById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PlantBatch) client.plantBatches.getPlantBatchesById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of plant batch types.
     * GET /plantbatches/v2/types
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PlantBatchesType> getPlantBatchesTypes(
        String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PlantBatchesType>) client.plantBatches.getPlantBatchesTypes(
                pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PlantBatchesWaste> getPlantBatchesWaste(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PlantBatchesWaste>) client.plantBatches.getPlantBatchesWaste(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
     * GET /plantbatches/v2/waste/reasons
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<PlantBatchesWasteReason> getPlantBatchesWasteReasons(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<PlantBatchesWasteReason>) client.plantBatches.getPlantBatchesWasteReasons(
                licenseNumber
            )
        );
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
    public WriteResponse updatePlantBatchesName(
        String licenseNumber, List<PlantBatchesUpdateNameRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.updatePlantBatchesName(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantBatchesLocation(
        String licenseNumber, List<PlantBatchesUpdatePlantBatchesLocationRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.updatePlantBatchesLocation(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantBatchesStrain(
        String licenseNumber, List<PlantBatchesUpdatePlantBatchesStrainRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.updatePlantBatchesStrain(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantBatchesTag(
        String licenseNumber, List<PlantBatchesUpdatePlantBatchesTagRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plantBatches.updatePlantBatchesTag(
                licenseNumber, body
            )
        );
    }

}

