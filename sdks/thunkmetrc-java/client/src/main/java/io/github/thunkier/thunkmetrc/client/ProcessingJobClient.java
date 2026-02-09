package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class ProcessingJobClient {
    private final MetrcClient client;

    public ProcessingJobClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
     * Permissions Required:
     * - Manage Processing Job
     * POST CreateAdjustProcessingJob
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createAdjustProcessingJob(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/adjust");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
     * Permissions Required:
     * - Manage Processing Job
     * POST CreateJobTypes
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createProcessingJobJobTypes(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Creates packages from processing jobs at a Facility, including optional location and note assignments.
     * Permissions Required:
     * - Manage Processing Job
     * POST CreateProcessingJobPackages
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createProcessingJobPackages(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/createpackages");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Archives a Processing Job Type at a Facility, making it inactive for future use.
     * Permissions Required:
     * - Manage Processing Job
     * DELETE DeleteJobTypeById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteProcessingJobJobTypeById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
     * Permissions Required:
     * - Manage Processing Job
     * DELETE DeleteProcessingJobById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteProcessingJobById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Completes processing jobs at a Facility by recording final notes and waste measurements.
     * Permissions Required:
     * - Manage Processing Job
     * PUT FinishProcessingJob
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object finishProcessingJobProcessingJob(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/finish");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Retrieves a list of all active processing job types defined within a Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET GetActiveJobTypes
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getProcessingJobActiveJobTypes(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/active");
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
     * Retrieves active processing jobs at a specified Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET GetActiveProcessingJob
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getActiveProcessingJob(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/active");
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
     * Retrieves a list of all inactive processing job types defined within a Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET GetInactiveJobTypes
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getProcessingJobInactiveJobTypes(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/inactive");
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
     * Retrieves inactive processing jobs at a specified Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET GetInactiveProcessingJob
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getInactiveProcessingJob(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/inactive");
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
     * Retrieves all processing job attributes available for a Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET GetJobTypesAttributes
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getProcessingJobJobTypesAttributes(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/attributes");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves all processing job categories available for a specified Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET GetJobTypesCategories
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getProcessingJobJobTypesCategories(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes/categories");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a ProcessingJob by Id.
     * Permissions Required:
     * - Manage Processing Job
     * GET GetProcessingJobById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getProcessingJobById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Initiates new processing jobs at a Facility, including job details and associated packages.
     * Permissions Required:
     * - Manage Processing Job
     * POST StartProcessingJob
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object startProcessingJobProcessingJob(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/start");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
     * Permissions Required:
     * - Manage Processing Job
     * PUT UnfinishProcessingJob
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object unfinishProcessingJobProcessingJob(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/unfinish");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
     * Permissions Required:
     * - Manage Processing Job
     * PUT UpdateJobTypes
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateProcessingJobJobTypes(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/processing/v2/jobtypes");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

