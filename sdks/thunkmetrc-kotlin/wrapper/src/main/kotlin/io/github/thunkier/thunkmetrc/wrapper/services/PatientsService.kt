package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class PatientsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

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
    suspend fun createPatients(licenseNumber: String? = null, body: List<PatientsCreatePatientsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.patients.createPatients(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun deletePatientsById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.patients.deletePatientsById(id, licenseNumber, ) 
        } as? Any
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
    suspend fun getActivePatients(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Patient>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.patients.getActivePatients(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Patient>
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
    suspend fun getPatientsById(id: String, licenseNumber: String? = null): Patient? {
        return rateLimiter.execute(null,true,
        ) { 
            client.patients.getPatientsById(id, licenseNumber, ) 
        } as? Patient
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
    suspend fun updatePatients(licenseNumber: String? = null, body: List<PatientsUpdatePatientsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.patients.updatePatients(licenseNumber, body) 
        } as? WriteResponse
    }
}

