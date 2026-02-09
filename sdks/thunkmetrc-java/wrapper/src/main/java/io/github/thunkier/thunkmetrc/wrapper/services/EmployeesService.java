package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class EmployeesService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public EmployeesService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Employee> getEmployees(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Employee>) client.employees.getEmployees(
                licenseNumber, pageNumber, pageSize
            )
        );
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
    public Object getEmployeesPermissions(
        String employeeLicenseNumber, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.employees.getEmployeesPermissions(
                employeeLicenseNumber, licenseNumber
            )
        );
    }

}

