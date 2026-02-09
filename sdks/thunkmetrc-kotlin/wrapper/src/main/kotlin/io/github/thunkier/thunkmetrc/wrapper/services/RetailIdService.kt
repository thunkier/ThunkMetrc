package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class RetailIdService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     * - WebApi Retail ID Read Write State (All or WriteOnly)
     * - Industry/View Packages
     * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
     * POST /retailid/v2/associate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createAssociate(licenseNumber: String? = null, body: List<RetailIdCreateAssociateRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.retailId.createRetailIdAssociate(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     * - WebApi Retail ID Read Write State (All or WriteOnly)
     * - Industry/View Packages
     * - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
     * POST /retailid/v2/generate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createGenerate(licenseNumber: String? = null, body: RetailIdCreateGenerateRequest): List<WriteResponse>? {
        return rateLimiter.execute(null,false,
        ) { 
            client.retailId.createRetailIdGenerate(licenseNumber, body) 
        } as? List<WriteResponse>
    }

    /**
     * Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     * - WebApi Retail ID Read Write State (All or WriteOnly)
     * - Key Value Settings/Retail ID Merge Packages Enabled
     * POST /retailid/v2/merge
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createMerge(licenseNumber: String? = null, body: RetailIdCreateMergeRequest): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.retailId.createRetailIdMerge(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Retrieves Package information for given list of Package labels.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     * - WebApi Retail ID Read Write State (All or WriteOnly)
     * - Industry/View Packages
     * - Admin/Employees/Packages Page/Product Labels(Manage)
     * POST /retailid/v2/packages/info
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createPackagesInfo(licenseNumber: String? = null, body: RetailIdCreatePackagesInfoRequest): List<WriteResponse>? {
        return rateLimiter.execute(null,false,
        ) { 
            client.retailId.createRetailIdPackagesInfo(licenseNumber, body) 
        } as? List<WriteResponse>
    }

    /**
     * Retrieves the available Retail Item ID quota for a facility.
     * Permissions Required:
     * - Download Product Labels
     * - Manage Product Labels
     * - Manage Tag Orders
     * GET /retailid/v2/allotment
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getAllotment(licenseNumber: String? = null): List<Allotment>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.retailId.getRetailIdAllotment(licenseNumber, ) 
        } as? List<Allotment>
    }

    /**
     * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Manage RetailId
     * - WebApi Retail ID Read Write State (All or ReadOnly)
     * - Industry/View Packages
     * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
     * GET /retailid/v2/receive/{label}
     * @param label Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getReceiveByLabel(label: String, licenseNumber: String? = null): List<Receive>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.retailId.getRetailIdReceiveByLabel(label, licenseNumber, ) 
        } as? List<Receive>
    }

    /**
     * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Manage RetailId
     * - WebApi Retail ID Read Write State (All or ReadOnly)
     * - Industry/View Packages
     * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
     * GET /retailid/v2/receive/qr/{shortCode}
     * @param shortCode Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getReceiveQrByShortCode(shortCode: String, licenseNumber: String? = null): List<ReceiveQrByShortCode>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.retailId.getRetailIdReceiveQrByShortCode(shortCode, licenseNumber, ) 
        } as? List<ReceiveQrByShortCode>
    }
}

