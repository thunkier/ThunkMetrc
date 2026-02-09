package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class PlantsClient {
    private final MetrcClient client;

    public PlantsClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Records additive usage for plants based on their location within a specified Facility.
     * Permissions Required:
     * - Manage Plants
     * - Manage Plants Additives
     * POST CreateAdditivesByLocation
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPlantsAdditivesByLocation(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives/bylocation");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Records additive usage for plants by location using a predefined additive template at a specified Facility.
     * Permissions Required:
     * - Manage Plants Additives
     * POST CreateAdditivesByLocationUsingTemplate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPlantsAdditivesByLocationUsingTemplate(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives/bylocation/usingtemplate");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Creates harvest product records from plant batches at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manicure/Harvest Veg/Flower Plants
     * POST CreateManicure
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPlantsManicure(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/manicure");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Creates packages from plant batches at a specified Facility.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * - View Packages
     * - Create/Submit/Discontinue Packages
     * POST CreatePlantBatchPackages
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPlantsPlantBatchPackages(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/plantbatch/packages");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Records additive usage details applied to specific plants at a Facility.
     * Permissions Required:
     * - Manage Plants Additives
     * POST CreatePlantsAdditives
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPlantsAdditives(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Records additive usage for plants using predefined additive templates at a specified Facility.
     * Permissions Required:
     * - Manage Plants Additives
     * POST CreatePlantsAdditivesUsingTemplate
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPlantsAdditivesUsingTemplate(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives/usingtemplate");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Creates new plant batches at a specified Facility from existing plant data.
     * Permissions Required:
     * - View Immature Plants
     * - Manage Immature Plants Inventory
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * POST CreatePlantsPlantings
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPlantsPlantings(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/plantings");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Records waste events for plants at a Facility, including method, reason, and location details.
     * Permissions Required:
     * - Manage Plants Waste
     * POST CreatePlantsWaste
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPlantsWaste(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Removes plants from a Facility’s inventory while recording the reason for their disposal.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Destroy Veg/Flower Plants
     * DELETE DeletePlants
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deletePlants(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Retrieves additive records applied to plants at a specified Facility.
     * Permissions Required:
     * - View/Manage Plants Additives
     * GET GetAdditives
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsAdditives(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of all plant additive types defined within a Facility.
     * GET GetAdditivesTypes
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsAdditivesTypes(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/additives/types");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET GetFloweringPlants
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getFloweringPlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/flowering");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves the list of growth phases supported by a specified Facility.
     * GET GetGrowthPhases
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsGrowthPhases(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/growthphases");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves inactive plants at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET GetInactivePlants
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getInactivePlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves inactive mother-phase plants at a specified Facility.
     * Permissions Required:
     * - View Mother Plants
     * GET GetMotherInactivePlants
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getMotherInactivePlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/mother/inactive");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
     * Permissions Required:
     * - View Mother Plants
     * GET GetMotherOnHoldPlants
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getMotherOnHoldPlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/mother/onhold");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves mother-phase plants at a specified Facility.
     * Permissions Required:
     * - View Mother Plants
     * GET GetMotherPlants
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getMotherPlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/mother");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves plants that are currently on hold at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET GetOnHoldPlants
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getOnHoldPlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/onhold");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Plant by Id.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET GetPlantsById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Plant by label.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET GetPlantsByLabel
     * @param label Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsByLabel(
        String label, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/"+URLEncoder.encode(label, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of recorded plant waste events for a specific Facility.
     * Permissions Required:
     * - View Plants Waste
     * GET GetPlantsWaste
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsWaste(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of all available plant waste methods for use within a Facility.
     * GET GetPlantsWasteMethods
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsWasteMethods(
        String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste/methods/all");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retriveves available reasons for recording mature plant waste at a specified Facility.
     * GET GetPlantsWasteReasons
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsWasteReasons(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste/reasons");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
     * Permissions Required:
     * - View Veg/Flower Plants
     * GET GetVegetativePlants
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getVegetativePlants(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/vegetative");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
     * Permissions Required:
     * - View Plants Waste
     * GET GetWasteById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsWasteById(
        String id, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/plant");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
     * Permissions Required:
     * - View Plants Waste
     * GET GetWastePackageById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPlantsWastePackageById(
        String id, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/waste/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/package");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Adjusts the recorded count of plants at a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT UpdateAdjustPlants
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateAdjustPlants(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/adjust");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Changes the growth phases of plants within a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT UpdateGrowthPhase
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updatePlantsGrowthPhase(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/growthphase");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manicure/Harvest Veg/Flower Plants
     * PUT UpdateHarvest
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updatePlantsHarvest(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/harvest");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Merges multiple plant groups into a single group within a Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manicure/Harvest Veg/Flower Plants
     * PUT UpdateMerge
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updatePlantsMerge(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/merge");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Moves plant batches to new locations within a specified Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT UpdatePlantsLocation
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updatePlantsLocation(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/location");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Updates the strain information for plants within a Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT UpdatePlantsStrain
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updatePlantsStrain(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/strain");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Replaces existing plant tags with new tags for plants within a Facility.
     * Permissions Required:
     * - View Veg/Flower Plants
     * - Manage Veg/Flower Plants Inventory
     * PUT UpdatePlantsTag
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updatePlantsTag(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/tag");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Splits an existing plant group into multiple groups within a Facility.
     * Permissions Required:
     * - View Plant
     * PUT UpdateSplit
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updatePlantsSplit(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/plants/v2/split");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

