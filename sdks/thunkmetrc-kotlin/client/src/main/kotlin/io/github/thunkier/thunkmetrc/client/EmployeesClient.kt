package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class EmployeesClient(private val client: MetrcClient) {
    /**
     * Retrieves a list of employees for a specified Facility.
     * Permissions Required:
     * - Manage Employees
     * - View Employees
     * GET GetEmployees
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getEmployees(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/employees/v2")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
     * Permissions Required:
     * - Manage Employees
     * GET GetPermissions
     * @param employeeLicenseNumber Query parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getEmployeesPermissions(employeeLicenseNumber: String? = null, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/employees/v2/permissions")
        val query = ArrayList<String>()
        if (employeeLicenseNumber != null) {
            query.add("employeeLicenseNumber=${URLEncoder.encode(employeeLicenseNumber, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

}

