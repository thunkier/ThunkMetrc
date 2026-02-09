package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class ProcessingJobService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

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
    suspend fun createAdjustProcessingJob(licenseNumber: String? = null, body: List<ProcessingJobCreateAdjustProcessingJobRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.createAdjustProcessingJob(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun createJobTypes(licenseNumber: String? = null, body: List<ProcessingJobCreateJobTypesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.createProcessingJobJobTypes(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun createProcessingJobPackages(licenseNumber: String? = null, body: List<ProcessingJobCreateProcessingJobPackagesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.createProcessingJobPackages(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun deleteJobTypeById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.deleteProcessingJobJobTypeById(id, licenseNumber, ) 
        } as? Any
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
    suspend fun deleteProcessingJobById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.deleteProcessingJobById(id, licenseNumber, ) 
        } as? Any
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
    suspend fun finishProcessingJob(licenseNumber: String? = null, body: List<ProcessingJobFinishProcessingJobRequestItem>): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.finishProcessingJobProcessingJob(licenseNumber, body) 
        } as? Any
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
    suspend fun getActiveJobTypes(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<ActiveJobType>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.processingJob.getProcessingJobActiveJobTypes(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<ActiveJobType>
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
    suspend fun getActiveProcessingJob(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<ProcessingJob>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.processingJob.getActiveProcessingJob(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<ProcessingJob>
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
    suspend fun getInactiveJobTypes(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<InactiveJobType>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.processingJob.getProcessingJobInactiveJobTypes(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<InactiveJobType>
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
    suspend fun getInactiveProcessingJob(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<ProcessingJob>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.processingJob.getInactiveProcessingJob(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<ProcessingJob>
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
    suspend fun getJobTypesAttributes(licenseNumber: String? = null): PaginatedResponse<JobTypesAttribute>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.processingJob.getProcessingJobJobTypesAttributes(licenseNumber, ) 
        } as? PaginatedResponse<JobTypesAttribute>
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
    suspend fun getJobTypesCategories(licenseNumber: String? = null): PaginatedResponse<JobTypesCategory>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.processingJob.getProcessingJobJobTypesCategories(licenseNumber, ) 
        } as? PaginatedResponse<JobTypesCategory>
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
    suspend fun getProcessingJobById(id: String, licenseNumber: String? = null): ProcessingJob? {
        return rateLimiter.execute(null,true,
        ) { 
            client.processingJob.getProcessingJobById(id, licenseNumber, ) 
        } as? ProcessingJob
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
    suspend fun startProcessingJob(licenseNumber: String? = null, body: List<ProcessingJobStartProcessingJobRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.startProcessingJobProcessingJob(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun unfinishProcessingJob(licenseNumber: String? = null, body: List<ProcessingJobUnfinishProcessingJobRequestItem>): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.unfinishProcessingJobProcessingJob(licenseNumber, body) 
        } as? Any
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
    suspend fun updateJobTypes(licenseNumber: String? = null, body: List<ProcessingJobUpdateJobTypesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.processingJob.updateProcessingJobJobTypes(licenseNumber, body) 
        } as? WriteResponse
    }
}

