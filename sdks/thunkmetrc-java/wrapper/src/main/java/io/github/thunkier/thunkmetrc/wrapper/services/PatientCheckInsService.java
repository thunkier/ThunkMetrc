package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class PatientCheckInsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public PatientCheckInsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Records patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * POST /patient-checkins/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createPatientCheckIns(
        String licenseNumber, List<PatientCheckInsCreatePatientCheckInsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.patientCheckIns.createPatientCheckIns(
                licenseNumber, body
            )
        );
    }

    /**
     * Archives a Patient Check-In, identified by its Id, for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * DELETE /patient-checkins/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deletePatientCheckInsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.patientCheckIns.deletePatientCheckInsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of Patient Check-In locations.
     * GET /patient-checkins/v2/locations
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<PatientCheckInsLocation> getPatientCheckInsLocations(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<PatientCheckInsLocation>) client.patientCheckIns.getPatientCheckInsLocations(
                
            )
        );
    }

    /**
     * Retrieves a list of patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * GET /patient-checkins/v2
     * @param checkinDateEnd Query parameter
     * @param checkinDateStart Query parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<PatientCheckIn> getPatientCheckIns(
        String checkinDateEnd, String checkinDateStart, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<PatientCheckIn>) client.patientCheckIns.getPatientCheckIns(
                checkinDateEnd, checkinDateStart, licenseNumber
            )
        );
    }

    /**
     * Updates patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * PUT /patient-checkins/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updatePatientCheckIns(
        String licenseNumber, List<PatientCheckInsUpdatePatientCheckInsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.patientCheckIns.updatePatientCheckIns(
                licenseNumber, body
            )
        );
    }

}

