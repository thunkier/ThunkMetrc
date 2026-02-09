package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class HarvestsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public HarvestsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

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
    public WriteResponse createHarvestsPackages(
        String licenseNumber, List<HarvestsCreateHarvestsPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.harvests.createHarvestsPackages(
                licenseNumber, body
            )
        );
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
    public WriteResponse createHarvestsWaste(
        String licenseNumber, List<HarvestsCreateHarvestsWasteRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.harvests.createHarvestsWaste(
                licenseNumber, body
            )
        );
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
    public WriteResponse createHarvestsPackagesTesting(
        String licenseNumber, List<HarvestsCreatePackagesTestingRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.harvests.createHarvestsPackagesTesting(
                licenseNumber, body
            )
        );
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
    public Object deleteHarvestsWasteById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.harvests.deleteHarvestsWasteById(
                id, licenseNumber
            )
        );
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
    public WriteResponse finishHarvestsHarvests(
        String licenseNumber, List<HarvestsFinishHarvestsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.harvests.finishHarvestsHarvests(
                licenseNumber, body
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Harvest> getActiveHarvests(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Harvest>) client.harvests.getActiveHarvests(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    public Harvest getHarvestsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Harvest) client.harvests.getHarvestsById(
                id, licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<HarvestsWaste> getHarvestsWaste(
        String harvestId, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<HarvestsWaste>) client.harvests.getHarvestsWaste(
                harvestId, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Harvest> getInactiveHarvests(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Harvest>) client.harvests.getInactiveHarvests(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Harvest> getOnHoldHarvests(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Harvest>) client.harvests.getOnHoldHarvests(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of Waste types for harvests.
     * GET /harvests/v2/waste/types
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<WasteType> getHarvestsWasteTypes(
        String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<WasteType>) client.harvests.getHarvestsWasteTypes(
                pageNumber, pageSize
            )
        );
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
    public WriteResponse unfinishHarvestsHarvests(
        String licenseNumber, List<HarvestsUnfinishHarvestsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.harvests.unfinishHarvestsHarvests(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateHarvestsLocation(
        String licenseNumber, List<HarvestsUpdateHarvestsLocationRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.harvests.updateHarvestsLocation(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateHarvestsRename(
        String licenseNumber, List<HarvestsUpdateRenameRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.harvests.updateHarvestsRename(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateHarvestsRestoreHarvestedPlants(
        String licenseNumber, List<HarvestsUpdateRestoreHarvestedPlantsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.harvests.updateHarvestsRestoreHarvestedPlants(
                licenseNumber, body
            )
        );
    }

}

