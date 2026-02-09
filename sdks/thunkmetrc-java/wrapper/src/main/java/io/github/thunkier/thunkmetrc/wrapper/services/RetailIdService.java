package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class RetailIdService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public RetailIdService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

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
    public WriteResponse createRetailIdAssociate(
        String licenseNumber, List<RetailIdCreateAssociateRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.retailId.createRetailIdAssociate(
                licenseNumber, body
            )
        );
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
    @SuppressWarnings("unchecked")
    public List<WriteResponse> createRetailIdGenerate(
        String licenseNumber, RetailIdCreateGenerateRequest body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (List<WriteResponse>) client.retailId.createRetailIdGenerate(
                licenseNumber, body
            )
        );
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
    public WriteResponse createRetailIdMerge(
        String licenseNumber, RetailIdCreateMergeRequest body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.retailId.createRetailIdMerge(
                licenseNumber, body
            )
        );
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
    @SuppressWarnings("unchecked")
    public List<WriteResponse> createRetailIdPackagesInfo(
        String licenseNumber, RetailIdCreatePackagesInfoRequest body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (List<WriteResponse>) client.retailId.createRetailIdPackagesInfo(
                licenseNumber, body
            )
        );
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
    @SuppressWarnings("unchecked")
    public List<Allotment> getRetailIdAllotment(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<Allotment>) client.retailId.getRetailIdAllotment(
                licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public List<Receive> getRetailIdReceiveByLabel(
        String label, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<Receive>) client.retailId.getRetailIdReceiveByLabel(
                label, licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public List<ReceiveQrByShortCode> getRetailIdReceiveQrByShortCode(
        String shortCode, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<ReceiveQrByShortCode>) client.retailId.getRetailIdReceiveQrByShortCode(
                shortCode, licenseNumber
            )
        );
    }

}

