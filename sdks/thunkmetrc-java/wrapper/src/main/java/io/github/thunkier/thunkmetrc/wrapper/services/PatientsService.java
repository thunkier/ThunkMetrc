package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class PatientsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public PatientsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Adds new patients to a specified Facility.
     * Permissions Required:
     * - Manage Patients
     * POST /patients/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createPatients(
        String licenseNumber, List<PatientsCreatePatientsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.patients.createPatients(
                licenseNumber, body
            )
        );
    }

    /**
     * Removes a Patient, identified by an Id, from a specified Facility.
     * Permissions Required:
     * - Manage Patients
     * DELETE /patients/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deletePatientsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.patients.deletePatientsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of active patients for a specified Facility.
     * Permissions Required:
     * - Manage Patients
     * GET /patients/v2/active
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Patient> getActivePatients(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Patient>) client.patients.getActivePatients(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a Patient by Id.
     * Permissions Required:
     * - Manage Patients
     * GET /patients/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Patient getPatientsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Patient) client.patients.getPatientsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Updates Patient information for a specified Facility.
     * Permissions Required:
     * - Manage Patients
     * PUT /patients/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePatients(
        String licenseNumber, List<PatientsUpdatePatientsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.patients.updatePatients(
                licenseNumber, body
            )
        );
    }

}

