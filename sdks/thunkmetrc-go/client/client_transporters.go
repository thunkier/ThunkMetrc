package client

import (
	"context"
	"net/url"
)


// POST /transporters/v2/drivers
// Creates new driver records for a Facility.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateTransportersDrivers(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/transporters/v2/drivers", queryParams, body)
}

// POST /transporters/v2/vehicles
// Creates new vehicle records for a Facility.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateTransportersVehicles(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/transporters/v2/vehicles", queryParams, body)
}

// DELETE /transporters/v2/drivers/{id}
// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteTransportersDriverById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/transporters/v2/drivers/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE /transporters/v2/vehicles/{id}
// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteTransportersVehicleById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/transporters/v2/vehicles/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /transporters/v2/drivers/{id}
// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
// Permissions Required:
// - Transporters
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetTransportersDriverById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/transporters/v2/drivers/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /transporters/v2/drivers
// Retrieves a list of drivers for a Facility.
// Permissions Required:
// - Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransportersDrivers(
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

	return c.send(ctx, "GET", "/transporters/v2/drivers", queryParams, nil)
}

// GET /transporters/v2/vehicles/{id}
// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
// Permissions Required:
// - Transporters
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetTransportersVehicleById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/transporters/v2/vehicles/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /transporters/v2/vehicles
// Retrieves a list of vehicles for a Facility.
// Permissions Required:
// - Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransportersVehicles(
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

	return c.send(ctx, "GET", "/transporters/v2/vehicles", queryParams, nil)
}

// PUT /transporters/v2/drivers
// Updates existing driver records for a Facility.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateTransportersDrivers(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/transporters/v2/drivers", queryParams, body)
}

// PUT /transporters/v2/vehicles
// Updates existing vehicle records for a facility.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateTransportersVehicles(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/transporters/v2/vehicles", queryParams, body)
}


