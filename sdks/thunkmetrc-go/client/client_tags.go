package client

import (
	"context"
)


// GET /tags/v2/package/available
// Returns a list of available package tags. NOTE: This is a premium endpoint.
// Permissions Required:
// - Manage Tags
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetTagsAvailablePackage(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/tags/v2/package/available", queryParams, nil)
}

// GET /tags/v2/plant/available
// Returns a list of available plant tags. NOTE: This is a premium endpoint.
// Permissions Required:
// - Manage Tags
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetTagsAvailablePlant(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/tags/v2/plant/available", queryParams, nil)
}

// GET /tags/v2/staged
// Returns a list of staged tags. NOTE: This is a premium endpoint.
// Permissions Required:
// - Manage Tags
// - RetailId.AllowPackageStaging Key Value enabled
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetStagedTags(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/tags/v2/staged", queryParams, nil)
}


