package client

import (
	"context"
	"net/url"
)


// POST /plantbatches/v2/adjust
// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateAdjustPlantBatches(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/adjust", queryParams, body)
}

// POST /plantbatches/v2/growthphase
// Updates the growth phase of plants at a specified Facility based on tracking information.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantBatchesGrowthPhase(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/growthphase", queryParams, body)
}

// POST /plantbatches/v2/packages/frommotherplant
// Creates packages from mother plants at the specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantBatchesPackagesFromMotherPlant(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/packages/frommotherplant", queryParams, body)
}

// POST /plantbatches/v2/additives
// Records Additive usage details for plant batches at a specific Facility.
// Permissions Required:
// - Manage Plants Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantBatchesAdditives(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/additives", queryParams, body)
}

// POST /plantbatches/v2/additives/usingtemplate
// Records Additive usage for plant batches at a Facility using predefined additive templates.
// Permissions Required:
// - Manage Plants Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantBatchesAdditivesUsingTemplate(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/additives/usingtemplate", queryParams, body)
}

// POST /plantbatches/v2/packages
// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantBatchesPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/packages", queryParams, body)
}

// POST /plantbatches/v2/plantings
// Creates new plantings for a Facility by generating plant batches based on provided planting details.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantBatchesPlantings(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/plantings", queryParams, body)
}

// POST /plantbatches/v2/waste
// Records waste information for plant batches based on the submitted data for the specified Facility.
// Permissions Required:
// - Manage Plants Waste
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantBatchesWaste(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/waste", queryParams, body)
}

// POST /plantbatches/v2/split
// Splits an existing Plant Batch into multiple groups at the specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantBatchesSplit(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plantbatches/v2/split", queryParams, body)
}

// DELETE /plantbatches/v2
// Completes the destruction of plant batches based on the provided input data.
// Permissions Required:
// - View Immature Plants
// - Destroy Immature Plants
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeletePlantBatches(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/plantbatches/v2", queryParams, nil)
}

// GET /plantbatches/v2/active
// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
// Permissions Required:
// - View Immature Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetActivePlantBatches(
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

	return c.send(ctx, "GET", "/plantbatches/v2/active", queryParams, nil)
}

// GET /plantbatches/v2/inactive
// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
// Permissions Required:
// - View Immature Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetInactivePlantBatches(
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

	return c.send(ctx, "GET", "/plantbatches/v2/inactive", queryParams, nil)
}

// GET /plantbatches/v2/{id}
// Retrieves a Plant Batch by Id.
// Permissions Required:
// - View Immature Plants
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPlantBatchesById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/plantbatches/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /plantbatches/v2/types
// Retrieves a list of plant batch types.
// Parameters:
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPlantBatchesTypes(
	ctx context.Context,pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/plantbatches/v2/types", queryParams, nil)
}

// GET /plantbatches/v2/waste
// Retrieves waste details associated with plant batches at a specified Facility.
// Permissions Required:
// - View Plants Waste
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPlantBatchesWaste(
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

	return c.send(ctx, "GET", "/plantbatches/v2/waste", queryParams, nil)
}

// GET /plantbatches/v2/waste/reasons
// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPlantBatchesWasteReasons(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/plantbatches/v2/waste/reasons", queryParams, nil)
}

// PUT /plantbatches/v2/name
// Renames plant batches at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantBatchesName(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plantbatches/v2/name", queryParams, body)
}

// PUT /plantbatches/v2/location
// Moves one or more plant batches to new locations with in a specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantBatchesLocation(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plantbatches/v2/location", queryParams, body)
}

// PUT /plantbatches/v2/strain
// Changes the strain of plant batches at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantBatchesStrain(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plantbatches/v2/strain", queryParams, body)
}

// PUT /plantbatches/v2/tag
// Replaces tags for plant batches at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantBatchesTag(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plantbatches/v2/tag", queryParams, body)
}


