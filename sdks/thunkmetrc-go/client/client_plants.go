package client

import (
	"context"
	"net/url"
)


// POST /plants/v2/additives/bylocation
// Records additive usage for plants based on their location within a specified Facility.
// Permissions Required:
// - Manage Plants
// - Manage Plants Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantsAdditivesByLocation(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plants/v2/additives/bylocation", queryParams, body)
}

// POST /plants/v2/additives/bylocation/usingtemplate
// Records additive usage for plants by location using a predefined additive template at a specified Facility.
// Permissions Required:
// - Manage Plants Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantsAdditivesByLocationUsingTemplate(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plants/v2/additives/bylocation/usingtemplate", queryParams, body)
}

// POST /plants/v2/manicure
// Creates harvest product records from plant batches at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manicure/Harvest Veg/Flower Plants
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantsManicure(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plants/v2/manicure", queryParams, body)
}

// POST /plants/v2/plantbatch/packages
// Creates packages from plant batches at a specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantsPlantBatchPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plants/v2/plantbatch/packages", queryParams, body)
}

// POST /plants/v2/additives
// Records additive usage details applied to specific plants at a Facility.
// Permissions Required:
// - Manage Plants Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantsAdditives(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plants/v2/additives", queryParams, body)
}

// POST /plants/v2/additives/usingtemplate
// Records additive usage for plants using predefined additive templates at a specified Facility.
// Permissions Required:
// - Manage Plants Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantsAdditivesUsingTemplate(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plants/v2/additives/usingtemplate", queryParams, body)
}

// POST /plants/v2/plantings
// Creates new plant batches at a specified Facility from existing plant data.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantsPlantings(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plants/v2/plantings", queryParams, body)
}

// POST /plants/v2/waste
// Records waste events for plants at a Facility, including method, reason, and location details.
// Permissions Required:
// - Manage Plants Waste
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePlantsWaste(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/plants/v2/waste", queryParams, body)
}

// DELETE /plants/v2
// Removes plants from a Facility’s inventory while recording the reason for their disposal.
// Permissions Required:
// - View Veg/Flower Plants
// - Destroy Veg/Flower Plants
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeletePlants(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/plants/v2", queryParams, nil)
}

// GET /plants/v2/additives
// Retrieves additive records applied to plants at a specified Facility.
// Permissions Required:
// - View/Manage Plants Additives
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPlantsAdditives(
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

	return c.send(ctx, "GET", "/plants/v2/additives", queryParams, nil)
}

// GET /plants/v2/additives/types
// Retrieves a list of all plant additive types defined within a Facility.
// Parameters:
func (c *MetrcClient) GetPlantsAdditivesTypes(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/plants/v2/additives/types", queryParams, nil)
}

// GET /plants/v2/flowering
// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
// Permissions Required:
// - View Veg/Flower Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetFloweringPlants(
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

	return c.send(ctx, "GET", "/plants/v2/flowering", queryParams, nil)
}

// GET /plants/v2/growthphases
// Retrieves the list of growth phases supported by a specified Facility.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPlantsGrowthPhases(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/plants/v2/growthphases", queryParams, nil)
}

// GET /plants/v2/inactive
// Retrieves inactive plants at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetInactivePlants(
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

	return c.send(ctx, "GET", "/plants/v2/inactive", queryParams, nil)
}

// GET /plants/v2/mother/inactive
// Retrieves inactive mother-phase plants at a specified Facility.
// Permissions Required:
// - View Mother Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetMotherInactivePlants(
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

	return c.send(ctx, "GET", "/plants/v2/mother/inactive", queryParams, nil)
}

// GET /plants/v2/mother/onhold
// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
// Permissions Required:
// - View Mother Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetMotherOnHoldPlants(
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

	return c.send(ctx, "GET", "/plants/v2/mother/onhold", queryParams, nil)
}

// GET /plants/v2/mother
// Retrieves mother-phase plants at a specified Facility.
// Permissions Required:
// - View Mother Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetMotherPlants(
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

	return c.send(ctx, "GET", "/plants/v2/mother", queryParams, nil)
}

// GET /plants/v2/onhold
// Retrieves plants that are currently on hold at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetOnHoldPlants(
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

	return c.send(ctx, "GET", "/plants/v2/onhold", queryParams, nil)
}

// GET /plants/v2/{id}
// Retrieves a Plant by Id.
// Permissions Required:
// - View Veg/Flower Plants
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPlantsById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/plants/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /plants/v2/{label}
// Retrieves a Plant by label.
// Permissions Required:
// - View Veg/Flower Plants
// Parameters:
// - label (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPlantsByLabel(
	ctx context.Context,label string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/plants/v2/"+url.QueryEscape(label)+"", queryParams, nil)
}

// GET /plants/v2/waste
// Retrieves a list of recorded plant waste events for a specific Facility.
// Permissions Required:
// - View Plants Waste
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPlantsWaste(
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

	return c.send(ctx, "GET", "/plants/v2/waste", queryParams, nil)
}

// GET /plants/v2/waste/methods/all
// Retrieves a list of all available plant waste methods for use within a Facility.
// Parameters:
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPlantsWasteMethods(
	ctx context.Context,pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/plants/v2/waste/methods/all", queryParams, nil)
}

// GET /plants/v2/waste/reasons
// Retriveves available reasons for recording mature plant waste at a specified Facility.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPlantsWasteReasons(
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

	return c.send(ctx, "GET", "/plants/v2/waste/reasons", queryParams, nil)
}

// GET /plants/v2/vegetative
// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
// Permissions Required:
// - View Veg/Flower Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetVegetativePlants(
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

	return c.send(ctx, "GET", "/plants/v2/vegetative", queryParams, nil)
}

// GET /plants/v2/waste/{id}/plant
// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
// Permissions Required:
// - View Plants Waste
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPlantsWasteById(
	ctx context.Context,id string, licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
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

	return c.send(ctx, "GET", "/plants/v2/waste/"+url.QueryEscape(id)+"/plant", queryParams, nil)
}

// GET /plants/v2/waste/{id}/package
// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
// Permissions Required:
// - View Plants Waste
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPlantsWastePackageById(
	ctx context.Context,id string, licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
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

	return c.send(ctx, "GET", "/plants/v2/waste/"+url.QueryEscape(id)+"/package", queryParams, nil)
}

// PUT /plants/v2/adjust
// Adjusts the recorded count of plants at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateAdjustPlants(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plants/v2/adjust", queryParams, body)
}

// PUT /plants/v2/growthphase
// Changes the growth phases of plants within a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantsGrowthPhase(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plants/v2/growthphase", queryParams, body)
}

// PUT /plants/v2/harvest
// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
// Permissions Required:
// - View Veg/Flower Plants
// - Manicure/Harvest Veg/Flower Plants
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantsHarvest(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plants/v2/harvest", queryParams, body)
}

// PUT /plants/v2/merge
// Merges multiple plant groups into a single group within a Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manicure/Harvest Veg/Flower Plants
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantsMerge(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plants/v2/merge", queryParams, body)
}

// PUT /plants/v2/location
// Moves plant batches to new locations within a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantsLocation(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plants/v2/location", queryParams, body)
}

// PUT /plants/v2/strain
// Updates the strain information for plants within a Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantsStrain(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plants/v2/strain", queryParams, body)
}

// PUT /plants/v2/tag
// Replaces existing plant tags with new tags for plants within a Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantsTag(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plants/v2/tag", queryParams, body)
}

// PUT /plants/v2/split
// Splits an existing plant group into multiple groups within a Facility.
// Permissions Required:
// - View Plant
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePlantsSplit(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/plants/v2/split", queryParams, body)
}


