package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class EmployeesService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Retrieves a list of employees for a specified Facility.
     * Permissions Required:
     * - Manage Employees
     * - View Employees
     * GET /employees/v2
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getEmployees(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Employee>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.employees.getEmployees(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Employee>
    }

    /**
     * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
     * Permissions Required:
     * - Manage Employees
     * GET /employees/v2/permissions
     * @param employeeLicenseNumber Query parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPermissions(employeeLicenseNumber: String? = null, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.employees.getEmployeesPermissions(employeeLicenseNumber, licenseNumber, ) 
        } as? Any
    }
}

