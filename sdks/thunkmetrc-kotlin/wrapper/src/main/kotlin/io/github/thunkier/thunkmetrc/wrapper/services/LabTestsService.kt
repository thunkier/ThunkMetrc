package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class LabTestsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * POST /labtests/v2/record
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createRecord(licenseNumber: String? = null, body: List<LabTestsCreateRecordRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.labTests.createLabTestsRecord(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Retrieves a list of Lab Test batches.
     * GET /labtests/v2/batches
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getBatches(pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Batch>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.labTests.getLabTestsBatches(pageNumber, pageSize, ) 
        } as? PaginatedResponse<Batch>
    }

    /**
     * Retrieves a specific Lab Test result document by its Id for a given Facility.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * GET /labtests/v2/labtestdocument/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getLabTestDocumentById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.labTests.getLabTestsLabTestDocumentById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Returns a list of Lab Test types.
     * GET /labtests/v2/types
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getLabTestsTypes(pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<LabTestsType>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.labTests.getLabTestsTypes(pageNumber, pageSize, ) 
        } as? PaginatedResponse<LabTestsType>
    }

    /**
     * Retrieves Lab Test results for a specified Package.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * GET /labtests/v2/results
     * @param licenseNumber Query parameter
     * @param packageId Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getResults(licenseNumber: String? = null, packageId: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Result>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.labTests.getLabTestsResults(licenseNumber, packageId, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Result>
    }

    /**
     * Returns a list of all lab testing states.
     * GET /labtests/v2/states
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getStates(): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.labTests.getLabTestsStates() 
        } as? Any
    }

    /**
     * Updates one or more documents for previously submitted lab tests.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /labtests/v2/labtestdocument
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateLabTestDocument(licenseNumber: String? = null, body: List<LabTestsUpdateLabTestDocumentRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.labTests.updateLabTestsLabTestDocument(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Releases Lab Test results for one or more packages.
     * Permissions Required:
     * - View Packages
     * - Manage Packages Inventory
     * PUT /labtests/v2/results/release
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateResultsRelease(licenseNumber: String? = null, body: List<LabTestsUpdateResultsReleaseRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.labTests.updateLabTestsResultsRelease(licenseNumber, body) 
        } as? WriteResponse
    }
}

