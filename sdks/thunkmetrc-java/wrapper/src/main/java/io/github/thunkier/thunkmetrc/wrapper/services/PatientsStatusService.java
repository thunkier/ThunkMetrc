package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class PatientsStatusService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public PatientsStatusService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * Permissions Required:
     * - Lookup Patients
     * GET /patients/v2/statuses/{patientLicenseNumber}
     * @param patientLicenseNumber Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<PatientsStatus> getPatientsStatusesByPatientLicenseNumber(
        String patientLicenseNumber, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<PatientsStatus>) client.patientsStatus.getPatientsStatusesByPatientLicenseNumber(
                patientLicenseNumber, licenseNumber
            )
        );
    }

}

