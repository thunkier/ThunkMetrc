package com.thunkmetrc.client;

import okhttp3.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.io.IOException;
import java.util.Map;
import java.util.List;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;

/**
 * Metrc Client for interacting with the Metrc API.
 */
public class MetrcClient {
    private final OkHttpClient client;
    private final ObjectMapper mapper;
    private final String baseUrl;
    private final String apiKey;

    /**
     * Constructor.
     * @param baseUrl Base URL for the API
     * @param vendorKey Vendor API Key
     * @param userKey User API Key
     */
    public MetrcClient(String baseUrl, String vendorKey, String userKey) {
        this.baseUrl = baseUrl.replaceAll("/$", "");
        this.apiKey = Credentials.basic(vendorKey, userKey);
        this.client = new OkHttpClient();
        this.mapper = new ObjectMapper();
    }

    private Object send(String method, String path, Object body) throws IOException {
        String url = this.baseUrl + path;
        Request.Builder builder = new Request.Builder()
            .url(url)
            .header("Authorization", this.apiKey);

        if (body != null && (method.equals("POST") || method.equals("PUT"))) {
            String json = this.mapper.writeValueAsString(body);
            builder.method(method, RequestBody.create(json, MediaType.parse("application/json")));
        } else if (method.equals("POST") || method.equals("PUT")) {
            builder.method(method, RequestBody.create("", null));
        } else {
            builder.method(method, null);
        }

        try (Response response = client.newCall(builder.build()).execute()) {
            if (!response.isSuccessful()) throw new IOException("Unexpected code " + response);
            if (response.code() == 204) return null;
            return this.mapper.readValue(response.body().string(), Object.class);
        }
    }

