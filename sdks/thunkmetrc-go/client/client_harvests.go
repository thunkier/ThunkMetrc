package client

import (
	"context"
	"net/url"
)


// POST /harvests/v2/packages
// Creates packages from harvested products for a specified Facility.
// Permissions Required:
// - View Harvests
// - Manage Harvests
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateHarvestsPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/harvests/v2/packages", queryParams, body)
}

// POST /harvests/v2/waste
// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
// Permissions Required:
// - View Harvests
// - Manage Harvests
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateHarvestsWaste(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/harvests/v2/waste", queryParams, body)
}

// POST /harvests/v2/packages/testing
// Creates packages for testing from harvested products for a specified Facility.
// Permissions Required:
// - View Harvests
// - Manage Harvests
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateHarvestsPackagesTesting(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/harvests/v2/packages/testing", queryParams, body)
}

// DELETE /harvests/v2/waste/{id}
// Discontinues a specific harvest waste record by Id for the specified Facility.
// Permissions Required:
// - View Harvests
// - Discontinue Harvest Waste
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteHarvestsWasteById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/harvests/v2/waste/"+url.QueryEscape(id)+"", queryParams, nil)
}

// PUT /harvests/v2/finish
// Marks one or more harvests as finished for the specified Facility.
// Permissions Required:
// - View Harvests
// - Finish/Discontinue Harvests
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) FinishHarvestsHarvests(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/harvests/v2/finish", queryParams, body)
}

// GET /harvests/v2/active
// Retrieves a list of active harvests for a specified Facility.
// Permissions Required:
// - View Harvests
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetActiveHarvests(
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

	return c.send(ctx, "GET", "/harvests/v2/active", queryParams, nil)
}

// GET /harvests/v2/{id}
// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
// Permissions Required:
// - View Harvests
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetHarvestsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/harvests/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /harvests/v2/waste
// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
// Permissions Required:
// - View Harvests
// Parameters:
// - harvestId (optional): Filter by harvestId
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetHarvestsWaste(
	ctx context.Context,harvestId string, licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if harvestId != "" {
		queryParams["harvestId"] = harvestId
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

	return c.send(ctx, "GET", "/harvests/v2/waste", queryParams, nil)
}

// GET /harvests/v2/inactive
// Retrieves a list of inactive harvests for a specified Facility.
// Permissions Required:
// - View Harvests
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetInactiveHarvests(
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

	return c.send(ctx, "GET", "/harvests/v2/inactive", queryParams, nil)
}

// GET /harvests/v2/onhold
// Retrieves a list of harvests on hold for a specified Facility.
// Permissions Required:
// - View Harvests
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetOnHoldHarvests(
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

	return c.send(ctx, "GET", "/harvests/v2/onhold", queryParams, nil)
}

// GET /harvests/v2/waste/types
// Retrieves a list of Waste types for harvests.
// Parameters:
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetHarvestsWasteTypes(
	ctx context.Context,pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/harvests/v2/waste/types", queryParams, nil)
}

// PUT /harvests/v2/unfinish
// Reopens one or more previously finished harvests for the specified Facility.
// Permissions Required:
// - View Harvests
// - Finish/Discontinue Harvests
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UnfinishHarvestsHarvests(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/harvests/v2/unfinish", queryParams, body)
}

// PUT /harvests/v2/location
// Updates the Location of Harvest for a specified Facility.
// Permissions Required:
// - View Harvests
// - Manage Harvests
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateHarvestsLocation(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/harvests/v2/location", queryParams, body)
}

// PUT /harvests/v2/rename
// Renames one or more harvests for the specified Facility.
// Permissions Required:
// - View Harvests
// - Manage Harvests
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateHarvestsRename(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/harvests/v2/rename", queryParams, body)
}

// PUT /harvests/v2/restore/harvestedplants
// Restores previously harvested plants to their original state for the specified Facility.
// Permissions Required:
// - View Harvests
// - Finish/Discontinue Harvests
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateHarvestsRestoreHarvestedPlants(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/harvests/v2/restore/harvestedplants", queryParams, body)
}


