package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class LabTestsClient(private val client: MetrcClient) {
    /**
     * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * POST CreateRecord
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createLabTestsRecord(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v2/record")
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
     * Retrieves a list of Lab Test batches.
     * GET GetBatches
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getLabTestsBatches(pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/labtests/v2/batches")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a specific Lab Test result document by its Id for a given Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * GET GetLabTestDocumentById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getLabTestsLabTestDocumentById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/labtests/v2/labtestdocument/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
     * Returns a list of Lab Test types.
     * GET GetLabTestsTypes
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getLabTestsTypes(pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/labtests/v2/types")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves Lab Test results for a specified Package.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * GET GetResults
     * @param licenseNumber Query parameter
     * @param packageId Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getLabTestsResults(licenseNumber: String? = null, packageId: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/labtests/v2/results")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (packageId != null) {
            query.add("packageId=${URLEncoder.encode(packageId, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of all lab testing states.
     * GET GetStates
     * @throws IOException If request fails
     * @return Response object
     */
    fun getLabTestsStates(): Any? {
        val path = StringBuilder("/labtests/v2/states")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates one or more documents for previously submitted lab tests.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT UpdateLabTestDocument
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateLabTestsLabTestDocument(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v2/labtestdocument")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

    /**
     * Releases Lab Test results for one or more packages.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT UpdateResultsRelease
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateLabTestsResultsRelease(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v2/results/release")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

}