    /**
     * Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/additives");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records additive usage details applied to specific plants at a Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Plants
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocation V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesBylocationV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/additives/bylocation");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records additive usage for plants based on their location within a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocation V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesBylocationV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives/bylocation");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records additive usage for plants by location using a predefined additive template at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocationUsingtemplate V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesBylocationUsingtemplateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives/bylocation/usingtemplate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records additive usage for plants using predefined additive templates at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesUsingtemplate V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesUsingtemplateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives/usingtemplate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateChangegrowthphases V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateChangegrowthphasesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/changegrowthphases");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateHarvestplants V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateHarvestplantsV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/harvestplants");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates harvest product records from plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateManicure V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateManicureV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/manicure");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateManicureplants V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateManicureplantsV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/manicureplants");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateMoveplants V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateMoveplantsV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/moveplants");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreatePlantbatchPackageV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/create/plantbatch/packages");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreatePlantbatchPackageV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/plantbatch/packages");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreatePlantings V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreatePlantingsV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/create/plantings");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreatePlantingsV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/plantings");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateWasteV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/waste");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records waste events for plants at a Facility, including method, reason, and location details.
     * 
     *   Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsCreateWasteV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Destroy Veg/Flower Plants
     *
     * DELETE Delete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsDeleteV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Removes plants from a Facility’s inventory while recording the reason for their disposal.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Destroy Veg/Flower Plants
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsDeleteV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Plant by Id.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View/Manage Plants Additives
     *
     * GET GetAdditives V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetAdditivesV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/additives");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves additive records applied to plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View/Manage Plants Additives
     *
     * GET GetAdditives V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetAdditivesV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetAdditivesTypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetAdditivesTypesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/additives/types");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of all plant additive types defined within a Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAdditivesTypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetAdditivesTypesV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives/types");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetByLabel V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetByLabelV1(String label, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/"+URLEncoder.encode(label, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Plant by label.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetByLabel V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetByLabelV2(String label, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/"+URLEncoder.encode(label, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetFlowering V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetFloweringV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/flowering");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetFlowering V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetFloweringV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/flowering");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetGrowthPhases V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetGrowthPhasesV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/growthphases");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves the list of growth phases supported by a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetGrowthPhases V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetGrowthPhasesV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/growthphases");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves inactive plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves mother-phase plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMother V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetMotherV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/mother");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves inactive mother-phase plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMotherInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetMotherInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/mother/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMotherOnhold V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetMotherOnholdV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/mother/onhold");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetOnhold V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetOnholdV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/onhold");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves plants that are currently on hold at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetOnhold V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetOnholdV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/onhold");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetVegetative V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetVegetativeV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/vegetative");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetVegetative V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetVegetativeV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/vegetative");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of recorded plant waste events for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWaste V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetWasteV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteMethodsAll V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetWasteMethodsAllV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/waste/methods/all");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of all available plant waste methods for use within a Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteMethodsAll V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetWasteMethodsAllV2(String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste/methods/all");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWastePackage V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetWastePackageV2(String id, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/package");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWastePlant V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetWastePlantV2(String id, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/plant");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetWasteReasonsV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v1/waste/reasons");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retriveves available reasons for recording mature plant waste at a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsGetWasteReasonsV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste/reasons");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Adjusts the recorded count of plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateAdjust V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsUpdateAdjustV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/adjust");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Changes the growth phases of plants within a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateGrowthphase V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsUpdateGrowthphaseV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/growthphase");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * PUT UpdateHarvest V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsUpdateHarvestV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/harvest");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Moves plant batches to new locations within a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateLocation V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsUpdateLocationV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/location");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Merges multiple plant groups into a single group within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * PUT UpdateMerge V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsUpdateMergeV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/merge");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Splits an existing plant group into multiple groups within a Facility.
     * 
     *   Permissions Required:
     *   - View Plant
     *
     * PUT UpdateSplit V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsUpdateSplitV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/split");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates the strain information for plants within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateStrain V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsUpdateStrainV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/strain");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Replaces existing plant tags with new tags for plants within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateTag V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantsUpdateTagV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/tag");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Creates new additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * POST Create V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object additivesTemplatesCreateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/additivestemplates/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Retrieves an Additive Template by its Id.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object additivesTemplatesGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/additivestemplates/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object additivesTemplatesGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/additivestemplates/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of inactive additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object additivesTemplatesGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/additivestemplates/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Updates existing additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * PUT Update V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object additivesTemplatesUpdateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/additivestemplates/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Caregivers
     *
     * GET GetByCaregiverLicenseNumber V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object caregiversStatusGetByCaregiverLicenseNumberV1(String caregiverlicensenumber, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/caregivers/v1/status/"+URLEncoder.encode(caregiverlicensenumber, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Caregivers
     *
     * GET GetByCaregiverLicenseNumber V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object caregiversStatusGetByCaregiverLicenseNumberV2(String caregiverlicensenumber, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/caregivers/v2/status/"+URLEncoder.encode(caregiverlicensenumber, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * POST CreateFinish V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsCreateFinishV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/finish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackage V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsCreatePackageV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/create/packages");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsCreatePackageV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/packages");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageTesting V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsCreatePackageTestingV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/create/packages/testing");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsCreatePackageTestingV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/packages/testing");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * POST CreateRemoveWaste V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsCreateRemoveWasteV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/removewaste");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * POST CreateUnfinish V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsCreateUnfinishV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/unfinish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * POST CreateWaste V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsCreateWasteV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/waste");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Discontinues a specific harvest waste record by Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Discontinue Harvest Waste
     *
     * DELETE DeleteWaste V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsDeleteWasteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/waste/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active harvests for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of inactive harvests for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetOnhold V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetOnholdV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/onhold");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of harvests on hold for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetOnhold V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetOnholdV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/onhold");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetWaste V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetWasteV2(String harvestid, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/waste");
        StringBuilder query = new StringBuilder();
        if (harvestid != null) {
            if (query.length() > 0) query.append("&");
            query.append("harvestId=").append(URLEncoder.encode(harvestid, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteTypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetWasteTypesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/waste/types");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of Waste types for harvests.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteTypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsGetWasteTypesV2(String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/waste/types");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Marks one or more harvests as finished for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateFinish V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsUpdateFinishV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/finish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates the Location of Harvest for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateLocation V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsUpdateLocationV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/location");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateMove V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsUpdateMoveV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/move");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateRename V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsUpdateRenameV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v1/rename");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Renames one or more harvests for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateRename V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsUpdateRenameV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/rename");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Restores previously harvested plants to their original state for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateRestoreHarvestedPlants V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsUpdateRestoreHarvestedPlantsV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/restore/harvestedplants");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Reopens one or more previously finished harvests for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateUnfinish V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object harvestsUpdateUnfinishV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/unfinish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRecord V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsCreateRecordV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v1/record");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRecord V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsCreateRecordV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/record");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Retrieves a list of Lab Test batches.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetBatches V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetBatchesV2(String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/batches");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetLabtestdocument V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetLabtestdocumentV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v1/labtestdocument/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a specific Lab Test result document by its Id for a given Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetLabtestdocument V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetLabtestdocumentV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/labtestdocument/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetResults V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetResultsV1(String licensenumber, String packageid) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v1/results");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (packageid != null) {
            if (query.length() > 0) query.append("&");
            query.append("packageId=").append(URLEncoder.encode(packageid, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves Lab Test results for a specified Package.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetResults V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetResultsV2(String licensenumber, String packageid, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/results");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (packageid != null) {
            if (query.length() > 0) query.append("&");
            query.append("packageId=").append(URLEncoder.encode(packageid, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetStates V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetStatesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v1/states");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of all lab testing states.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetStates V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetStatesV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/states");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetTypesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v1/types");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of Lab Test types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsGetTypesV2(String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/types");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateLabtestdocument V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsUpdateLabtestdocumentV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v1/labtestdocument");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates one or more documents for previously submitted lab tests.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateLabtestdocument V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsUpdateLabtestdocumentV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/labtestdocument");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateResultRelease V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsUpdateResultReleaseV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v1/results/release");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Releases Lab Test results for one or more packages.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateResultRelease V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object labTestsUpdateResultReleaseV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/results/release");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * POST CreateDelivery V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailer V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerDepart V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerDepartV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer/depart");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerDepartV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/depart");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerEnd V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerEndV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer/end");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerEndV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/end");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerRestock V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerRestockV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer/restock");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerRestockV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/restock");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerSale V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerSaleV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer/sale");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerSaleV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/sale");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales
     *
     * POST CreateReceipt V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateReceiptV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/receipts");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateReceiptV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * POST CreateTransactionByDate V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesCreateTransactionByDateV1(String date, String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/transactions/"+URLEncoder.encode(date, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * DELETE DeleteDelivery V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesDeleteDeliveryV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * DELETE DeleteDelivery V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesDeleteDeliveryV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * DELETE DeleteDeliveryRetailer V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesDeleteDeliveryRetailerV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesDeleteDeliveryRetailerV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * DELETE DeleteReceipt V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesDeleteReceiptV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/receipts/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * DELETE DeleteReceipt V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesDeleteReceiptV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCounties V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetCountiesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/counties");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of counties available for sales deliveries.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCounties V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetCountiesV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/counties");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCustomertypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetCustomertypesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/customertypes");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of customer types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCustomertypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetCustomertypesV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/customertypes");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String salesdateend, String salesdatestart) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (salesdateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateEnd=").append(URLEncoder.encode(salesdateend, StandardCharsets.UTF_8));
        }
        if (salesdatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateStart=").append(URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDeliveriesActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, String salesdateend, String salesdatestart) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (salesdateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateEnd=").append(URLEncoder.encode(salesdateend, StandardCharsets.UTF_8));
        }
        if (salesdatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateStart=").append(URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String salesdateend, String salesdatestart) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (salesdateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateEnd=").append(URLEncoder.encode(salesdateend, StandardCharsets.UTF_8));
        }
        if (salesdatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateStart=").append(URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDeliveriesInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, String salesdateend, String salesdatestart) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (salesdateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateEnd=").append(URLEncoder.encode(salesdateend, StandardCharsets.UTF_8));
        }
        if (salesdatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateStart=").append(URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveriesRetailerActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesRetailerActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveriesRetailerActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesRetailerActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveriesRetailerInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesRetailerInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveriesRetailerInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesRetailerInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetDeliveriesReturnreasons V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesReturnreasonsV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/returnreasons");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of return reasons for sales deliveries based on the provided License Number.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesReturnreasons V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesReturnreasonsV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/returnreasons");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDelivery V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveryV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a sales delivery record by its Id, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDelivery V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveryV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveryRetailer V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveryRetailerV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a retailer delivery record by its ID, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveryRetailer V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetDeliveryRetailerV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetPatientRegistrationsLocations V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetPatientRegistrationsLocationsV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/patientregistration/locations");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of valid Patient registration locations for sales.
     * 
     *   Permissions Required:
     *   -
     *
     * GET GetPatientRegistrationsLocations V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetPatientRegistrationsLocationsV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/patientregistration/locations");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetPaymenttypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetPaymenttypesV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/paymenttypes");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of available payment types for the specified License Number.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetPaymenttypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetPaymenttypesV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/paymenttypes");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceipt V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetReceiptV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/receipts/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a sales receipt by its Id, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceipt V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetReceiptV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceiptsActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetReceiptsActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String salesdateend, String salesdatestart) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/receipts/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (salesdateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateEnd=").append(URLEncoder.encode(salesdateend, StandardCharsets.UTF_8));
        }
        if (salesdatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateStart=").append(URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetReceiptsActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, String salesdateend, String salesdatestart) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (salesdateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateEnd=").append(URLEncoder.encode(salesdateend, StandardCharsets.UTF_8));
        }
        if (salesdatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateStart=").append(URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Sales Receipt by its external number, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsExternalByExternalNumber V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetReceiptsExternalByExternalNumberV2(String externalnumber, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/external/"+URLEncoder.encode(externalnumber, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceiptsInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetReceiptsInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String salesdateend, String salesdatestart) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/receipts/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (salesdateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateEnd=").append(URLEncoder.encode(salesdateend, StandardCharsets.UTF_8));
        }
        if (salesdatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateStart=").append(URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetReceiptsInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, String salesdateend, String salesdatestart) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (salesdateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateEnd=").append(URLEncoder.encode(salesdateend, StandardCharsets.UTF_8));
        }
        if (salesdatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("salesDateStart=").append(URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetTransactions V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetTransactionsV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/transactions");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesGetTransactionsBySalesDateStartAndSalesDateEndV1(String salesdatestart, String salesdateend, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/transactions/"+URLEncoder.encode(salesdatestart, StandardCharsets.UTF_8)+"/"+URLEncoder.encode(salesdateend, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDelivery V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * PUT UpdateDelivery V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDeliveryComplete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryCompleteV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/complete");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * PUT UpdateDeliveryComplete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryCompleteV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/complete");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDeliveryHub V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/hub");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery, Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHub V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/hub");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubAccept V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubAcceptV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/hub/accept");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubAccept V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubAcceptV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/hub/accept");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubDepart V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubDepartV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/hub/depart");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubDepart V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubDepartV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/hub/depart");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubVerifyID V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubVerifyIdV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/hub/verifyID");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubVerifyID V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubVerifyIdV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/hub/verifyID");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * PUT UpdateDeliveryRetailer V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryRetailerV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/deliveries/retailer");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryRetailerV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales
     *
     * PUT UpdateReceipt V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateReceiptV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/receipts");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceipt V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateReceiptV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceiptFinalize V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateReceiptFinalizeV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/finalize");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceiptUnfinalize V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateReceiptUnfinalizeV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/unfinalize");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateTransactionByDate V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object salesUpdateTransactionByDateV1(String date, String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v1/transactions/"+URLEncoder.encode(date, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Creates new driver records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * POST CreateDriver V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersCreateDriverV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates new vehicle records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * POST CreateVehicle V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersCreateVehicleV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * DELETE DeleteDriver V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersDeleteDriverV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * DELETE DeleteVehicle V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersDeleteVehicleV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetDriver V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersGetDriverV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of drivers for a Facility.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetDrivers V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersGetDriversV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetVehicle V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersGetVehicleV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of vehicles for a Facility.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetVehicles V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersGetVehiclesV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Updates existing driver records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * PUT UpdateDriver V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersUpdateDriverV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates existing vehicle records for a facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * PUT UpdateVehicle V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transportersUpdateVehicleV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object facilitiesGetAllV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/facilities/v1");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object facilitiesGetAllV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/facilities/v2");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Create V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/create");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates new packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Create V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateAdjust V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateAdjustV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/adjust");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records a list of adjustments for packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateAdjust V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateAdjustV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/adjust");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateChangeItem V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateChangeItemV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/change/item");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateChangeLocation V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateChangeLocationV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/change/locations");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateFinish V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateFinishV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/finish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreatePlantings V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreatePlantingsV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/create/plantings");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreatePlantingsV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/plantings");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRemediate V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateRemediateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/remediate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateTesting V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateTestingV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/create/testing");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates new packages for testing for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateTesting V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateTestingV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/testing");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateUnfinish V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesCreateUnfinishV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/unfinish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Discontinues a Package at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesDeleteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Package by its Id.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetAdjustReasons V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetAdjustReasonsV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/adjust/reasons");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of adjustment reasons for packages at a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAdjustReasons V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetAdjustReasonsV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/adjust/reasons");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetByLabel V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetByLabelV1(String label, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/"+URLEncoder.encode(label, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Package by its label.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetByLabel V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetByLabelV2(String label, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/"+URLEncoder.encode(label, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of inactive packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of packages in transit for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetIntransit V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetIntransitV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/intransit");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetLabsamples V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetLabsamplesV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/labsamples");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetOnhold V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetOnholdV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/onhold");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of packages on hold for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetOnhold V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetOnholdV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/onhold");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves the source harvests for a Package by its Id.
     * 
     *   Permissions Required:
     *   - View Package Source Harvests
     *
     * GET GetSourceHarvest V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetSourceHarvestV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/source/harvests");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transferred packages for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetTransferred V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetTransferredV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/transferred");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetTypesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/types");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of available Package types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesGetTypesV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/types");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Set the final quantity for a Package.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateAdjust V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateAdjustV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/adjust");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *   - Manage Package Notes
     *
     * PUT UpdateChangeNote V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateChangeNoteV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v1/change/note");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates the Product decontaminate information for a list of packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDecontaminate V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateDecontaminateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/decontaminate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Flags one or more packages for donation at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDonationFlag V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateDonationFlagV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/donation/flag");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Removes the donation flag from one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDonationUnflag V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateDonationUnflagV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/donation/unflag");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateExternalidV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/externalid");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates a list of packages as finished for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateFinish V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateFinishV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/finish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates the associated Item for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateItem V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateItemV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/item");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates the list of required lab test batches for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateLabTestRequired V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateLabTestRequiredV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/labtests/required");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates the Location and Sublocation for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateLocation V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateLocationV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/location");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateNoteV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/note");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates a list of Product remediations for packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateRemediate V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateRemediateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/remediate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Flags or unflags one or more packages at the specified Facility as trade samples.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateTradesampleFlag V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateTradesampleFlagV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/tradesample/flag");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Removes the trade sample flag from one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateTradesampleUnflag V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateTradesampleUnflagV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/tradesample/unflag");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates a list of packages as unfinished for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateUnfinish V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateUnfinishV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/unfinish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates the use-by date for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateUsebydate V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object packagesUpdateUsebydateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/packages/v2/usebydate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * POST Create V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsCreateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v1");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * POST Create V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsCreateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * DELETE Delete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsDeleteV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives a Patient Check-In, identified by its Id, for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsDeleteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * GET GetAll V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsGetAllV1(String checkindateend, String checkindatestart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v1");
        StringBuilder query = new StringBuilder();
        if (checkindateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("checkinDateEnd=").append(URLEncoder.encode(checkindateend, StandardCharsets.UTF_8));
        }
        if (checkindatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("checkinDateStart=").append(URLEncoder.encode(checkindatestart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * GET GetAll V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsGetAllV2(String checkindateend, String checkindatestart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2");
        StringBuilder query = new StringBuilder();
        if (checkindateend != null) {
            if (query.length() > 0) query.append("&");
            query.append("checkinDateEnd=").append(URLEncoder.encode(checkindateend, StandardCharsets.UTF_8));
        }
        if (checkindatestart != null) {
            if (query.length() > 0) query.append("&");
            query.append("checkinDateStart=").append(URLEncoder.encode(checkindatestart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetLocations V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsGetLocationsV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v1/locations");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of Patient Check-In locations.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetLocations V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsGetLocationsV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2/locations");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * PUT Update V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsUpdateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v1");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * PUT Update V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientCheckInsUpdateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object unitsOfMeasureGetActiveV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/unitsofmeasure/v1/active");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves all active units of measure.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object unitsOfMeasureGetActiveV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/unitsofmeasure/v2/active");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves all inactive units of measure.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object unitsOfMeasureGetInactiveV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/unitsofmeasure/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves all available waste methods.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object wasteMethodsGetAllV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/wastemethods/v2");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Employees
     *
     * GET GetAll V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object employeesGetAllV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/employees/v1");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of employees for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Employees
     *   - View Employees
     *
     * GET GetAll V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object employeesGetAllV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/employees/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
     * 
     *   Permissions Required:
     *   - Manage Employees
     *
     * GET GetPermissions V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object employeesGetPermissionsV2(String employeelicensenumber, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/employees/v2/permissions");
        StringBuilder query = new StringBuilder();
        if (employeelicensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("employeeLicenseNumber=").append(URLEncoder.encode(employeelicensenumber, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST Create V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsCreateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/create");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST Create V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsCreateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates one or more new item brands for the specified Facility identified by the License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreateBrand V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsCreateBrandV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/brand");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreateFile V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsCreateFileV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/file");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreatePhoto V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsCreatePhotoV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/photo");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreatePhoto V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsCreatePhotoV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/photo");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * POST CreateUpdate V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsCreateUpdateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/update");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * DELETE Delete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsDeleteV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives the specified Product by Id for the given Facility License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsDeleteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives the specified Item Brand by Id for the given Facility License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * DELETE DeleteBrand V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsDeleteBrandV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/brand/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves detailed information about a specific Item by Id.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetActiveV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/active");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of active items for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetBrands V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetBrandsV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/brands");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active item brands for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetBrands V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetBrandsV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/brands");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCategories V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetCategoriesV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/categories");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of item categories.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCategories V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetCategoriesV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/categories");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a file by its Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetFile V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetFileV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/file/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetInactiveV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/inactive");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of inactive items for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetInactiveV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetPhoto V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetPhotoV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v1/photo/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves an image by its Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetPhoto V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsGetPhotoV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/photo/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Updates one or more existing products for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * PUT Update V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsUpdateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates one or more existing item brands for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * PUT UpdateBrand V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object itemsUpdateBrandV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/items/v2/brand");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * POST Create V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsCreateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v1/create");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates new locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * POST Create V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsCreateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * POST CreateUpdate V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsCreateUpdateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v1/update");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsDeleteV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives a specified Location, identified by its Id, for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsDeleteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Location by its Id.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsGetActiveV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v1/active");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of inactive locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsGetInactiveV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET GetTypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsGetTypesV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v1/types");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active location types for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetTypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsGetTypesV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v2/types");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Updates existing locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * PUT Update V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object locationsUpdateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/locations/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Adds new patients to a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * POST Create V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsCreateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * POST CreateAdd V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsCreateAddV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v1/add");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * POST CreateUpdate V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsCreateUpdateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v1/update");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * DELETE Delete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsDeleteV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Removes a Patient, identified by an Id, from a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsDeleteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Patient by Id.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsGetActiveV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v1/active");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active patients for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsGetActiveV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v2/active");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Updates Patient information for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * PUT Update V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsUpdateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object retailIdCreateAssociateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/associate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object retailIdCreateGenerateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/generate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object retailIdCreateMergeV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/merge");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object retailIdCreatePackageInfoV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/packages/info");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object retailIdGetReceiveByLabelV2(String label, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/receive/"+URLEncoder.encode(label, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object retailIdGetReceiveQrByShortCodeV2(String shortcode, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/retailid/v2/receive/qr/"+URLEncoder.encode(shortcode, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
     * 
     *   Permissions Required:
     *   - None
     *
     * POST CreateIntegratorSetup V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object sandboxCreateIntegratorSetupV2(String userkey, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sandbox/v2/integrator/setup");
        StringBuilder query = new StringBuilder();
        if (userkey != null) {
            if (query.length() > 0) query.append("&");
            query.append("userKey=").append(URLEncoder.encode(userkey, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * POST Create V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsCreateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v1/create");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates new strain records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * POST Create V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsCreateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * POST CreateUpdate V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsCreateUpdateV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v1/update");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * DELETE Delete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsDeleteV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives an existing strain record for a Facility
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsDeleteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Strain record by its Id, with an optional license number.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsGetActiveV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v1/active");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsGetInactiveV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Updates existing strain records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * PUT Update V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object strainsUpdateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/strains/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Returns a list of available package tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *
     * GET GetPackageAvailable V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object tagsGetPackageAvailableV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/tags/v2/package/available");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of available plant tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *
     * GET GetPlantAvailable V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object tagsGetPlantAvailableV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/tags/v2/plant/available");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of staged tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *   - RetailId.AllowPackageStaging Key Value enabled
     *
     * GET GetStaged V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object tagsGetStagedV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/tags/v2/staged");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Patients
     *
     * GET GetStatusesByPatientLicenseNumber V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsStatusGetStatusesByPatientLicenseNumberV1(String patientlicensenumber, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v1/statuses/"+URLEncoder.encode(patientlicensenumber, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Patients
     *
     * GET GetStatusesByPatientLicenseNumber V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object patientsStatusGetStatusesByPatientLicenseNumberV2(String patientlicensenumber, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v2/statuses/"+URLEncoder.encode(patientlicensenumber, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdditivesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/additives");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records Additive usage details for plant batches at a specific Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdditivesV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/additives");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records Additive usage for plant batches at a Facility using predefined additive templates.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesUsingtemplate V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdditivesUsingtemplateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/additives/usingtemplate");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateAdjust V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdjustV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/adjust");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateAdjust V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdjustV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/adjust");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateChangegrowthphase V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateChangegrowthphaseV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/changegrowthphase");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateGrowthphaseV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/growthphase");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreatePackageV2(String isfrommotherplant, String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/packages");
        StringBuilder query = new StringBuilder();
        if (isfrommotherplant != null) {
            if (query.length() > 0) query.append("&");
            query.append("isFromMotherPlant=").append(URLEncoder.encode(isfrommotherplant, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageFrommotherplant V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreatePackageFrommotherplantV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/create/packages/frommotherplant");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
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
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreatePackageFrommotherplantV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/packages/frommotherplant");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates new plantings for a Facility by generating plant batches based on provided planting details.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreatePlantings V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreatePlantingsV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/plantings");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateSplit V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateSplitV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/split");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Splits an existing Plant Batch into multiple groups at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateSplit V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateSplitV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/split");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateWasteV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/waste");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Records waste information for plant batches based on the submitted data for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateWasteV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/waste");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Createpackages V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreatepackagesV1(String isfrommotherplant, String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/createpackages");
        StringBuilder query = new StringBuilder();
        if (isfrommotherplant != null) {
            if (query.length() > 0) query.append("&");
            query.append("isFromMotherPlant=").append(URLEncoder.encode(isfrommotherplant, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST Createplantings V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesCreateplantingsV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/createplantings");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Destroy Immature Plants
     *
     * DELETE Delete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesDeleteV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Completes the destruction of plant batches based on the provided input data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Destroy Immature Plants
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesDeleteV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Plant Batch by Id.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET GetInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetTypesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/types");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of plant batch types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetTypesV2(String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/types");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves waste details associated with plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWaste V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetWasteV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/waste");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetWasteReasonsV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/waste/reasons");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesGetWasteReasonsV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/waste/reasons");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Moves one or more plant batches to new locations with in a specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *
     * PUT UpdateLocation V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateLocationV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/location");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * PUT UpdateMoveplantbatches V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateMoveplantbatchesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v1/moveplantbatches");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Renames plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateName V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateNameV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/name");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Changes the strain of plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateStrain V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateStrainV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/strain");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Replaces tags for plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateTag V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateTagV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/plantbatches/v2/tag");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST CreateAdjust V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsCreateAdjustV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/adjust");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateAdjust V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsCreateAdjustV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/adjust");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateJobtypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsCreateJobtypesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/jobtypes");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateJobtypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsCreateJobtypesV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST CreateStart V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsCreateStartV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/start");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Initiates new processing jobs at a Facility, including job details and associated packages.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateStart V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsCreateStartV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/start");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST Createpackages V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsCreatepackagesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/createpackages");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates packages from processing jobs at a Facility, including optional location and note assignments.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST Createpackages V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsCreatepackagesV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/createpackages");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE Delete V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsDeleteV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsDeleteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE DeleteJobtypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsDeleteJobtypesV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/jobtypes/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives a Processing Job Type at a Facility, making it inactive for future use.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE DeleteJobtypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsDeleteJobtypesV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET Get V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a ProcessingJob by Id.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves active processing jobs at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves inactive processing jobs at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesActive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/jobtypes/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of all active processing job types defined within a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesAttributes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesAttributesV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/jobtypes/attributes");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves all processing job attributes available for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesAttributes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesAttributesV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/attributes");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesCategories V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesCategoriesV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/jobtypes/categories");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves all processing job categories available for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesCategories V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesCategoriesV2(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/categories");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesInactive V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/jobtypes/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of all inactive processing job types defined within a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/inactive");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateFinish V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsUpdateFinishV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/finish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Completes processing jobs at a Facility by recording final notes and waste measurements.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateFinish V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsUpdateFinishV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/finish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateJobtypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsUpdateJobtypesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/jobtypes");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateJobtypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsUpdateJobtypesV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateUnfinish V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsUpdateUnfinishV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v1/unfinish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateUnfinish V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object processingJobsUpdateUnfinishV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/unfinish");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Creates new sublocation records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * POST Create V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object sublocationsCreateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Archives an existing Sublocation record for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object sublocationsDeleteV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Retrieves a Sublocation by its Id, with an optional license number.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET Get V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object sublocationsGetV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object sublocationsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2/active");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of inactive sublocations for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetInactive V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object sublocationsGetInactiveV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Updates existing sublocation records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * PUT Update V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object sublocationsUpdateV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * POST CreateExternalIncoming V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersCreateExternalIncomingV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/external/incoming");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates external incoming shipment plans for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * POST CreateExternalIncoming V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersCreateExternalIncomingV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/external/incoming");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * POST CreateTemplates V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersCreateTemplatesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/templates");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Creates new transfer templates for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * POST CreateTemplatesOutgoing V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersCreateTemplatesOutgoingV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("POST", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * DELETE DeleteExternalIncoming V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersDeleteExternalIncomingV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/external/incoming/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Voids an external incoming shipment plan for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * DELETE DeleteExternalIncoming V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersDeleteExternalIncomingV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/external/incoming/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * DELETE DeleteTemplates V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersDeleteTemplatesV1(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/templates/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Archives a transfer template for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * DELETE DeleteTemplatesOutgoing V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersDeleteTemplatesOutgoingV2(String id, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("DELETE", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetDeliveriesPackagesStates V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveriesPackagesStatesV1(String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/deliveries/packages/states");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Returns a list of available shipment Package states.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetDeliveriesPackagesStates V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveriesPackagesStatesV2(String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/packages/states");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: that the {id} parameter above represents a Shipment Plan ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDelivery V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/deliveries");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDelivery V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/deliveries");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackage V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackage V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackageRequiredlabtestbatches V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageRequiredlabtestbatchesV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/deliveries/package/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/requiredlabtestbatches");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackageRequiredlabtestbatches V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageRequiredlabtestbatchesV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/package/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/requiredlabtestbatches");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackageWholesale V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageWholesaleV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages/wholesale");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackageWholesale V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageWholesaleV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages/wholesale");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: that the {id} parameter above represents a Shipment Delivery ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryTransporters V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryTransportersV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryTransporters V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryTransportersV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: The {id} parameter above represents a Shipment Delivery ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryTransportersDetails V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryTransportersDetailsV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters/details");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryTransportersDetails V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryTransportersDetailsV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters/details");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetHub V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetHubV2(String estimatedarrivalend, String estimatedarrivalstart, String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/hub");
        StringBuilder query = new StringBuilder();
        if (estimatedarrivalend != null) {
            if (query.length() > 0) query.append("&");
            query.append("estimatedArrivalEnd=").append(URLEncoder.encode(estimatedarrivalend, StandardCharsets.UTF_8));
        }
        if (estimatedarrivalstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("estimatedArrivalStart=").append(URLEncoder.encode(estimatedarrivalstart, StandardCharsets.UTF_8));
        }
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetIncoming V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetIncomingV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/incoming");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetIncoming V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetIncomingV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/incoming");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetOutgoing V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetOutgoingV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/outgoing");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetOutgoing V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetOutgoingV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/outgoing");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetRejected V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetRejectedV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/rejected");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of shipments with rejected packages for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetRejected V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetRejectedV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/rejected");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplates V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/templates");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDelivery V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesDeliveryV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/templates/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/deliveries");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetTemplatesDeliveryPackage V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesDeliveryPackageV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/templates/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDeliveryTransporters V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesDeliveryTransportersV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/templates/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDeliveryTransportersDetails V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesDeliveryTransportersDetailsV1(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/templates/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters/details");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoing V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing");
        StringBuilder query = new StringBuilder();
        if (lastmodifiedend != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastmodifiedend, StandardCharsets.UTF_8));
        }
        if (lastmodifiedstart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastmodifiedstart, StandardCharsets.UTF_8));
        }
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of deliveries associated with a specific transfer template.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDelivery V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingDeliveryV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/deliveries");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryPackage V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingDeliveryPackageV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryTransporters V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingDeliveryTransportersV2(String id, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters");
        StringBuilder query = new StringBuilder();
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryTransportersDetails V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(String id, String no) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters/details");
        StringBuilder query = new StringBuilder();
        if (no != null) {
            if (query.length() > 0) query.append("&");
            query.append("No=").append(URLEncoder.encode(no, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTypesV1(String licensenumber) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/types");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of available transfer types for a Facility based on its license number.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersGetTypesV2(String licensenumber, String pagenumber, String pagesize) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/types");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (pagenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pagenumber, StandardCharsets.UTF_8));
        }
        if (pagesize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pagesize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("GET", path.toString(), null);
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * PUT UpdateExternalIncoming V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersUpdateExternalIncomingV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/external/incoming");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates external incoming shipment plans for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * PUT UpdateExternalIncoming V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersUpdateExternalIncomingV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/external/incoming");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * PUT UpdateTemplates V1
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersUpdateTemplatesV1(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v1/templates");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

    /**
     * Updates existing transfer templates for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * PUT UpdateTemplatesOutgoing V2
     * @throws IOException If request fails
     * @return Response object
     */
    public Object transfersUpdateTemplatesOutgoingV2(String licensenumber, Object body) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing");
        StringBuilder query = new StringBuilder();
        if (licensenumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licensenumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);
        return send("PUT", path.toString(), body);
    }

}
