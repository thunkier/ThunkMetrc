package client

import (
	"context"
	"net/url"
)


// POST /patients/v2
// Adds new patients to a specified Facility.
// Permissions Required:
// - Manage Patients
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePatients(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/patients/v2", queryParams, body)
}

// DELETE /patients/v2/{id}
// Removes a Patient, identified by an Id, from a specified Facility.
// Permissions Required:
// - Manage Patients
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeletePatientsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/patients/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /patients/v2/active
// Retrieves a list of active patients for a specified Facility.
// Permissions Required:
// - Manage Patients
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetActivePatients(
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

	return c.send(ctx, "GET", "/patients/v2/active", queryParams, nil)
}

// GET /patients/v2/{id}
// Retrieves a Patient by Id.
// Permissions Required:
// - Manage Patients
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPatientsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/patients/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// PUT /patients/v2
// Updates Patient information for a specified Facility.
// Permissions Required:
// - Manage Patients
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePatients(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/patients/v2", queryParams, body)
}


