package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class RetailIdClient {
    private final MetrcClient client;

    public RetailIdClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     * - WebApi Retail ID Read Write State (All or WriteOnly)
     * - Industry/View Packages
     * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
     * POST CreateAssociate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createRetailIdAssociate(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/associate");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     * - WebApi Retail ID Read Write State (All or WriteOnly)
     * - Industry/View Packages
     * - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
     * POST CreateGenerate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createRetailIdGenerate(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/generate");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     * - WebApi Retail ID Read Write State (All or WriteOnly)
     * - Key Value Settings/Retail ID Merge Packages Enabled
     * POST CreateMerge
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createRetailIdMerge(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/merge");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Retrieves Package information for given list of Package labels.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     * - WebApi Retail ID Read Write State (All or WriteOnly)
     * - Industry/View Packages
     * - Admin/Employees/Packages Page/Product Labels(Manage)
     * POST CreatePackagesInfo
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createRetailIdPackagesInfo(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/packages/info");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Retrieves the available Retail Item ID quota for a facility.
     * Permissions Required:
     * - Download Product Labels
     * - Manage Product Labels
     * - Manage Tag Orders
     * GET GetAllotment
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getRetailIdAllotment(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/allotment");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Manage RetailId
     * - WebApi Retail ID Read Write State (All or ReadOnly)
     * - Industry/View Packages
     * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
     * GET GetReceiveByLabel
     * @param label Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getRetailIdReceiveByLabel(
        String label, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/receive/"+URLEncoder.encode(label, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Manage RetailId
     * - WebApi Retail ID Read Write State (All or ReadOnly)
     * - Industry/View Packages
     * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
     * GET GetReceiveQrByShortCode
     * @param shortCode Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getRetailIdReceiveQrByShortCode(
        String shortCode, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/receive/qr/"+URLEncoder.encode(shortCode, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

}

