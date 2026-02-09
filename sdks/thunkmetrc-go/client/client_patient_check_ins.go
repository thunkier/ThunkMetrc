package client

import (
	"context"
	"net/url"
)


// POST /patient-checkins/v2
// Records patient check-ins for a specified Facility.
// Permissions Required:
// - ManagePatientsCheckIns
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePatientCheckIns(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/patient-checkins/v2", queryParams, body)
}

// DELETE /patient-checkins/v2/{id}
// Archives a Patient Check-In, identified by its Id, for a specified Facility.
// Permissions Required:
// - ManagePatientsCheckIns
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeletePatientCheckInsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/patient-checkins/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /patient-checkins/v2/locations
// Retrieves a list of Patient Check-In locations.
// Parameters:
func (c *MetrcClient) GetPatientCheckInsLocations(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/patient-checkins/v2/locations", queryParams, nil)
}

// GET /patient-checkins/v2
// Retrieves a list of patient check-ins for a specified Facility.
// Permissions Required:
// - ManagePatientsCheckIns
// Parameters:
// - checkinDateEnd (optional): Filter by checkinDateEnd
// - checkinDateStart (optional): Filter by checkinDateStart
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPatientCheckIns(
	ctx context.Context,checkinDateEnd string, checkinDateStart string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if checkinDateEnd != "" {
		queryParams["checkinDateEnd"] = checkinDateEnd
	}
	if checkinDateStart != "" {
		queryParams["checkinDateStart"] = checkinDateStart
	}
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/patient-checkins/v2", queryParams, nil)
}

// PUT /patient-checkins/v2
// Updates patient check-ins for a specified Facility.
// Permissions Required:
// - ManagePatientsCheckIns
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePatientCheckIns(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/patient-checkins/v2", queryParams, body)
}


