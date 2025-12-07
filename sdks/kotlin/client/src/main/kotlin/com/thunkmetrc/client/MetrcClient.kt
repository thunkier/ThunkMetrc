package com.thunkmetrc.client

import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class MetrcClient(private val baseUrl: String, vendorKey: String, userKey: String) {
    private val client = OkHttpClient()
    private val mapper = jacksonObjectMapper()
    private val auth = Credentials.basic(vendorKey, userKey)

    private inline fun <reified T> send(method: String, path: String, body: Any? = null): T? {
        val url = "${baseUrl.trimEnd('/')}$path"
        val builder = Request.Builder().url(url).header("Authorization", auth)

        if (body != null && (method == "POST" || method == "PUT")) {
            val json = mapper.writeValueAsString(body)
            builder.method(method, json.toRequestBody("application/json".toMediaType()))
        } else if (method == "POST" || method == "PUT") {
            builder.method(method, "".toRequestBody(null))
        } else {
            builder.method(method, null)
        }

        client.newCall(builder.build()).execute().use { response ->
            if (!response.isSuccessful) throw RuntimeException("Unexpected code $response")
            if (response.code == 204) return null
            return mapper.readValue(response.body!!.string(), T::class.java)
        }
    }

    /**
     * Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V1
     */
    fun plantsCreateAdditivesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/additives")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records additive usage details applied to specific plants at a Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V2
     */
    fun plantsCreateAdditivesV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/additives")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Plants
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocation V1
     */
    fun plantsCreateAdditivesBylocationV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/additives/bylocation")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records additive usage for plants based on their location within a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocation V2
     */
    fun plantsCreateAdditivesBylocationV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/additives/bylocation")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records additive usage for plants by location using a predefined additive template at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocationUsingtemplate V2
     */
    fun plantsCreateAdditivesBylocationUsingtemplateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/additives/bylocation/usingtemplate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records additive usage for plants using predefined additive templates at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesUsingtemplate V2
     */
    fun plantsCreateAdditivesUsingtemplateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/additives/usingtemplate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateChangegrowthphases V1
     */
    fun plantsCreateChangegrowthphasesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/changegrowthphases")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateHarvestplants V1
     */
    fun plantsCreateHarvestplantsV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/harvestplants")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates harvest product records from plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateManicure V2
     */
    fun plantsCreateManicureV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/manicure")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateManicureplants V1
     */
    fun plantsCreateManicureplantsV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/manicureplants")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateMoveplants V1
     */
    fun plantsCreateMoveplantsV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/moveplants")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePlantbatchPackage V1
     */
    fun plantsCreatePlantbatchPackageV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/create/plantbatch/packages")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates packages from plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePlantbatchPackage V2
     */
    fun plantsCreatePlantbatchPackageV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/plantbatch/packages")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreatePlantings V1
     */
    fun plantsCreatePlantingsV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/create/plantings")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new plant batches at a specified Facility from existing plant data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreatePlantings V2
     */
    fun plantsCreatePlantingsV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/plantings")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V1
     */
    fun plantsCreateWasteV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v1/waste")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records waste events for plants at a Facility, including method, reason, and location details.
     * 
     *   Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V2
     */
    fun plantsCreateWasteV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/waste")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Destroy Veg/Flower Plants
     *
     * DELETE Delete V1
     */
    fun plantsDeleteV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Removes plants from a Facility’s inventory while recording the reason for their disposal.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Destroy Veg/Flower Plants
     *
     * DELETE Delete V2
     */
    fun plantsDeleteV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET Get V1
     */
    fun plantsGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Plant by Id.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET Get V2
     */
    fun plantsGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View/Manage Plants Additives
     *
     * GET GetAdditives V1
     */
    fun plantsGetAdditivesV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/additives")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves additive records applied to plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View/Manage Plants Additives
     *
     * GET GetAdditives V2
     */
    fun plantsGetAdditivesV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/additives")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetAdditivesTypes V1
     */
    fun plantsGetAdditivesTypesV1(no: String? = null): Any? {
        val path = StringBuilder("/plants/v1/additives/types")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of all plant additive types defined within a Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAdditivesTypes V2
     */
    fun plantsGetAdditivesTypesV2(no: String? = null): Any? {
        val path = StringBuilder("/plants/v2/additives/types")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetByLabel V1
     */
    fun plantsGetByLabelV1(label: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/${URLEncoder.encode(label, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Plant by label.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetByLabel V2
     */
    fun plantsGetByLabelV2(label: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v2/${URLEncoder.encode(label, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetFlowering V1
     */
    fun plantsGetFloweringV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/flowering")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetFlowering V2
     */
    fun plantsGetFloweringV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/flowering")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetGrowthPhases V1
     */
    fun plantsGetGrowthPhasesV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/growthphases")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves the list of growth phases supported by a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetGrowthPhases V2
     */
    fun plantsGetGrowthPhasesV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v2/growthphases")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetInactive V1
     */
    fun plantsGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves inactive plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetInactive V2
     */
    fun plantsGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves mother-phase plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMother V2
     */
    fun plantsGetMotherV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/mother")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves inactive mother-phase plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMotherInactive V2
     */
    fun plantsGetMotherInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/mother/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMotherOnhold V2
     */
    fun plantsGetMotherOnholdV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/mother/onhold")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetOnhold V1
     */
    fun plantsGetOnholdV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/onhold")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves plants that are currently on hold at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetOnhold V2
     */
    fun plantsGetOnholdV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/onhold")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetVegetative V1
     */
    fun plantsGetVegetativeV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/vegetative")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetVegetative V2
     */
    fun plantsGetVegetativeV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/vegetative")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of recorded plant waste events for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWaste V2
     */
    fun plantsGetWasteV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/waste")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteMethodsAll V1
     */
    fun plantsGetWasteMethodsAllV1(no: String? = null): Any? {
        val path = StringBuilder("/plants/v1/waste/methods/all")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of all available plant waste methods for use within a Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteMethodsAll V2
     */
    fun plantsGetWasteMethodsAllV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/waste/methods/all")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWastePackage V2
     */
    fun plantsGetWastePackageV2(id: String, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/waste/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/package")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWastePlant V2
     */
    fun plantsGetWastePlantV2(id: String, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/waste/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/plant")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V1
     */
    fun plantsGetWasteReasonsV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plants/v1/waste/reasons")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retriveves available reasons for recording mature plant waste at a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V2
     */
    fun plantsGetWasteReasonsV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plants/v2/waste/reasons")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Adjusts the recorded count of plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateAdjust V2
     */
    fun plantsUpdateAdjustV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/adjust")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Changes the growth phases of plants within a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateGrowthphase V2
     */
    fun plantsUpdateGrowthphaseV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/growthphase")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * PUT UpdateHarvest V2
     */
    fun plantsUpdateHarvestV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/harvest")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Moves plant batches to new locations within a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateLocation V2
     */
    fun plantsUpdateLocationV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/location")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Merges multiple plant groups into a single group within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * PUT UpdateMerge V2
     */
    fun plantsUpdateMergeV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/merge")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Splits an existing plant group into multiple groups within a Facility.
     * 
     *   Permissions Required:
     *   - View Plant
     *
     * PUT UpdateSplit V2
     */
    fun plantsUpdateSplitV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/split")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates the strain information for plants within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateStrain V2
     */
    fun plantsUpdateStrainV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/strain")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Replaces existing plant tags with new tags for plants within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateTag V2
     */
    fun plantsUpdateTagV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plants/v2/tag")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     *   - WebApi Retail ID Read Write State (All or WriteOnly)
     *   - Industry/View Packages
     *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
     *
     * POST CreateAssociate V2
     */
    fun retailIdCreateAssociateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/retailid/v2/associate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     *   - WebApi Retail ID Read Write State (All or WriteOnly)
     *   - Industry/View Packages
     *   - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
     *
     * POST CreateGenerate V2
     */
    fun retailIdCreateGenerateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/retailid/v2/generate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     *   - WebApi Retail ID Read Write State (All or WriteOnly)
     *   - Key Value Settings/Retail ID Merge Packages Enabled
     *
     * POST CreateMerge V2
     */
    fun retailIdCreateMergeV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/retailid/v2/merge")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Retrieves Package information for given list of Package labels.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     *   - WebApi Retail ID Read Write State (All or WriteOnly)
     *   - Industry/View Packages
     *   - Admin/Employees/Packages Page/Product Labels(Manage)
     *
     * POST CreatePackageInfo V2
     */
    fun retailIdCreatePackageInfoV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/retailid/v2/packages/info")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Manage RetailId
     *   - WebApi Retail ID Read Write State (All or ReadOnly)
     *   - Industry/View Packages
     *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
     *
     * GET GetReceiveByLabel V2
     */
    fun retailIdGetReceiveByLabelV2(label: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/retailid/v2/receive/${URLEncoder.encode(label, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Manage RetailId
     *   - WebApi Retail ID Read Write State (All or ReadOnly)
     *   - Industry/View Packages
     *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
     *
     * GET GetReceiveQrByShortCode V2
     */
    fun retailIdGetReceiveQrByShortCodeV2(shortcode: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/retailid/v2/receive/qr/${URLEncoder.encode(shortcode, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
     * 
     *   Permissions Required:
     *   - None
     *
     * POST CreateIntegratorSetup V2
     */
    fun sandboxCreateIntegratorSetupV2(userkey: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sandbox/v2/integrator/setup")
        val query = ArrayList<String>()
        if (userkey != null) {
            query.add("userKey=${URLEncoder.encode(userkey, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V1
     */
    fun facilitiesGetAllV1(no: String? = null): Any? {
        val path = StringBuilder("/facilities/v1")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V2
     */
    fun facilitiesGetAllV2(no: String? = null): Any? {
        val path = StringBuilder("/facilities/v2")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Adds new patients to a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * POST Create V2
     */
    fun patientsCreateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patients/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * POST CreateAdd V1
     */
    fun patientsCreateAddV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patients/v1/add")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * POST CreateUpdate V1
     */
    fun patientsCreateUpdateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patients/v1/update")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * DELETE Delete V1
     */
    fun patientsDeleteV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patients/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Removes a Patient, identified by an Id, from a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * DELETE Delete V2
     */
    fun patientsDeleteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patients/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * GET Get V1
     */
    fun patientsGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patients/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Patient by Id.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * GET Get V2
     */
    fun patientsGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patients/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * GET GetActive V1
     */
    fun patientsGetActiveV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/patients/v1/active")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active patients for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * GET GetActive V2
     */
    fun patientsGetActiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/patients/v2/active")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates Patient information for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * PUT Update V2
     */
    fun patientsUpdateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patients/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * POST CreateExternalIncoming V1
     */
    fun transfersCreateExternalIncomingV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v1/external/incoming")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates external incoming shipment plans for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * POST CreateExternalIncoming V2
     */
    fun transfersCreateExternalIncomingV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/external/incoming")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * POST CreateTemplates V1
     */
    fun transfersCreateTemplatesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v1/templates")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new transfer templates for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * POST CreateTemplatesOutgoing V2
     */
    fun transfersCreateTemplatesOutgoingV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * DELETE DeleteExternalIncoming V1
     */
    fun transfersDeleteExternalIncomingV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/external/incoming/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Voids an external incoming shipment plan for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * DELETE DeleteExternalIncoming V2
     */
    fun transfersDeleteExternalIncomingV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/external/incoming/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * DELETE DeleteTemplates V1
     */
    fun transfersDeleteTemplatesV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/templates/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives a transfer template for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * DELETE DeleteTemplatesOutgoing V2
     */
    fun transfersDeleteTemplatesOutgoingV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetDeliveriesPackagesStates V1
     */
    fun transfersGetDeliveriesPackagesStatesV1(no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/deliveries/packages/states")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of available shipment Package states.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetDeliveriesPackagesStates V2
     */
    fun transfersGetDeliveriesPackagesStatesV2(no: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/packages/states")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: that the {id} parameter above represents a Shipment Plan ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDelivery V1
     */
    fun transfersGetDeliveryV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/deliveries")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDelivery V2
     */
    fun transfersGetDeliveryV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/deliveries")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackage V1
     */
    fun transfersGetDeliveryPackageV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackage V2
     */
    fun transfersGetDeliveryPackageV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackageRequiredlabtestbatches V1
     */
    fun transfersGetDeliveryPackageRequiredlabtestbatchesV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/deliveries/package/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/requiredlabtestbatches")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackageRequiredlabtestbatches V2
     */
    fun transfersGetDeliveryPackageRequiredlabtestbatchesV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/package/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/requiredlabtestbatches")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackageWholesale V1
     */
    fun transfersGetDeliveryPackageWholesaleV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages/wholesale")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackageWholesale V2
     */
    fun transfersGetDeliveryPackageWholesaleV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages/wholesale")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: that the {id} parameter above represents a Shipment Delivery ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryTransporters V1
     */
    fun transfersGetDeliveryTransportersV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryTransporters V2
     */
    fun transfersGetDeliveryTransportersV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: The {id} parameter above represents a Shipment Delivery ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryTransportersDetails V1
     */
    fun transfersGetDeliveryTransportersDetailsV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters/details")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryTransportersDetails V2
     */
    fun transfersGetDeliveryTransportersDetailsV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters/details")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetHub V2
     */
    fun transfersGetHubV2(estimatedarrivalend: String? = null, estimatedarrivalstart: String? = null, lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/hub")
        val query = ArrayList<String>()
        if (estimatedarrivalend != null) {
            query.add("estimatedArrivalEnd=${URLEncoder.encode(estimatedarrivalend, StandardCharsets.UTF_8)}")
        }
        if (estimatedarrivalstart != null) {
            query.add("estimatedArrivalStart=${URLEncoder.encode(estimatedarrivalstart, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetIncoming V1
     */
    fun transfersGetIncomingV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/incoming")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetIncoming V2
     */
    fun transfersGetIncomingV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/incoming")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetOutgoing V1
     */
    fun transfersGetOutgoingV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/outgoing")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetOutgoing V2
     */
    fun transfersGetOutgoingV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/outgoing")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetRejected V1
     */
    fun transfersGetRejectedV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/rejected")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of shipments with rejected packages for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetRejected V2
     */
    fun transfersGetRejectedV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/rejected")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplates V1
     */
    fun transfersGetTemplatesV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/templates")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDelivery V1
     */
    fun transfersGetTemplatesDeliveryV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/templates/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/deliveries")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetTemplatesDeliveryPackage V1
     */
    fun transfersGetTemplatesDeliveryPackageV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/templates/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDeliveryTransporters V1
     */
    fun transfersGetTemplatesDeliveryTransportersV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/templates/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDeliveryTransportersDetails V1
     */
    fun transfersGetTemplatesDeliveryTransportersDetailsV1(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/templates/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters/details")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoing V2
     */
    fun transfersGetTemplatesOutgoingV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of deliveries associated with a specific transfer template.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDelivery V2
     */
    fun transfersGetTemplatesOutgoingDeliveryV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/deliveries")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryPackage V2
     */
    fun transfersGetTemplatesOutgoingDeliveryPackageV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryTransporters V2
     */
    fun transfersGetTemplatesOutgoingDeliveryTransportersV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryTransportersDetails V2
     */
    fun transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id: String, no: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters/details")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     */
    fun transfersGetTypesV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/transfers/v1/types")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of available transfer types for a Facility based on its license number.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     */
    fun transfersGetTypesV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transfers/v2/types")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * PUT UpdateExternalIncoming V1
     */
    fun transfersUpdateExternalIncomingV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v1/external/incoming")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates external incoming shipment plans for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * PUT UpdateExternalIncoming V2
     */
    fun transfersUpdateExternalIncomingV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/external/incoming")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * PUT UpdateTemplates V1
     */
    fun transfersUpdateTemplatesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v1/templates")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates existing transfer templates for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * PUT UpdateTemplatesOutgoing V2
     */
    fun transfersUpdateTemplatesOutgoingV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Creates new driver records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * POST CreateDriver V2
     */
    fun transportersCreateDriverV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transporters/v2/drivers")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new vehicle records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * POST CreateVehicle V2
     */
    fun transportersCreateVehicleV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transporters/v2/vehicles")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * DELETE DeleteDriver V2
     */
    fun transportersDeleteDriverV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transporters/v2/drivers/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * DELETE DeleteVehicle V2
     */
    fun transportersDeleteVehicleV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transporters/v2/vehicles/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetDriver V2
     */
    fun transportersGetDriverV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transporters/v2/drivers/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of drivers for a Facility.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetDrivers V2
     */
    fun transportersGetDriversV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transporters/v2/drivers")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetVehicle V2
     */
    fun transportersGetVehicleV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/transporters/v2/vehicles/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of vehicles for a Facility.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetVehicles V2
     */
    fun transportersGetVehiclesV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/transporters/v2/vehicles")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates existing driver records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * PUT UpdateDriver V2
     */
    fun transportersUpdateDriverV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transporters/v2/drivers")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates existing vehicle records for a facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * PUT UpdateVehicle V2
     */
    fun transportersUpdateVehicleV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transporters/v2/vehicles")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetActive V1
     */
    fun unitsOfMeasureGetActiveV1(no: String? = null): Any? {
        val path = StringBuilder("/unitsofmeasure/v1/active")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves all active units of measure.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetActive V2
     */
    fun unitsOfMeasureGetActiveV2(no: String? = null): Any? {
        val path = StringBuilder("/unitsofmeasure/v2/active")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves all inactive units of measure.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetInactive V2
     */
    fun unitsOfMeasureGetInactiveV2(no: String? = null): Any? {
        val path = StringBuilder("/unitsofmeasure/v2/inactive")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves all available waste methods.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V2
     */
    fun wasteMethodsGetAllV2(no: String? = null): Any? {
        val path = StringBuilder("/wastemethods/v2")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Caregivers
     *
     * GET GetByCaregiverLicenseNumber V1
     */
    fun caregiversStatusGetByCaregiverLicenseNumberV1(caregiverlicensenumber: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/caregivers/v1/status/${URLEncoder.encode(caregiverlicensenumber, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Caregivers
     *
     * GET GetByCaregiverLicenseNumber V2
     */
    fun caregiversStatusGetByCaregiverLicenseNumberV2(caregiverlicensenumber: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/caregivers/v2/status/${URLEncoder.encode(caregiverlicensenumber, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Employees
     *
     * GET GetAll V1
     */
    fun employeesGetAllV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/employees/v1")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of employees for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Employees
     *   - View Employees
     *
     * GET GetAll V2
     */
    fun employeesGetAllV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/employees/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
     * 
     *   Permissions Required:
     *   - Manage Employees
     *
     * GET GetPermissions V2
     */
    fun employeesGetPermissionsV2(employeelicensenumber: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/employees/v2/permissions")
        val query = ArrayList<String>()
        if (employeelicensenumber != null) {
            query.add("employeeLicenseNumber=${URLEncoder.encode(employeelicensenumber, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST Create V1
     */
    fun itemsCreateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v1/create")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST Create V2
     */
    fun itemsCreateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates one or more new item brands for the specified Facility identified by the License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreateBrand V2
     */
    fun itemsCreateBrandV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v2/brand")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreateFile V2
     */
    fun itemsCreateFileV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v2/file")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreatePhoto V1
     */
    fun itemsCreatePhotoV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v1/photo")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreatePhoto V2
     */
    fun itemsCreatePhotoV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v2/photo")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * POST CreateUpdate V1
     */
    fun itemsCreateUpdateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v1/update")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * DELETE Delete V1
     */
    fun itemsDeleteV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives the specified Product by Id for the given Facility License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * DELETE Delete V2
     */
    fun itemsDeleteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives the specified Item Brand by Id for the given Facility License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * DELETE DeleteBrand V2
     */
    fun itemsDeleteBrandV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v2/brand/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET Get V1
     */
    fun itemsGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves detailed information about a specific Item by Id.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET Get V2
     */
    fun itemsGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetActive V1
     */
    fun itemsGetActiveV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v1/active")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of active items for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetActive V2
     */
    fun itemsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/items/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetBrands V1
     */
    fun itemsGetBrandsV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v1/brands")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active item brands for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetBrands V2
     */
    fun itemsGetBrandsV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/items/v2/brands")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCategories V1
     */
    fun itemsGetCategoriesV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v1/categories")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of item categories.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCategories V2
     */
    fun itemsGetCategoriesV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/items/v2/categories")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a file by its Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetFile V2
     */
    fun itemsGetFileV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v2/file/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetInactive V1
     */
    fun itemsGetInactiveV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v1/inactive")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of inactive items for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetInactive V2
     */
    fun itemsGetInactiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/items/v2/inactive")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetPhoto V1
     */
    fun itemsGetPhotoV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v1/photo/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves an image by its Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetPhoto V2
     */
    fun itemsGetPhotoV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/items/v2/photo/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates one or more existing products for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * PUT Update V2
     */
    fun itemsUpdateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates one or more existing item brands for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * PUT UpdateBrand V2
     */
    fun itemsUpdateBrandV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/items/v2/brand")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Create V1
     */
    fun packagesCreateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/create")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Create V2
     */
    fun packagesCreateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateAdjust V1
     */
    fun packagesCreateAdjustV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/adjust")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records a list of adjustments for packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateAdjust V2
     */
    fun packagesCreateAdjustV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/adjust")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateChangeItem V1
     */
    fun packagesCreateChangeItemV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/change/item")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateChangeLocation V1
     */
    fun packagesCreateChangeLocationV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/change/locations")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateFinish V1
     */
    fun packagesCreateFinishV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/finish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreatePlantings V1
     */
    fun packagesCreatePlantingsV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/create/plantings")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new plantings from packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreatePlantings V2
     */
    fun packagesCreatePlantingsV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/plantings")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRemediate V1
     */
    fun packagesCreateRemediateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/remediate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateTesting V1
     */
    fun packagesCreateTestingV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/create/testing")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new packages for testing for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateTesting V2
     */
    fun packagesCreateTestingV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/testing")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateUnfinish V1
     */
    fun packagesCreateUnfinishV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/unfinish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Discontinues a Package at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * DELETE Delete V2
     */
    fun packagesDeleteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET Get V1
     */
    fun packagesGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Package by its Id.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET Get V2
     */
    fun packagesGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetActive V1
     */
    fun packagesGetActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v1/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetActive V2
     */
    fun packagesGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/packages/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetAdjustReasons V1
     */
    fun packagesGetAdjustReasonsV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v1/adjust/reasons")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of adjustment reasons for packages at a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAdjustReasons V2
     */
    fun packagesGetAdjustReasonsV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/packages/v2/adjust/reasons")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetByLabel V1
     */
    fun packagesGetByLabelV1(label: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v1/${URLEncoder.encode(label, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Package by its label.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetByLabel V2
     */
    fun packagesGetByLabelV2(label: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v2/${URLEncoder.encode(label, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetInactive V1
     */
    fun packagesGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v1/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of inactive packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetInactive V2
     */
    fun packagesGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/packages/v2/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of packages in transit for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetIntransit V2
     */
    fun packagesGetIntransitV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/packages/v2/intransit")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetLabsamples V2
     */
    fun packagesGetLabsamplesV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/packages/v2/labsamples")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetOnhold V1
     */
    fun packagesGetOnholdV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v1/onhold")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of packages on hold for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetOnhold V2
     */
    fun packagesGetOnholdV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/packages/v2/onhold")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves the source harvests for a Package by its Id.
     * 
     *   Permissions Required:
     *   - View Package Source Harvests
     *
     * GET GetSourceHarvest V2
     */
    fun packagesGetSourceHarvestV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/packages/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/source/harvests")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of transferred packages for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetTransferred V2
     */
    fun packagesGetTransferredV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/packages/v2/transferred")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     */
    fun packagesGetTypesV1(no: String? = null): Any? {
        val path = StringBuilder("/packages/v1/types")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of available Package types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     */
    fun packagesGetTypesV2(no: String? = null): Any? {
        val path = StringBuilder("/packages/v2/types")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Set the final quantity for a Package.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateAdjust V2
     */
    fun packagesUpdateAdjustV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/adjust")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *   - Manage Package Notes
     *
     * PUT UpdateChangeNote V1
     */
    fun packagesUpdateChangeNoteV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v1/change/note")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates the Product decontaminate information for a list of packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDecontaminate V2
     */
    fun packagesUpdateDecontaminateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/decontaminate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Flags one or more packages for donation at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDonationFlag V2
     */
    fun packagesUpdateDonationFlagV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/donation/flag")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Removes the donation flag from one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDonationUnflag V2
     */
    fun packagesUpdateDonationUnflagV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/donation/unflag")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates the external identifiers for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Package Inventory
     *   - External Id Enabled
     *
     * PUT UpdateExternalid V2
     */
    fun packagesUpdateExternalidV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/externalid")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates a list of packages as finished for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateFinish V2
     */
    fun packagesUpdateFinishV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/finish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates the associated Item for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateItem V2
     */
    fun packagesUpdateItemV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/item")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates the list of required lab test batches for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateLabTestRequired V2
     */
    fun packagesUpdateLabTestRequiredV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/labtests/required")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates the Location and Sublocation for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateLocation V2
     */
    fun packagesUpdateLocationV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/location")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates notes associated with one or more packages for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *   - Manage Package Notes
     *
     * PUT UpdateNote V2
     */
    fun packagesUpdateNoteV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/note")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates a list of Product remediations for packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateRemediate V2
     */
    fun packagesUpdateRemediateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/remediate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Flags or unflags one or more packages at the specified Facility as trade samples.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateTradesampleFlag V2
     */
    fun packagesUpdateTradesampleFlagV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/tradesample/flag")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Removes the trade sample flag from one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateTradesampleUnflag V2
     */
    fun packagesUpdateTradesampleUnflagV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/tradesample/unflag")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates a list of packages as unfinished for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateUnfinish V2
     */
    fun packagesUpdateUnfinishV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/unfinish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates the use-by date for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateUsebydate V2
     */
    fun packagesUpdateUsebydateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/packages/v2/usebydate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * POST Create V1
     */
    fun patientCheckInsCreateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patient-checkins/v1")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * POST Create V2
     */
    fun patientCheckInsCreateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patient-checkins/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * DELETE Delete V1
     */
    fun patientCheckInsDeleteV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patient-checkins/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives a Patient Check-In, identified by its Id, for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * DELETE Delete V2
     */
    fun patientCheckInsDeleteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patient-checkins/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * GET GetAll V1
     */
    fun patientCheckInsGetAllV1(checkindateend: String? = null, checkindatestart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patient-checkins/v1")
        val query = ArrayList<String>()
        if (checkindateend != null) {
            query.add("checkinDateEnd=${URLEncoder.encode(checkindateend, StandardCharsets.UTF_8)}")
        }
        if (checkindatestart != null) {
            query.add("checkinDateStart=${URLEncoder.encode(checkindatestart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * GET GetAll V2
     */
    fun patientCheckInsGetAllV2(checkindateend: String? = null, checkindatestart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patient-checkins/v2")
        val query = ArrayList<String>()
        if (checkindateend != null) {
            query.add("checkinDateEnd=${URLEncoder.encode(checkindateend, StandardCharsets.UTF_8)}")
        }
        if (checkindatestart != null) {
            query.add("checkinDateStart=${URLEncoder.encode(checkindatestart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetLocations V1
     */
    fun patientCheckInsGetLocationsV1(no: String? = null): Any? {
        val path = StringBuilder("/patient-checkins/v1/locations")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of Patient Check-In locations.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetLocations V2
     */
    fun patientCheckInsGetLocationsV2(no: String? = null): Any? {
        val path = StringBuilder("/patient-checkins/v2/locations")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * PUT Update V1
     */
    fun patientCheckInsUpdateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patient-checkins/v1")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * PUT Update V2
     */
    fun patientCheckInsUpdateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patient-checkins/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V1
     */
    fun plantBatchesCreateAdditivesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/additives")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records Additive usage details for plant batches at a specific Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V2
     */
    fun plantBatchesCreateAdditivesV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/additives")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records Additive usage for plant batches at a Facility using predefined additive templates.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesUsingtemplate V2
     */
    fun plantBatchesCreateAdditivesUsingtemplateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/additives/usingtemplate")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateAdjust V1
     */
    fun plantBatchesCreateAdjustV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/adjust")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateAdjust V2
     */
    fun plantBatchesCreateAdjustV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/adjust")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateChangegrowthphase V1
     */
    fun plantBatchesCreateChangegrowthphaseV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/changegrowthphase")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Updates the growth phase of plants at a specified Facility based on tracking information.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateGrowthphase V2
     */
    fun plantBatchesCreateGrowthphaseV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/growthphase")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackage V2
     */
    fun plantBatchesCreatePackageV2(isfrommotherplant: String? = null, licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/packages")
        val query = ArrayList<String>()
        if (isfrommotherplant != null) {
            query.add("isFromMotherPlant=${URLEncoder.encode(isfrommotherplant, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageFrommotherplant V1
     */
    fun plantBatchesCreatePackageFrommotherplantV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/create/packages/frommotherplant")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates packages from mother plants at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageFrommotherplant V2
     */
    fun plantBatchesCreatePackageFrommotherplantV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/packages/frommotherplant")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new plantings for a Facility by generating plant batches based on provided planting details.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreatePlantings V2
     */
    fun plantBatchesCreatePlantingsV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/plantings")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateSplit V1
     */
    fun plantBatchesCreateSplitV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/split")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Splits an existing Plant Batch into multiple groups at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateSplit V2
     */
    fun plantBatchesCreateSplitV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/split")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V1
     */
    fun plantBatchesCreateWasteV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/waste")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records waste information for plant batches based on the submitted data for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V2
     */
    fun plantBatchesCreateWasteV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/waste")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Createpackages V1
     */
    fun plantBatchesCreatepackagesV1(isfrommotherplant: String? = null, licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/createpackages")
        val query = ArrayList<String>()
        if (isfrommotherplant != null) {
            query.add("isFromMotherPlant=${URLEncoder.encode(isfrommotherplant, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST Createplantings V1
     */
    fun plantBatchesCreateplantingsV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/createplantings")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Destroy Immature Plants
     *
     * DELETE Delete V1
     */
    fun plantBatchesDeleteV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v1")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Completes the destruction of plant batches based on the provided input data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Destroy Immature Plants
     *
     * DELETE Delete V2
     */
    fun plantBatchesDeleteV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET Get V1
     */
    fun plantBatchesGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Plant Batch by Id.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET Get V2
     */
    fun plantBatchesGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET GetActive V1
     */
    fun plantBatchesGetActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET GetActive V2
     */
    fun plantBatchesGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET GetInactive V1
     */
    fun plantBatchesGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET GetInactive V2
     */
    fun plantBatchesGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     */
    fun plantBatchesGetTypesV1(no: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/types")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of plant batch types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     */
    fun plantBatchesGetTypesV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/types")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves waste details associated with plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWaste V2
     */
    fun plantBatchesGetWasteV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/waste")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V1
     */
    fun plantBatchesGetWasteReasonsV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/waste/reasons")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V2
     */
    fun plantBatchesGetWasteReasonsV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/waste/reasons")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Moves one or more plant batches to new locations with in a specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *
     * PUT UpdateLocation V2
     */
    fun plantBatchesUpdateLocationV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/location")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * PUT UpdateMoveplantbatches V1
     */
    fun plantBatchesUpdateMoveplantbatchesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v1/moveplantbatches")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Renames plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateName V2
     */
    fun plantBatchesUpdateNameV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/name")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Changes the strain of plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateStrain V2
     */
    fun plantBatchesUpdateStrainV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/strain")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Replaces tags for plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateTag V2
     */
    fun plantBatchesUpdateTagV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/plantbatches/v2/tag")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST CreateAdjust V1
     */
    fun processingJobsCreateAdjustV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v1/adjust")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateAdjust V2
     */
    fun processingJobsCreateAdjustV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v2/adjust")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateJobtypes V1
     */
    fun processingJobsCreateJobtypesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v1/jobtypes")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateJobtypes V2
     */
    fun processingJobsCreateJobtypesV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v2/jobtypes")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST CreateStart V1
     */
    fun processingJobsCreateStartV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v1/start")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Initiates new processing jobs at a Facility, including job details and associated packages.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateStart V2
     */
    fun processingJobsCreateStartV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v2/start")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST Createpackages V1
     */
    fun processingJobsCreatepackagesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v1/createpackages")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates packages from processing jobs at a Facility, including optional location and note assignments.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST Createpackages V2
     */
    fun processingJobsCreatepackagesV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v2/createpackages")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE Delete V1
     */
    fun processingJobsDeleteV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE Delete V2
     */
    fun processingJobsDeleteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE DeleteJobtypes V1
     */
    fun processingJobsDeleteJobtypesV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/jobtypes/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives a Processing Job Type at a Facility, making it inactive for future use.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE DeleteJobtypes V2
     */
    fun processingJobsDeleteJobtypesV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v2/jobtypes/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET Get V1
     */
    fun processingJobsGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a ProcessingJob by Id.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET Get V2
     */
    fun processingJobsGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetActive V1
     */
    fun processingJobsGetActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves active processing jobs at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetActive V2
     */
    fun processingJobsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/processing/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetInactive V1
     */
    fun processingJobsGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves inactive processing jobs at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetInactive V2
     */
    fun processingJobsGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/processing/v2/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesActive V1
     */
    fun processingJobsGetJobtypesActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/jobtypes/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of all active processing job types defined within a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesActive V2
     */
    fun processingJobsGetJobtypesActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/processing/v2/jobtypes/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesAttributes V1
     */
    fun processingJobsGetJobtypesAttributesV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/jobtypes/attributes")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves all processing job attributes available for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesAttributes V2
     */
    fun processingJobsGetJobtypesAttributesV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v2/jobtypes/attributes")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesCategories V1
     */
    fun processingJobsGetJobtypesCategoriesV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/jobtypes/categories")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves all processing job categories available for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesCategories V2
     */
    fun processingJobsGetJobtypesCategoriesV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v2/jobtypes/categories")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesInactive V1
     */
    fun processingJobsGetJobtypesInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/processing/v1/jobtypes/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of all inactive processing job types defined within a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesInactive V2
     */
    fun processingJobsGetJobtypesInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/processing/v2/jobtypes/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateFinish V1
     */
    fun processingJobsUpdateFinishV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v1/finish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Completes processing jobs at a Facility by recording final notes and waste measurements.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateFinish V2
     */
    fun processingJobsUpdateFinishV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v2/finish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateJobtypes V1
     */
    fun processingJobsUpdateJobtypesV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v1/jobtypes")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateJobtypes V2
     */
    fun processingJobsUpdateJobtypesV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v2/jobtypes")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateUnfinish V1
     */
    fun processingJobsUpdateUnfinishV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v1/unfinish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateUnfinish V2
     */
    fun processingJobsUpdateUnfinishV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/processing/v2/unfinish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Creates new sublocation records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * POST Create V2
     */
    fun sublocationsCreateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sublocations/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Archives an existing Sublocation record for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V2
     */
    fun sublocationsDeleteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sublocations/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Retrieves a Sublocation by its Id, with an optional license number.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET Get V2
     */
    fun sublocationsGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sublocations/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V2
     */
    fun sublocationsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/sublocations/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of inactive sublocations for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetInactive V2
     */
    fun sublocationsGetInactiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/sublocations/v2/inactive")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates existing sublocation records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * PUT Update V2
     */
    fun sublocationsUpdateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sublocations/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * POST CreateDelivery V1
     */
    fun salesCreateDeliveryV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *
     * POST CreateDelivery V2
     */
    fun salesCreateDeliveryV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailer V1
     */
    fun salesCreateDeliveryRetailerV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *   - Manage Retailer Delivery
     *
     * POST CreateDeliveryRetailer V2
     */
    fun salesCreateDeliveryRetailerV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerDepart V1
     */
    fun salesCreateDeliveryRetailerDepartV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer/depart")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - Manage Retailer Delivery
     *
     * POST CreateDeliveryRetailerDepart V2
     */
    fun salesCreateDeliveryRetailerDepartV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/depart")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerEnd V1
     */
    fun salesCreateDeliveryRetailerEndV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer/end")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - Manage Retailer Delivery
     *
     * POST CreateDeliveryRetailerEnd V2
     */
    fun salesCreateDeliveryRetailerEndV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/end")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerRestock V1
     */
    fun salesCreateDeliveryRetailerRestockV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer/restock")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - Manage Retailer Delivery
     *
     * POST CreateDeliveryRetailerRestock V2
     */
    fun salesCreateDeliveryRetailerRestockV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/restock")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerSale V1
     */
    fun salesCreateDeliveryRetailerSaleV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer/sale")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records sales deliveries originating from a retailer delivery for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *
     * POST CreateDeliveryRetailerSale V2
     */
    fun salesCreateDeliveryRetailerSaleV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/sale")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales
     *
     * POST CreateReceipt V1
     */
    fun salesCreateReceiptV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/receipts")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales (Write)
     *   - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
     *   - Industry/Facility Type/Advanced Sales
     *   - WebApi Sales Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *
     * POST CreateReceipt V2
     */
    fun salesCreateReceiptV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * POST CreateTransactionByDate V1
     */
    fun salesCreateTransactionByDateV1(date: String, licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/transactions/${URLEncoder.encode(date, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * DELETE DeleteDelivery V1
     */
    fun salesDeleteDeliveryV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * DELETE DeleteDelivery V2
     */
    fun salesDeleteDeliveryV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * DELETE DeleteDeliveryRetailer V1
     */
    fun salesDeleteDeliveryRetailerV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - Manage Retailer Delivery
     *
     * DELETE DeleteDeliveryRetailer V2
     */
    fun salesDeleteDeliveryRetailerV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * DELETE DeleteReceipt V1
     */
    fun salesDeleteReceiptV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/receipts/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * DELETE DeleteReceipt V2
     */
    fun salesDeleteReceiptV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCounties V1
     */
    fun salesGetCountiesV1(no: String? = null): Any? {
        val path = StringBuilder("/sales/v1/counties")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of counties available for sales deliveries.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCounties V2
     */
    fun salesGetCountiesV2(no: String? = null): Any? {
        val path = StringBuilder("/sales/v2/counties")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCustomertypes V1
     */
    fun salesGetCustomertypesV1(no: String? = null): Any? {
        val path = StringBuilder("/sales/v1/customertypes")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of customer types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCustomertypes V2
     */
    fun salesGetCustomertypesV2(no: String? = null): Any? {
        val path = StringBuilder("/sales/v2/customertypes")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesActive V1
     */
    fun salesGetDeliveriesActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (salesdateend != null) {
            query.add("salesDateEnd=${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        }
        if (salesdatestart != null) {
            query.add("salesDateStart=${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDeliveriesActive V2
     */
    fun salesGetDeliveriesActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (salesdateend != null) {
            query.add("salesDateEnd=${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        }
        if (salesdatestart != null) {
            query.add("salesDateStart=${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesInactive V1
     */
    fun salesGetDeliveriesInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (salesdateend != null) {
            query.add("salesDateEnd=${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        }
        if (salesdatestart != null) {
            query.add("salesDateStart=${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDeliveriesInactive V2
     */
    fun salesGetDeliveriesInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (salesdateend != null) {
            query.add("salesDateEnd=${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        }
        if (salesdatestart != null) {
            query.add("salesDateStart=${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveriesRetailerActive V1
     */
    fun salesGetDeliveriesRetailerActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveriesRetailerActive V2
     */
    fun salesGetDeliveriesRetailerActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveriesRetailerInactive V1
     */
    fun salesGetDeliveriesRetailerInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveriesRetailerInactive V2
     */
    fun salesGetDeliveriesRetailerInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetDeliveriesReturnreasons V1
     */
    fun salesGetDeliveriesReturnreasonsV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/returnreasons")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of return reasons for sales deliveries based on the provided License Number.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesReturnreasons V2
     */
    fun salesGetDeliveriesReturnreasonsV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/returnreasons")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDelivery V1
     */
    fun salesGetDeliveryV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a sales delivery record by its Id, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDelivery V2
     */
    fun salesGetDeliveryV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveryRetailer V1
     */
    fun salesGetDeliveryRetailerV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a retailer delivery record by its ID, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveryRetailer V2
     */
    fun salesGetDeliveryRetailerV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetPatientRegistrationsLocations V1
     */
    fun salesGetPatientRegistrationsLocationsV1(no: String? = null): Any? {
        val path = StringBuilder("/sales/v1/patientregistration/locations")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of valid Patient registration locations for sales.
     * 
     *   Permissions Required:
     *   -
     *
     * GET GetPatientRegistrationsLocations V2
     */
    fun salesGetPatientRegistrationsLocationsV2(no: String? = null): Any? {
        val path = StringBuilder("/sales/v2/patientregistration/locations")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetPaymenttypes V1
     */
    fun salesGetPaymenttypesV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/paymenttypes")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of available payment types for the specified License Number.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetPaymenttypes V2
     */
    fun salesGetPaymenttypesV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v2/paymenttypes")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceipt V1
     */
    fun salesGetReceiptV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/receipts/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a sales receipt by its Id, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceipt V2
     */
    fun salesGetReceiptV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceiptsActive V1
     */
    fun salesGetReceiptsActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        val path = StringBuilder("/sales/v1/receipts/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (salesdateend != null) {
            query.add("salesDateEnd=${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        }
        if (salesdatestart != null) {
            query.add("salesDateStart=${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsActive V2
     */
    fun salesGetReceiptsActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (salesdateend != null) {
            query.add("salesDateEnd=${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        }
        if (salesdatestart != null) {
            query.add("salesDateStart=${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Sales Receipt by its external number, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsExternalByExternalNumber V2
     */
    fun salesGetReceiptsExternalByExternalNumberV2(externalnumber: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/external/${URLEncoder.encode(externalnumber, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceiptsInactive V1
     */
    fun salesGetReceiptsInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        val path = StringBuilder("/sales/v1/receipts/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (salesdateend != null) {
            query.add("salesDateEnd=${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        }
        if (salesdatestart != null) {
            query.add("salesDateStart=${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsInactive V2
     */
    fun salesGetReceiptsInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (salesdateend != null) {
            query.add("salesDateEnd=${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        }
        if (salesdatestart != null) {
            query.add("salesDateStart=${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetTransactions V1
     */
    fun salesGetTransactionsV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/transactions")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
     */
    fun salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesdatestart: String, salesdateend: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/sales/v1/transactions/${URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)}/${URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDelivery V1
     */
    fun salesUpdateDeliveryV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * PUT UpdateDelivery V2
     */
    fun salesUpdateDeliveryV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDeliveryComplete V1
     */
    fun salesUpdateDeliveryCompleteV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/complete")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * PUT UpdateDeliveryComplete V2
     */
    fun salesUpdateDeliveryCompleteV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/complete")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDeliveryHub V1
     */
    fun salesUpdateDeliveryHubV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/hub")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery, Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHub V2
     */
    fun salesUpdateDeliveryHubV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/hub")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubAccept V1
     */
    fun salesUpdateDeliveryHubAcceptV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/hub/accept")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubAccept V2
     */
    fun salesUpdateDeliveryHubAcceptV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/hub/accept")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubDepart V1
     */
    fun salesUpdateDeliveryHubDepartV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/hub/depart")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubDepart V2
     */
    fun salesUpdateDeliveryHubDepartV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/hub/depart")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubVerifyID V1
     */
    fun salesUpdateDeliveryHubVerifyIdV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/hub/verifyID")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubVerifyID V2
     */
    fun salesUpdateDeliveryHubVerifyIdV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/hub/verifyID")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * PUT UpdateDeliveryRetailer V1
     */
    fun salesUpdateDeliveryRetailerV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/deliveries/retailer")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *   - Manage Retailer Delivery
     *
     * PUT UpdateDeliveryRetailer V2
     */
    fun salesUpdateDeliveryRetailerV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales
     *
     * PUT UpdateReceipt V1
     */
    fun salesUpdateReceiptV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/receipts")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceipt V2
     */
    fun salesUpdateReceiptV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceiptFinalize V2
     */
    fun salesUpdateReceiptFinalizeV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/finalize")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceiptUnfinalize V2
     */
    fun salesUpdateReceiptUnfinalizeV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/unfinalize")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateTransactionByDate V1
     */
    fun salesUpdateTransactionByDateV1(date: String, licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v1/transactions/${URLEncoder.encode(date, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * POST Create V1
     */
    fun strainsCreateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/strains/v1/create")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new strain records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * POST Create V2
     */
    fun strainsCreateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/strains/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * POST CreateUpdate V1
     */
    fun strainsCreateUpdateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/strains/v1/update")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * DELETE Delete V1
     */
    fun strainsDeleteV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/strains/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives an existing strain record for a Facility
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * DELETE Delete V2
     */
    fun strainsDeleteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/strains/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * GET Get V1
     */
    fun strainsGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/strains/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Strain record by its Id, with an optional license number.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET Get V2
     */
    fun strainsGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/strains/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * GET GetActive V1
     */
    fun strainsGetActiveV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/strains/v1/active")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET GetActive V2
     */
    fun strainsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/strains/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET GetInactive V2
     */
    fun strainsGetInactiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/strains/v2/inactive")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates existing strain records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * PUT Update V2
     */
    fun strainsUpdateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/strains/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Returns a list of available package tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *
     * GET GetPackageAvailable V2
     */
    fun tagsGetPackageAvailableV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/tags/v2/package/available")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of available plant tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *
     * GET GetPlantAvailable V2
     */
    fun tagsGetPlantAvailableV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/tags/v2/plant/available")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of staged tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *   - RetailId.AllowPackageStaging Key Value enabled
     *
     * GET GetStaged V2
     */
    fun tagsGetStagedV2(licensenumber: String? = null): Any? {
        val path = StringBuilder("/tags/v2/staged")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Creates new additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * POST Create V2
     */
    fun additivesTemplatesCreateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/additivestemplates/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Retrieves an Additive Template by its Id.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET Get V2
     */
    fun additivesTemplatesGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/additivestemplates/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET GetActive V2
     */
    fun additivesTemplatesGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/additivestemplates/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of inactive additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET GetInactive V2
     */
    fun additivesTemplatesGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/additivestemplates/v2/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates existing additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * PUT Update V2
     */
    fun additivesTemplatesUpdateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/additivestemplates/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * POST CreateFinish V1
     */
    fun harvestsCreateFinishV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v1/finish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackage V1
     */
    fun harvestsCreatePackageV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v1/create/packages")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates packages from harvested products for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackage V2
     */
    fun harvestsCreatePackageV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/packages")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageTesting V1
     */
    fun harvestsCreatePackageTestingV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v1/create/packages/testing")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates packages for testing from harvested products for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageTesting V2
     */
    fun harvestsCreatePackageTestingV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/packages/testing")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * POST CreateRemoveWaste V1
     */
    fun harvestsCreateRemoveWasteV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v1/removewaste")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * POST CreateUnfinish V1
     */
    fun harvestsCreateUnfinishV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v1/unfinish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * POST CreateWaste V2
     */
    fun harvestsCreateWasteV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/waste")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Discontinues a specific harvest waste record by Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Discontinue Harvest Waste
     *
     * DELETE DeleteWaste V2
     */
    fun harvestsDeleteWasteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/harvests/v2/waste/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET Get V1
     */
    fun harvestsGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/harvests/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET Get V2
     */
    fun harvestsGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/harvests/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetActive V1
     */
    fun harvestsGetActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/harvests/v1/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active harvests for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetActive V2
     */
    fun harvestsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/harvests/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetInactive V1
     */
    fun harvestsGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/harvests/v1/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of inactive harvests for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetInactive V2
     */
    fun harvestsGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/harvests/v2/inactive")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetOnhold V1
     */
    fun harvestsGetOnholdV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        val path = StringBuilder("/harvests/v1/onhold")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of harvests on hold for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetOnhold V2
     */
    fun harvestsGetOnholdV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/harvests/v2/onhold")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetWaste V2
     */
    fun harvestsGetWasteV2(harvestid: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/harvests/v2/waste")
        val query = ArrayList<String>()
        if (harvestid != null) {
            query.add("harvestId=${URLEncoder.encode(harvestid, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteTypes V1
     */
    fun harvestsGetWasteTypesV1(no: String? = null): Any? {
        val path = StringBuilder("/harvests/v1/waste/types")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of Waste types for harvests.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteTypes V2
     */
    fun harvestsGetWasteTypesV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/harvests/v2/waste/types")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Marks one or more harvests as finished for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateFinish V2
     */
    fun harvestsUpdateFinishV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/finish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates the Location of Harvest for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateLocation V2
     */
    fun harvestsUpdateLocationV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/location")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateMove V1
     */
    fun harvestsUpdateMoveV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v1/move")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateRename V1
     */
    fun harvestsUpdateRenameV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v1/rename")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Renames one or more harvests for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateRename V2
     */
    fun harvestsUpdateRenameV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/rename")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Restores previously harvested plants to their original state for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateRestoreHarvestedPlants V2
     */
    fun harvestsUpdateRestoreHarvestedPlantsV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/restore/harvestedplants")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Reopens one or more previously finished harvests for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateUnfinish V2
     */
    fun harvestsUpdateUnfinishV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/harvests/v2/unfinish")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRecord V1
     */
    fun labTestsCreateRecordV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v1/record")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRecord V2
     */
    fun labTestsCreateRecordV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v2/record")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Retrieves a list of Lab Test batches.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetBatches V2
     */
    fun labTestsGetBatchesV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/labtests/v2/batches")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetLabtestdocument V1
     */
    fun labTestsGetLabtestdocumentV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/labtests/v1/labtestdocument/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a specific Lab Test result document by its Id for a given Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetLabtestdocument V2
     */
    fun labTestsGetLabtestdocumentV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/labtests/v2/labtestdocument/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetResults V1
     */
    fun labTestsGetResultsV1(licensenumber: String? = null, packageid: String? = null): Any? {
        val path = StringBuilder("/labtests/v1/results")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (packageid != null) {
            query.add("packageId=${URLEncoder.encode(packageid, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves Lab Test results for a specified Package.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetResults V2
     */
    fun labTestsGetResultsV2(licensenumber: String? = null, packageid: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/labtests/v2/results")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (packageid != null) {
            query.add("packageId=${URLEncoder.encode(packageid, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetStates V1
     */
    fun labTestsGetStatesV1(no: String? = null): Any? {
        val path = StringBuilder("/labtests/v1/states")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of all lab testing states.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetStates V2
     */
    fun labTestsGetStatesV2(no: String? = null): Any? {
        val path = StringBuilder("/labtests/v2/states")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     */
    fun labTestsGetTypesV1(no: String? = null): Any? {
        val path = StringBuilder("/labtests/v1/types")
        val query = ArrayList<String>()
        if (no != null) {
            query.add("No=${URLEncoder.encode(no, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of Lab Test types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     */
    fun labTestsGetTypesV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/labtests/v2/types")
        val query = ArrayList<String>()
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateLabtestdocument V1
     */
    fun labTestsUpdateLabtestdocumentV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v1/labtestdocument")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates one or more documents for previously submitted lab tests.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateLabtestdocument V2
     */
    fun labTestsUpdateLabtestdocumentV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v2/labtestdocument")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateResultRelease V1
     */
    fun labTestsUpdateResultReleaseV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v1/results/release")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Releases Lab Test results for one or more packages.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateResultRelease V2
     */
    fun labTestsUpdateResultReleaseV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/labtests/v2/results/release")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * POST Create V1
     */
    fun locationsCreateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/locations/v1/create")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * POST Create V2
     */
    fun locationsCreateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/locations/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * POST CreateUpdate V1
     */
    fun locationsCreateUpdateV1(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/locations/v1/update")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("POST", path.toString(), body)
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V1
     */
    fun locationsDeleteV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/locations/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives a specified Location, identified by its Id, for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V2
     */
    fun locationsDeleteV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/locations/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET Get V1
     */
    fun locationsGetV1(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/locations/v1/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Location by its Id.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET Get V2
     */
    fun locationsGetV2(id: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/locations/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V1
     */
    fun locationsGetActiveV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/locations/v1/active")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V2
     */
    fun locationsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/locations/v2/active")
        val query = ArrayList<String>()
        if (lastmodifiedend != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8)}")
        }
        if (lastmodifiedstart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8)}")
        }
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of inactive locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetInactive V2
     */
    fun locationsGetInactiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/locations/v2/inactive")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET GetTypes V1
     */
    fun locationsGetTypesV1(licensenumber: String? = null): Any? {
        val path = StringBuilder("/locations/v1/types")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of active location types for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetTypes V2
     */
    fun locationsGetTypesV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        val path = StringBuilder("/locations/v2/types")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (pagenumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pagenumber, StandardCharsets.UTF_8)}")
        }
        if (pagesize != null) {
            query.add("pageSize=${URLEncoder.encode(pagesize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates existing locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * PUT Update V2
     */
    fun locationsUpdateV2(licensenumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/locations/v2")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("PUT", path.toString(), body)
    }

    /**
     * Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Patients
     *
     * GET GetStatusesByPatientLicenseNumber V1
     */
    fun patientsStatusGetStatusesByPatientLicenseNumberV1(patientlicensenumber: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patients/v1/statuses/${URLEncoder.encode(patientlicensenumber, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Patients
     *
     * GET GetStatusesByPatientLicenseNumber V2
     */
    fun patientsStatusGetStatusesByPatientLicenseNumberV2(patientlicensenumber: String, licensenumber: String? = null): Any? {
        val path = StringBuilder("/patients/v2/statuses/${URLEncoder.encode(patientlicensenumber, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licensenumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licensenumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }
        return send<Any>("GET", path.toString(), null)
    }

}
