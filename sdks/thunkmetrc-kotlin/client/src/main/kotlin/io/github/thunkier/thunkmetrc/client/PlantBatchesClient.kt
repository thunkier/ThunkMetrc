package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class PlantBatchesClient(private val client: MetrcClient) {
    /**
     * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * POST CreateAdjustPlantBatches
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createAdjustPlantBatches(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/adjust")
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
     * Updates the growth phase of plants at a specified Facility based on tracking information.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * POST CreateGrowthPhase
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPlantBatchesGrowthPhase(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/growthphase")
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
     * Creates packages from mother plants at the specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST CreatePackagesFromMotherPlant
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPlantBatchesPackagesFromMotherPlant(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/packages/frommotherplant")
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
     * Records Additive usage details for plant batches at a specific Facility.
     * Permissions Required:
     * - Manage Plants Additives
     * POST CreatePlantBatchesAdditives
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPlantBatchesAdditives(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/additives")
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
     * Records Additive usage for plant batches at a Facility using predefined additive templates.
     * Permissions Required:
     * - Manage Plants Additives
     * POST CreatePlantBatchesAdditivesUsingTemplate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPlantBatchesAdditivesUsingTemplate(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/additives/usingtemplate")
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
     * Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST CreatePlantBatchesPackages
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPlantBatchesPackages(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/packages")
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
     * Creates new plantings for a Facility by generating plant batches based on provided planting details.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * POST CreatePlantBatchesPlantings
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPlantBatchesPlantings(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/plantings")
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
     * Records waste information for plant batches based on the submitted data for the specified Facility.
     * Permissions Required:
     * - Manage Plants Waste
     * POST CreatePlantBatchesWaste
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPlantBatchesWaste(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/waste")
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
     * Splits an existing Plant Batch into multiple groups at the specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * POST CreateSplit
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPlantBatchesSplit(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/split")
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
     * Completes the destruction of plant batches based on the provided input data.
     * Permissions Required:
     * - View Immature Plants
     * - Destroy Immature Plants
     * DELETE DeletePlantBatches
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun deletePlantBatches(licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/plantbatches/v2")
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
     * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
     * Permissions Required:
     * - View Immature Plants
     * GET GetActivePlantBatches
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getActivePlantBatches(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/plantbatches/v2/active")
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
     * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
     * Permissions Required:
     * - View Immature Plants
     * GET GetInactivePlantBatches
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getInactivePlantBatches(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/plantbatches/v2/inactive")
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
     * Retrieves a Plant Batch by Id.
     * Permissions Required:
     * - View Immature Plants
     * GET GetPlantBatchesById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getPlantBatchesById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/plantbatches/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
     * Retrieves a list of plant batch types.
     * GET GetPlantBatchesTypes
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getPlantBatchesTypes(pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/plantbatches/v2/types")
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
     * Retrieves waste details associated with plant batches at a specified Facility.
     * Permissions Required:
     * - View Plants Waste
     * GET GetPlantBatchesWaste
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getPlantBatchesWaste(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/plantbatches/v2/waste")
        val query = ArrayList<String>()
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
     * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
     * GET GetPlantBatchesWasteReasons
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getPlantBatchesWasteReasons(licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/plantbatches/v2/waste/reasons")
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
     * Renames plant batches at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT UpdateName
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updatePlantBatchesName(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/name")
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
     * Moves one or more plant batches to new locations with in a specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants
     * PUT UpdatePlantBatchesLocation
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updatePlantBatchesLocation(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/location")
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
     * Changes the strain of plant batches at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT UpdatePlantBatchesStrain
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updatePlantBatchesStrain(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/strain")
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
     * Replaces tags for plant batches at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT UpdatePlantBatchesTag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updatePlantBatchesTag(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/tag")
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

