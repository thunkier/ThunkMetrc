package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class PatientsStatusClient(private val client: MetrcClient) {
    /**
     * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * Permissions Required:
     * - Lookup Patients
     * GET GetPatientsStatusesByPatientLicenseNumber
     * @param patientLicenseNumber Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getPatientsStatusesByPatientLicenseNumber(patientLicenseNumber: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/patients/v2/statuses/${URLEncoder.encode(patientLicenseNumber, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

}

