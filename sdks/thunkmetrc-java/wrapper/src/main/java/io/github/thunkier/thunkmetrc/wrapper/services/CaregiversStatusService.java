package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class CaregiversStatusService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public CaregiversStatusService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * Permissions Required:
     * - Lookup Caregivers
     * GET /caregivers/v2/status/{caregiverLicenseNumber}
     * @param caregiverLicenseNumber Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<CaregiversStatus> getCaregiversStatusByCaregiverLicenseNumber(
        String caregiverLicenseNumber, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<CaregiversStatus>) client.caregiversStatus.getCaregiversStatusByCaregiverLicenseNumber(
                caregiverLicenseNumber, licenseNumber
            )
        );
    }

}

