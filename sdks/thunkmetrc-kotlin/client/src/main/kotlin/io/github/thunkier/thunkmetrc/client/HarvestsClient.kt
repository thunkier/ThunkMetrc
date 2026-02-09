package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class HarvestsClient(private val client: MetrcClient) {
    /**
     * Creates packages from harvested products for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST CreateHarvestsPackages
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createHarvestsPackages(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/packages")
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
     * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * POST CreateHarvestsWaste
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createHarvestsWaste(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/waste")
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
     * Creates packages for testing from harvested products for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST CreatePackagesTesting
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createHarvestsPackagesTesting(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/packages/testing")
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
     * Discontinues a specific harvest waste record by Id for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Discontinue Harvest Waste
     * DELETE DeleteWasteById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun deleteHarvestsWasteById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/harvests/v2/waste/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Marks one or more harvests as finished for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Finish/Discontinue Harvests
     * PUT FinishHarvests
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun finishHarvestsHarvests(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/finish")
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
     * Retrieves a list of active harvests for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * GET GetActiveHarvests
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getActiveHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/harvests/v2/active")
        val query = ArrayList<String>()
        if (lastModifiedEnd != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8)}")
        }
        if (lastModifiedStart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
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
     * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
     * Permissions Required:
     * - View Harvests
     * GET GetHarvestsById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getHarvestsById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/harvests/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
     * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
     * Permissions Required:
     * - View Harvests
     * GET GetHarvestsWaste
     * @param harvestId Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getHarvestsWaste(harvestId: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/harvests/v2/waste")
        val query = ArrayList<String>()
        if (harvestId != null) {
            query.add("harvestId=${URLEncoder.encode(harvestId, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
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
     * Retrieves a list of inactive harvests for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * GET GetInactiveHarvests
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getInactiveHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/harvests/v2/inactive")
        val query = ArrayList<String>()
        if (lastModifiedEnd != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8)}")
        }
        if (lastModifiedStart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
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
     * Retrieves a list of harvests on hold for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * GET GetOnHoldHarvests
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getOnHoldHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/harvests/v2/onhold")
        val query = ArrayList<String>()
        if (lastModifiedEnd != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8)}")
        }
        if (lastModifiedStart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
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
     * Retrieves a list of Waste types for harvests.
     * GET GetWasteTypes
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getHarvestsWasteTypes(pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/harvests/v2/waste/types")
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
     * Reopens one or more previously finished harvests for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Finish/Discontinue Harvests
     * PUT UnfinishHarvests
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun unfinishHarvestsHarvests(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/unfinish")
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
     * Updates the Location of Harvest for a specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * PUT UpdateHarvestsLocation
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateHarvestsLocation(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/location")
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
     * Renames one or more harvests for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Manage Harvests
     * PUT UpdateRename
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateHarvestsRename(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/rename")
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
     * Restores previously harvested plants to their original state for the specified Facility.
     * Permissions Required:
     * - View Harvests
     * - Finish/Discontinue Harvests
     * PUT UpdateRestoreHarvestedPlants
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateHarvestsRestoreHarvestedPlants(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/restore/harvestedplants")
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

