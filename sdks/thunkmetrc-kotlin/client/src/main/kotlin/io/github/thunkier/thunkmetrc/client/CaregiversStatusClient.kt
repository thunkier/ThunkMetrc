package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class CaregiversStatusClient(private val client: MetrcClient) {
    /**
     * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * Permissions Required:
     * - Lookup Caregivers
     * GET GetCaregiversStatusByCaregiverLicenseNumber
     * @param caregiverLicenseNumber Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getCaregiversStatusByCaregiverLicenseNumber(caregiverLicenseNumber: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/caregivers/v2/status/${URLEncoder.encode(caregiverLicenseNumber, StandardCharsets.UTF_8)}")
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

