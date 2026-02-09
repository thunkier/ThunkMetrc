package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class TagsClient {
    private final MetrcClient client;

    public TagsClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Returns a list of available package tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * GET GetAvailablePackage
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTagsAvailablePackage(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/tags/v2/package/available");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Returns a list of available plant tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * GET GetAvailablePlant
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTagsAvailablePlant(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/tags/v2/plant/available");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Returns a list of staged tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * - RetailId.AllowPackageStaging Key Value enabled
     * GET GetStagedTags
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getStagedTags(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/tags/v2/staged");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

}

