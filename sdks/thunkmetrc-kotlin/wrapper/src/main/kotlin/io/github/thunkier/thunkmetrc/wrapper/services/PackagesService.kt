package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class PackagesService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

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
    suspend fun createAdjustPackages(licenseNumber: String? = null, body: List<PackagesCreateAdjustPackagesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.createAdjustPackages(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun createPackagesPackages(licenseNumber: String? = null, body: List<PackagesCreatePackagesPackagesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.createPackagesPackages(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun createPackagesPlantings(licenseNumber: String? = null, body: List<PackagesCreatePackagesPlantingsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.createPackagesPlantings(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun createTesting(licenseNumber: String? = null, body: List<PackagesCreateTestingRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.createPackagesTesting(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun deletePackagesById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.deletePackagesById(id, licenseNumber, ) 
        } as? Any
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
    suspend fun finishPackages(licenseNumber: String? = null, body: List<PackagesFinishPackagesRequestItem>): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.finishPackagesPackages(licenseNumber, body) 
        } as? Any
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
    suspend fun finishedgoodFlag(licenseNumber: String? = null, body: List<PackagesFinishedgoodFlagRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.finishedgoodFlagPackages(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun finishedgoodUnflag(licenseNumber: String? = null, body: List<PackagesFinishedgoodUnflagRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.finishedgoodUnflagPackages(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun getActivePackages(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PackagesPackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getActivePackages(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PackagesPackage>
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
    suspend fun getAdjustReasons(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<AdjustReason>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getPackagesAdjustReasons(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<AdjustReason>
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
    suspend fun getAdjustments(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Adjustment>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getPackagesAdjustments(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Adjustment>
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
    suspend fun getInTransitPackages(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<InTransit>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getInTransitPackages(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<InTransit>
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
    suspend fun getInactivePackages(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PackagesPackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getInactivePackages(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PackagesPackage>
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
    suspend fun getLabSamples(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PackagesPackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getPackagesLabSamples(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PackagesPackage>
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
    suspend fun getOnHoldPackages(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PackagesPackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getOnHoldPackages(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PackagesPackage>
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
    suspend fun getPackagesById(id: String, licenseNumber: String? = null): PackagesPackage? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getPackagesById(id, licenseNumber, ) 
        } as? PackagesPackage
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
    suspend fun getPackagesByLabel(label: String, licenseNumber: String? = null): List<PackagesPackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getPackagesByLabel(label, licenseNumber, ) 
        } as? List<PackagesPackage>
    }

    /**
     * Retrieves a list of available Package types.
     * GET /packages/v2/types
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPackagesTypes(): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getPackagesTypes() 
        } as? Any
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
    suspend fun getSourceHarvestById(id: String, licenseNumber: String? = null): SourceHarvest? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getPackagesSourceHarvestById(id, licenseNumber, ) 
        } as? SourceHarvest
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
    suspend fun getTransferred(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<PackagesPackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.packages.getPackagesTransferred(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<PackagesPackage>
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
    suspend fun unfinishPackages(licenseNumber: String? = null, body: List<PackagesUnfinishPackagesRequestItem>): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.unfinishPackagesPackages(licenseNumber, body) 
        } as? Any
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
    suspend fun updateAdjustPackages(licenseNumber: String? = null, body: List<PackagesUpdateAdjustPackagesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updateAdjustPackages(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateDecontaminate(licenseNumber: String? = null, body: List<PackagesUpdateDecontaminateRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesDecontaminate(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateDonationFlag(licenseNumber: String? = null, body: List<PackagesUpdateDonationFlagRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesDonationFlag(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateDonationUnflag(licenseNumber: String? = null, body: List<PackagesUpdateDonationUnflagRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesDonationUnflag(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateExternalid(licenseNumber: String? = null, body: List<PackagesUpdateExternalidRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesExternalid(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateItem(licenseNumber: String? = null, body: List<PackagesUpdateItemRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesItem(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateLabtestsRequired(licenseNumber: String? = null, body: List<PackagesUpdateLabtestsRequiredRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesLabtestsRequired(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateNote(licenseNumber: String? = null, body: List<PackagesUpdateNoteRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesNote(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updatePackagesLocation(licenseNumber: String? = null, body: List<PackagesUpdatePackagesLocationRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesLocation(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateRemediate(licenseNumber: String? = null, body: List<PackagesUpdateRemediateRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesRemediate(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateTradeSampleFlag(licenseNumber: String? = null, body: List<PackagesUpdateTradeSampleFlagRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesTradeSampleFlag(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateTradeSampleUnflag(licenseNumber: String? = null, body: List<PackagesUpdateTradeSampleUnflagRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesTradeSampleUnflag(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateUseByDate(licenseNumber: String? = null, body: List<PackagesUpdateUseByDateRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.packages.updatePackagesUseByDate(licenseNumber, body) 
        } as? WriteResponse
    }
}

