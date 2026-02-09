package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class PackagesService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public PackagesService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Records a list of adjustments for packages at a specific Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * POST /packages/v2/adjust
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createAdjustPackages(
        String licenseNumber, List<PackagesCreateAdjustPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.createAdjustPackages(
                licenseNumber, body
            )
        );
    }

    /**
     * Creates new packages for a specified Facility.
     * Permissions Required:
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST /packages/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createPackagesPackages(
        String licenseNumber, List<PackagesCreatePackagesPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.createPackagesPackages(
                licenseNumber, body
            )
        );
    }

    /**
     * Creates new plantings from packages for a specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants
     * - View Packages
     * - Manage Packages Inventory
     * POST /packages/v2/plantings
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createPackagesPlantings(
        String licenseNumber, List<PackagesCreatePackagesPlantingsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.createPackagesPlantings(
                licenseNumber, body
            )
        );
    }

    /**
     * Creates new packages for testing for a specified Facility.
     * Permissions Required:
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST /packages/v2/testing
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createPackagesTesting(
        String licenseNumber, List<PackagesCreateTestingRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.createPackagesTesting(
                licenseNumber, body
            )
        );
    }

    /**
     * Discontinues a Package at a specific Facility.
     * Permissions Required:
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * DELETE /packages/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deletePackagesById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.packages.deletePackagesById(
                id, licenseNumber
            )
        );
    }

    /**
     * Updates a list of packages as finished for a specific Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/finish
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public Object finishPackagesPackages(
        String licenseNumber, List<PackagesFinishPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.packages.finishPackagesPackages(
                licenseNumber, body
            )
        );
    }

    /**
     * Flags one or more Packages at the specified Facility as Finished Goods.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/finishedgood/flag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse finishedgoodFlagPackages(
        String licenseNumber, List<PackagesFinishedgoodFlagRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.finishedgoodFlagPackages(
                licenseNumber, body
            )
        );
    }

    /**
     * Removes the Finished Good flag one or more Packages at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/finishedgood/unflag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse finishedgoodUnflagPackages(
        String licenseNumber, List<PackagesFinishedgoodUnflagRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.finishedgoodUnflagPackages(
                licenseNumber, body
            )
        );
    }

    /**
     * Retrieves a list of active packages for a specified Facility.
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PackagesPackage> getActivePackages(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PackagesPackage>) client.packages.getActivePackages(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of adjustment reasons for packages at a specified Facility.
     * GET /packages/v2/adjust/reasons
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<AdjustReason> getPackagesAdjustReasons(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<AdjustReason>) client.packages.getPackagesAdjustReasons(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves the Package Adjustments for a Facility
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/adjustments
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Adjustment> getPackagesAdjustments(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Adjustment>) client.packages.getPackagesAdjustments(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of packages in transit for a specified Facility.
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/intransit
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<InTransit> getInTransitPackages(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<InTransit>) client.packages.getInTransitPackages(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of inactive packages for a specified Facility.
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PackagesPackage> getInactivePackages(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PackagesPackage>) client.packages.getInactivePackages(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/labsamples
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PackagesPackage> getPackagesLabSamples(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PackagesPackage>) client.packages.getPackagesLabSamples(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of packages on hold for a specified Facility.
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/onhold
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PackagesPackage> getOnHoldPackages(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PackagesPackage>) client.packages.getOnHoldPackages(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a Package by its Id.
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public PackagesPackage getPackagesById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PackagesPackage) client.packages.getPackagesById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a Package by its label.
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/{label}
     * @param label Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<PackagesPackage> getPackagesByLabel(
        String label, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<PackagesPackage>) client.packages.getPackagesByLabel(
                label, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of available Package types.
     * GET /packages/v2/types
     * @throws Exception If request fails
     * @return Response object
     */
    public Object getPackagesTypes(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.packages.getPackagesTypes(
                
            )
        );
    }

    /**
     * Retrieves the source harvests for a Package by its Id.
     * Permissions Required:
     * - View Package Source Harvests
     * GET /packages/v2/{id}/source/harvests
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public SourceHarvest getPackagesSourceHarvestById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (SourceHarvest) client.packages.getPackagesSourceHarvestById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of transferred packages for a specific Facility.
     * Permissions Required:
     * - View Packages
     * GET /packages/v2/transferred
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PackagesPackage> getPackagesTransferred(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PackagesPackage>) client.packages.getPackagesTransferred(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Updates a list of packages as unfinished for a specific Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/unfinish
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public Object unfinishPackagesPackages(
        String licenseNumber, List<PackagesUnfinishPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.packages.unfinishPackagesPackages(
                licenseNumber, body
            )
        );
    }

    /**
     * Set the final quantity for a Package.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/adjust
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateAdjustPackages(
        String licenseNumber, List<PackagesUpdateAdjustPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updateAdjustPackages(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates the Product decontaminate information for a list of packages at a specific Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/decontaminate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesDecontaminate(
        String licenseNumber, List<PackagesUpdateDecontaminateRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesDecontaminate(
                licenseNumber, body
            )
        );
    }

    /**
     * Flags one or more packages for donation at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/donation/flag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesDonationFlag(
        String licenseNumber, List<PackagesUpdateDonationFlagRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesDonationFlag(
                licenseNumber, body
            )
        );
    }

    /**
     * Removes the donation flag from one or more packages at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/donation/unflag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesDonationUnflag(
        String licenseNumber, List<PackagesUpdateDonationUnflagRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesDonationUnflag(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates the external identifiers for one or more packages at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Package Inventory
     * - External Id Enabled
     * PUT /packages/v2/externalid
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesExternalid(
        String licenseNumber, List<PackagesUpdateExternalidRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesExternalid(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates the associated Item for one or more packages at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * PUT /packages/v2/item
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesItem(
        String licenseNumber, List<PackagesUpdateItemRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesItem(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates the list of required lab test batches for one or more packages at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * PUT /packages/v2/labtests/required
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesLabtestsRequired(
        String licenseNumber, List<PackagesUpdateLabtestsRequiredRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesLabtestsRequired(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates notes associated with one or more packages for the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * - Manage Package Notes
     * PUT /packages/v2/note
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesNote(
        String licenseNumber, List<PackagesUpdateNoteRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesNote(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates the Location and Sublocation for one or more packages at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * PUT /packages/v2/location
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesLocation(
        String licenseNumber, List<PackagesUpdatePackagesLocationRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesLocation(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates a list of Product remediations for packages at a specific Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/remediate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesRemediate(
        String licenseNumber, List<PackagesUpdateRemediateRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesRemediate(
                licenseNumber, body
            )
        );
    }

    /**
     * Flags or unflags one or more packages at the specified Facility as trade samples.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/tradesample/flag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesTradeSampleFlag(
        String licenseNumber, List<PackagesUpdateTradeSampleFlagRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesTradeSampleFlag(
                licenseNumber, body
            )
        );
    }

    /**
     * Removes the trade sample flag from one or more packages at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /packages/v2/tradesample/unflag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesTradeSampleUnflag(
        String licenseNumber, List<PackagesUpdateTradeSampleUnflagRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesTradeSampleUnflag(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates the use-by date for one or more packages at the specified Facility.
     * Permissions Required:
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * PUT /packages/v2/usebydate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePackagesUseByDate(
        String licenseNumber, List<PackagesUpdateUseByDateRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.packages.updatePackagesUseByDate(
                licenseNumber, body
            )
        );
    }

}

