package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class LabTestsClient {
    private final MetrcClient client;

    public LabTestsClient(MetrcClient client) {
        this.client = client;
    }
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
    public Object createLabTestsRecord(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/record");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Retrieves a list of Lab Test batches.
     * GET GetBatches
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getLabTestsBatches(
        String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/batches");
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
    public Object getLabTestsLabTestDocumentById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/labtestdocument/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Returns a list of Lab Test types.
     * GET GetLabTestsTypes
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getLabTestsTypes(
        String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/types");
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
    public Object getLabTestsResults(
        String licenseNumber, String packageId, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/results");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (packageId != null) {
            if (query.length() > 0) query.append("&");
            query.append("packageId=").append(URLEncoder.encode(packageId, StandardCharsets.UTF_8));
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
     * Returns a list of all lab testing states.
     * GET GetStates
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getLabTestsStates(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/states");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
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
    public Object updateLabTestsLabTestDocument(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/labtestdocument");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
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
    public Object updateLabTestsResultsRelease(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/labtests/v2/results/release");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

