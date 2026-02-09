package client

import (
	"context"
	"net/url"
)


// POST /sublocations/v2
// Creates new sublocation records for a Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateSublocations(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/sublocations/v2", queryParams, body)
}

// DELETE /sublocations/v2/{id}
// Archives an existing Sublocation record for a Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteSublocationsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/sublocations/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /sublocations/v2/active
// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Locations
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetActiveSublocations(
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

	return c.send(ctx, "GET", "/sublocations/v2/active", queryParams, nil)
}

// GET /sublocations/v2/inactive
// Retrieves a list of inactive sublocations for the specified Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetInactiveSublocations(
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

	return c.send(ctx, "GET", "/sublocations/v2/inactive", queryParams, nil)
}

// GET /sublocations/v2/{id}
// Retrieves a Sublocation by its Id, with an optional license number.
// Permissions Required:
// - Manage Locations
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetSublocationsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/sublocations/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// PUT /sublocations/v2
// Updates existing sublocation records for a specified Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSublocations(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sublocations/v2", queryParams, body)
}


