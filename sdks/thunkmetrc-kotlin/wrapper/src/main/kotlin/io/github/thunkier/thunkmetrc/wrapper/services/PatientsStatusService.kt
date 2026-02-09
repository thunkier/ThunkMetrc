package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class PatientsStatusService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

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
    suspend fun getPatientsStatusesByPatientLicenseNumber(patientLicenseNumber: String, licenseNumber: String? = null): List<PatientsStatus>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.patientsStatus.getPatientsStatusesByPatientLicenseNumber(patientLicenseNumber, licenseNumber, ) 
        } as? List<PatientsStatus>
    }
}

