package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class HarvestsClient {
    private final MetrcClient client;

    public HarvestsClient(MetrcClient client) {
        this.client = client;
    }
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
    public Object createHarvestsPackages(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/packages");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
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
    public Object createHarvestsWaste(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/waste");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
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
    public Object createHarvestsPackagesTesting(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/packages/testing");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
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
    public Object deleteHarvestsWasteById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/waste/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
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
    public Object finishHarvestsHarvests(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/finish");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
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
    public Object getActiveHarvests(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/active");
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
     * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
     * Permissions Required:
     * - View Harvests
     * GET GetHarvestsById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getHarvestsById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
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
    public Object getHarvestsWaste(
        String harvestId, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/waste");
        StringBuilder query = new StringBuilder();
        if (harvestId != null) {
            if (query.length() > 0) query.append("&");
            query.append("harvestId=").append(URLEncoder.encode(harvestId, StandardCharsets.UTF_8));
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
    public Object getInactiveHarvests(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/inactive");
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
    public Object getOnHoldHarvests(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/onhold");
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
     * Retrieves a list of Waste types for harvests.
     * GET GetWasteTypes
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getHarvestsWasteTypes(
        String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/waste/types");
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
    public Object unfinishHarvestsHarvests(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/unfinish");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
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
    public Object updateHarvestsLocation(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/location");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
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
    public Object updateHarvestsRename(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/rename");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
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
    public Object updateHarvestsRestoreHarvestedPlants(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/harvests/v2/restore/harvestedplants");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

