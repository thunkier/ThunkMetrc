package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class PatientCheckInsClient(private val client: MetrcClient) {
    /**
     * Records patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * POST CreatePatientCheckIns
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createPatientCheckIns(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patient-checkins/v2")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
    }

    /**
     * Archives a Patient Check-In, identified by its Id, for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * DELETE DeletePatientCheckInsById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun deletePatientCheckInsById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/patient-checkins/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Retrieves a list of Patient Check-In locations.
     * GET GetLocations
     * @throws IOException If request fails
     * @return Response object
     */
    fun getPatientCheckInsLocations(): Any? {
        val path = StringBuilder("/patient-checkins/v2/locations")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a list of patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * GET GetPatientCheckIns
     * @param checkinDateEnd Query parameter
     * @param checkinDateStart Query parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getPatientCheckIns(checkinDateEnd: String? = null, checkinDateStart: String? = null, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/patient-checkins/v2")
        val query = ArrayList<String>()
        if (checkinDateEnd != null) {
            query.add("checkinDateEnd=${URLEncoder.encode(checkinDateEnd, StandardCharsets.UTF_8)}")
        }
        if (checkinDateStart != null) {
            query.add("checkinDateStart=${URLEncoder.encode(checkinDateStart, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * PUT UpdatePatientCheckIns
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updatePatientCheckIns(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/patient-checkins/v2")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

}

