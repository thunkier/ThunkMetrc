package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class SublocationsClient {
    private final MetrcClient client;

    public SublocationsClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Creates new sublocation records for a Facility.
     * Permissions Required:
     * - Manage Locations
     * POST CreateSublocations
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createSublocations(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Archives an existing Sublocation record for a Facility.
     * Permissions Required:
     * - Manage Locations
     * DELETE DeleteSublocationsById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteSublocationsById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Locations
     * GET GetActiveSublocations
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getActiveSublocations(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2/active");
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
     * Retrieves a list of inactive sublocations for the specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET GetInactiveSublocations
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getInactiveSublocations(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2/inactive");
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
     * Retrieves a Sublocation by its Id, with an optional license number.
     * Permissions Required:
     * - Manage Locations
     * GET GetSublocationsById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSublocationsById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Updates existing sublocation records for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * PUT UpdateSublocations
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSublocations(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sublocations/v2");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

