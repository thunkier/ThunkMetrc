package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class CaregiversStatusService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

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
    suspend fun getCaregiversStatusByCaregiverLicenseNumber(caregiverLicenseNumber: String, licenseNumber: String? = null): List<CaregiversStatus>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.caregiversStatus.getCaregiversStatusByCaregiverLicenseNumber(caregiverLicenseNumber, licenseNumber, ) 
        } as? List<CaregiversStatus>
    }
}

