package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class PatientCheckInsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

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
    suspend fun createPatientCheckIns(licenseNumber: String? = null, body: List<PatientCheckInsCreatePatientCheckInsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.patientCheckIns.createPatientCheckIns(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun deletePatientCheckInsById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.patientCheckIns.deletePatientCheckInsById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Retrieves a list of Patient Check-In locations.
     * GET /patient-checkins/v2/locations
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getLocations(): List<PatientCheckInsLocation>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.patientCheckIns.getPatientCheckInsLocations() 
        } as? List<PatientCheckInsLocation>
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
    suspend fun getPatientCheckIns(checkinDateEnd: String? = null, checkinDateStart: String? = null, licenseNumber: String? = null): PaginatedResponse<PatientCheckIn>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.patientCheckIns.getPatientCheckIns(checkinDateEnd, checkinDateStart, licenseNumber, ) 
        } as? PaginatedResponse<PatientCheckIn>
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
    suspend fun updatePatientCheckIns(licenseNumber: String? = null, body: List<PatientCheckInsUpdatePatientCheckInsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.patientCheckIns.updatePatientCheckIns(licenseNumber, body) 
        } as? WriteResponse
    }
}

