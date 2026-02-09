package client

import (
	"context"
	"net/url"
)


// POST /strains/v2
// Creates new strain records for a specified Facility.
// Permissions Required:
// - Manage Strains
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateStrains(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/strains/v2", queryParams, body)
}

// DELETE /strains/v2/{id}
// Archives an existing strain record for a Facility
// Permissions Required:
// - Manage Strains
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteStrainsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/strains/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /strains/v2/active
// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Strains
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetActiveStrains(
	ctx context.Context,lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastModifiedEnd != "" {
		queryParams["lastModifiedEnd"] = lastModifiedEnd
	}
	if lastModifiedStart != "" {
		queryParams["lastModifiedStart"] = lastModifiedStart
	}
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/strains/v2/active", queryParams, nil)
}

// GET /strains/v2/inactive
// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Strains
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetInactiveStrains(
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

	return c.send(ctx, "GET", "/strains/v2/inactive", queryParams, nil)
}

// GET /strains/v2/{id}
// Retrieves a Strain record by its Id, with an optional license number.
// Permissions Required:
// - Manage Strains
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetStrainsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/strains/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// PUT /strains/v2
// Updates existing strain records for a specified Facility.
// Permissions Required:
// - Manage Strains
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateStrains(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/strains/v2", queryParams, body)
}


