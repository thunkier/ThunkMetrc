package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class ProcessingJobService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public ProcessingJobService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
     * Permissions Required:
     * - Manage Processing Job
     * POST /processing/v2/adjust
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createAdjustProcessingJob(
        String licenseNumber, List<ProcessingJobCreateAdjustProcessingJobRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.processingJob.createAdjustProcessingJob(
                licenseNumber, body
            )
        );
    }

    /**
     * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
     * Permissions Required:
     * - Manage Processing Job
     * POST /processing/v2/jobtypes
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createProcessingJobJobTypes(
        String licenseNumber, List<ProcessingJobCreateJobTypesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.processingJob.createProcessingJobJobTypes(
                licenseNumber, body
            )
        );
    }

    /**
     * Creates packages from processing jobs at a Facility, including optional location and note assignments.
     * Permissions Required:
     * - Manage Processing Job
     * POST /processing/v2/createpackages
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createProcessingJobPackages(
        String licenseNumber, List<ProcessingJobCreateProcessingJobPackagesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.processingJob.createProcessingJobPackages(
                licenseNumber, body
            )
        );
    }

    /**
     * Archives a Processing Job Type at a Facility, making it inactive for future use.
     * Permissions Required:
     * - Manage Processing Job
     * DELETE /processing/v2/jobtypes/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteProcessingJobJobTypeById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.processingJob.deleteProcessingJobJobTypeById(
                id, licenseNumber
            )
        );
    }

    /**
     * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
     * Permissions Required:
     * - Manage Processing Job
     * DELETE /processing/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteProcessingJobById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.processingJob.deleteProcessingJobById(
                id, licenseNumber
            )
        );
    }

    /**
     * Completes processing jobs at a Facility by recording final notes and waste measurements.
     * Permissions Required:
     * - Manage Processing Job
     * PUT /processing/v2/finish
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public Object finishProcessingJobProcessingJob(
        String licenseNumber, List<ProcessingJobFinishProcessingJobRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.processingJob.finishProcessingJobProcessingJob(
                licenseNumber, body
            )
        );
    }

    /**
     * Retrieves a list of all active processing job types defined within a Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET /processing/v2/jobtypes/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<ActiveJobType> getProcessingJobActiveJobTypes(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<ActiveJobType>) client.processingJob.getProcessingJobActiveJobTypes(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves active processing jobs at a specified Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET /processing/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<ProcessingJob> getActiveProcessingJob(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<ProcessingJob>) client.processingJob.getActiveProcessingJob(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of all inactive processing job types defined within a Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET /processing/v2/jobtypes/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<InactiveJobType> getProcessingJobInactiveJobTypes(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<InactiveJobType>) client.processingJob.getProcessingJobInactiveJobTypes(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves inactive processing jobs at a specified Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET /processing/v2/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<ProcessingJob> getInactiveProcessingJob(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<ProcessingJob>) client.processingJob.getInactiveProcessingJob(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves all processing job attributes available for a Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET /processing/v2/jobtypes/attributes
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<JobTypesAttribute> getProcessingJobJobTypesAttributes(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<JobTypesAttribute>) client.processingJob.getProcessingJobJobTypesAttributes(
                licenseNumber
            )
        );
    }

    /**
     * Retrieves all processing job categories available for a specified Facility.
     * Permissions Required:
     * - Manage Processing Job
     * GET /processing/v2/jobtypes/categories
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<JobTypesCategory> getProcessingJobJobTypesCategories(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<JobTypesCategory>) client.processingJob.getProcessingJobJobTypesCategories(
                licenseNumber
            )
        );
    }

    /**
     * Retrieves a ProcessingJob by Id.
     * Permissions Required:
     * - Manage Processing Job
     * GET /processing/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public ProcessingJob getProcessingJobById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (ProcessingJob) client.processingJob.getProcessingJobById(
                id, licenseNumber
            )
        );
    }

    /**
     * Initiates new processing jobs at a Facility, including job details and associated packages.
     * Permissions Required:
     * - Manage Processing Job
     * POST /processing/v2/start
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse startProcessingJobProcessingJob(
        String licenseNumber, List<ProcessingJobStartProcessingJobRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.processingJob.startProcessingJobProcessingJob(
                licenseNumber, body
            )
        );
    }

    /**
     * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
     * Permissions Required:
     * - Manage Processing Job
     * PUT /processing/v2/unfinish
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public Object unfinishProcessingJobProcessingJob(
        String licenseNumber, List<ProcessingJobUnfinishProcessingJobRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.processingJob.unfinishProcessingJobProcessingJob(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
     * Permissions Required:
     * - Manage Processing Job
     * PUT /processing/v2/jobtypes
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateProcessingJobJobTypes(
        String licenseNumber, List<ProcessingJobUpdateJobTypesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.processingJob.updateProcessingJobJobTypes(
                licenseNumber, body
            )
        );
    }

}

