package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class RetailIdClient(private val client: MetrcClient) {
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
    fun createRetailIdAssociate(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/retailid/v2/associate")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun createRetailIdGenerate(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/retailid/v2/generate")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun createRetailIdMerge(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/retailid/v2/merge")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun createRetailIdPackagesInfo(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/retailid/v2/packages/info")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun getRetailIdAllotment(licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/retailid/v2/allotment")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getRetailIdReceiveByLabel(label: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/retailid/v2/receive/${URLEncoder.encode(label, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getRetailIdReceiveQrByShortCode(shortCode: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/retailid/v2/receive/qr/${URLEncoder.encode(shortCode, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

}

