package client

import (
	"context"
)


// GET /employees/v2
// Retrieves a list of employees for a specified Facility.
// Permissions Required:
// - Manage Employees
// - View Employees
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetEmployees(
	ctx context.Context,licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/employees/v2", queryParams, nil)
}

// GET /employees/v2/permissions
// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
// Permissions Required:
// - Manage Employees
// Parameters:
// - employeeLicenseNumber (optional): Filter by employeeLicenseNumber
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetEmployeesPermissions(
	ctx context.Context,employeeLicenseNumber string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if employeeLicenseNumber != "" {
		queryParams["employeeLicenseNumber"] = employeeLicenseNumber
	}
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/employees/v2/permissions", queryParams, nil)
}


