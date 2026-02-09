package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class LabTestsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public LabTestsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

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
    public WriteResponse createLabTestsRecord(
        String licenseNumber, List<LabTestsCreateRecordRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.labTests.createLabTestsRecord(
                licenseNumber, body
            )
        );
    }

    /**
     * Retrieves a list of Lab Test batches.
     * GET /labtests/v2/batches
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Batch> getLabTestsBatches(
        String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Batch>) client.labTests.getLabTestsBatches(
                pageNumber, pageSize
            )
        );
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
    public Object getLabTestsLabTestDocumentById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.labTests.getLabTestsLabTestDocumentById(
                id, licenseNumber
            )
        );
    }

    /**
     * Returns a list of Lab Test types.
     * GET /labtests/v2/types
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<LabTestsType> getLabTestsTypes(
        String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<LabTestsType>) client.labTests.getLabTestsTypes(
                pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Result> getLabTestsResults(
        String licenseNumber, String packageId, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Result>) client.labTests.getLabTestsResults(
                licenseNumber, packageId, pageNumber, pageSize
            )
        );
    }

    /**
     * Returns a list of all lab testing states.
     * GET /labtests/v2/states
     * @throws Exception If request fails
     * @return Response object
     */
    public Object getLabTestsStates(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.labTests.getLabTestsStates(
                
            )
        );
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
    public WriteResponse updateLabTestsLabTestDocument(
        String licenseNumber, List<LabTestsUpdateLabTestDocumentRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.labTests.updateLabTestsLabTestDocument(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateLabTestsResultsRelease(
        String licenseNumber, List<LabTestsUpdateResultsReleaseRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.labTests.updateLabTestsResultsRelease(
                licenseNumber, body
            )
        );
    }

}

