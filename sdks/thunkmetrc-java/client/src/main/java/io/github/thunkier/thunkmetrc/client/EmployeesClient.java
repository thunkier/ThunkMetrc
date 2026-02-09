package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class EmployeesClient {
    private final MetrcClient client;

    public EmployeesClient(MetrcClient client) {
        this.client = client;
    }
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
    public Object getEmployees(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/employees/v2");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
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
    public Object getEmployeesPermissions(
        String employeeLicenseNumber, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/employees/v2/permissions");
        StringBuilder query = new StringBuilder();
        if (employeeLicenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("employeeLicenseNumber=").append(URLEncoder.encode(employeeLicenseNumber, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

}

