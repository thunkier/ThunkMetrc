package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class PlantsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public PlantsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

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
    public WriteResponse createPlantsAdditivesByLocation(
        String licenseNumber, List<PlantsCreateAdditivesByLocationRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.createPlantsAdditivesByLocation(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantsAdditivesByLocationUsingTemplate(
        String licenseNumber, List<PlantsCreateAdditivesByLocationUsingTemplateRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.createPlantsAdditivesByLocationUsingTemplate(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantsManicure(
        String licenseNumber, List<PlantsCreateManicureRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.createPlantsManicure(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantsPlantBatchPackages(
        String licenseNumber, List<PlantsCreatePlantBatchPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.createPlantsPlantBatchPackages(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantsAdditives(
        String licenseNumber, List<PlantsCreatePlantsAdditivesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.createPlantsAdditives(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantsAdditivesUsingTemplate(
        String licenseNumber, List<PlantsCreatePlantsAdditivesUsingTemplateRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.createPlantsAdditivesUsingTemplate(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantsPlantings(
        String licenseNumber, List<PlantsCreatePlantsPlantingsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.createPlantsPlantings(
                licenseNumber, body
            )
        );
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
    public WriteResponse createPlantsWaste(
        String licenseNumber, List<PlantsCreatePlantsWasteRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.createPlantsWaste(
                licenseNumber, body
            )
        );
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
    public WriteResponse deletePlants(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.deletePlants(
                licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Additive> getPlantsAdditives(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Additive>) client.plants.getPlantsAdditives(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of all plant additive types defined within a Facility.
     * GET /plants/v2/additives/types
     * @throws Exception If request fails
     * @return Response object
     */
    public Object getPlantsAdditivesTypes(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.plants.getPlantsAdditivesTypes(
                
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Plant> getFloweringPlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Plant>) client.plants.getFloweringPlants(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves the list of growth phases supported by a specified Facility.
     * GET /plants/v2/growthphases
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object getPlantsGrowthPhases(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.plants.getPlantsGrowthPhases(
                licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Plant> getInactivePlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Plant>) client.plants.getInactivePlants(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Mother> getMotherInactivePlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Mother>) client.plants.getMotherInactivePlants(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Mother> getMotherOnHoldPlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Mother>) client.plants.getMotherOnHoldPlants(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Mother> getMotherPlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Mother>) client.plants.getMotherPlants(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Plant> getOnHoldPlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Plant>) client.plants.getOnHoldPlants(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    public Plant getPlantsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Plant) client.plants.getPlantsById(
                id, licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public List<Plant> getPlantsByLabel(
        String label, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<Plant>) client.plants.getPlantsByLabel(
                label, licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PlantsWaste> getPlantsWaste(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PlantsWaste>) client.plants.getPlantsWaste(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of all available plant waste methods for use within a Facility.
     * GET /plants/v2/waste/methods/all
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<WasteMethod> getPlantsWasteMethods(
        String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<WasteMethod>) client.plants.getPlantsWasteMethods(
                pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<WasteReason> getPlantsWasteReasons(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<WasteReason>) client.plants.getPlantsWasteReasons(
                licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Plant> getVegetativePlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Plant>) client.plants.getVegetativePlants(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PlantsWaste> getPlantsWasteById(
        String id, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PlantsWaste>) client.plants.getPlantsWasteById(
                id, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<WastePackage> getPlantsWastePackageById(
        String id, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<WastePackage>) client.plants.getPlantsWastePackageById(
                id, licenseNumber, pageNumber, pageSize
            )
        );
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
    public WriteResponse updateAdjustPlants(
        String licenseNumber, List<PlantsUpdateAdjustPlantsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.updateAdjustPlants(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantsGrowthPhase(
        String licenseNumber, List<PlantsUpdateGrowthPhaseRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.updatePlantsGrowthPhase(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantsHarvest(
        String licenseNumber, List<PlantsUpdateHarvestRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.updatePlantsHarvest(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantsMerge(
        String licenseNumber, List<PlantsUpdateMergeRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.updatePlantsMerge(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantsLocation(
        String licenseNumber, List<PlantsUpdatePlantsLocationRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.updatePlantsLocation(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantsStrain(
        String licenseNumber, List<PlantsUpdatePlantsStrainRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.updatePlantsStrain(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantsTag(
        String licenseNumber, List<PlantsUpdatePlantsTagRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.updatePlantsTag(
                licenseNumber, body
            )
        );
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
    public WriteResponse updatePlantsSplit(
        String licenseNumber, List<PlantsUpdateSplitRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.plants.updatePlantsSplit(
                licenseNumber, body
            )
        );
    }

}

