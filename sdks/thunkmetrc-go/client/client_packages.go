package client

import (
	"context"
	"net/url"
)


// POST /packages/v2/adjust
// Records a list of adjustments for packages at a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateAdjustPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/packages/v2/adjust", queryParams, body)
}

// POST /packages/v2
// Creates new packages for a specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePackagesPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/packages/v2", queryParams, body)
}

// POST /packages/v2/plantings
// Creates new plantings from packages for a specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePackagesPlantings(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/packages/v2/plantings", queryParams, body)
}

// POST /packages/v2/testing
// Creates new packages for testing for a specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreatePackagesTesting(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/packages/v2/testing", queryParams, body)
}

// DELETE /packages/v2/{id}
// Discontinues a Package at a specific Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeletePackagesById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/packages/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// PUT /packages/v2/finish
// Updates a list of packages as finished for a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) FinishPackagesPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/finish", queryParams, body)
}

// PUT /packages/v2/finishedgood/flag
// Flags one or more Packages at the specified Facility as Finished Goods.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) FinishedgoodFlagPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/finishedgood/flag", queryParams, body)
}

// PUT /packages/v2/finishedgood/unflag
// Removes the Finished Good flag one or more Packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) FinishedgoodUnflagPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/finishedgood/unflag", queryParams, body)
}

// GET /packages/v2/active
// Retrieves a list of active packages for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetActivePackages(
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

	return c.send(ctx, "GET", "/packages/v2/active", queryParams, nil)
}

// GET /packages/v2/adjust/reasons
// Retrieves a list of adjustment reasons for packages at a specified Facility.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPackagesAdjustReasons(
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

	return c.send(ctx, "GET", "/packages/v2/adjust/reasons", queryParams, nil)
}

// GET /packages/v2/adjustments
// Retrieves the Package Adjustments for a Facility
// Permissions Required:
// - View Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPackagesAdjustments(
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

	return c.send(ctx, "GET", "/packages/v2/adjustments", queryParams, nil)
}

// GET /packages/v2/intransit
// Retrieves a list of packages in transit for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetInTransitPackages(
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

	return c.send(ctx, "GET", "/packages/v2/intransit", queryParams, nil)
}

// GET /packages/v2/inactive
// Retrieves a list of inactive packages for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetInactivePackages(
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

	return c.send(ctx, "GET", "/packages/v2/inactive", queryParams, nil)
}

// GET /packages/v2/labsamples
// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPackagesLabSamples(
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

	return c.send(ctx, "GET", "/packages/v2/labsamples", queryParams, nil)
}

// GET /packages/v2/onhold
// Retrieves a list of packages on hold for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetOnHoldPackages(
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

	return c.send(ctx, "GET", "/packages/v2/onhold", queryParams, nil)
}

// GET /packages/v2/{id}
// Retrieves a Package by its Id.
// Permissions Required:
// - View Packages
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPackagesById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/packages/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /packages/v2/{label}
// Retrieves a Package by its label.
// Permissions Required:
// - View Packages
// Parameters:
// - label (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPackagesByLabel(
	ctx context.Context,label string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/packages/v2/"+url.QueryEscape(label)+"", queryParams, nil)
}

// GET /packages/v2/types
// Retrieves a list of available Package types.
// Parameters:
func (c *MetrcClient) GetPackagesTypes(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/packages/v2/types", queryParams, nil)
}

// GET /packages/v2/{id}/source/harvests
// Retrieves the source harvests for a Package by its Id.
// Permissions Required:
// - View Package Source Harvests
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPackagesSourceHarvestById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/packages/v2/"+url.QueryEscape(id)+"/source/harvests", queryParams, nil)
}

// GET /packages/v2/transferred
// Retrieves a list of transferred packages for a specific Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetPackagesTransferred(
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

	return c.send(ctx, "GET", "/packages/v2/transferred", queryParams, nil)
}

// PUT /packages/v2/unfinish
// Updates a list of packages as unfinished for a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UnfinishPackagesPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/unfinish", queryParams, body)
}

// PUT /packages/v2/adjust
// Set the final quantity for a Package.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateAdjustPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/adjust", queryParams, body)
}

// PUT /packages/v2/decontaminate
// Updates the Product decontaminate information for a list of packages at a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesDecontaminate(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/decontaminate", queryParams, body)
}

// PUT /packages/v2/donation/flag
// Flags one or more packages for donation at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesDonationFlag(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/donation/flag", queryParams, body)
}

// PUT /packages/v2/donation/unflag
// Removes the donation flag from one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesDonationUnflag(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/donation/unflag", queryParams, body)
}

// PUT /packages/v2/externalid
// Updates the external identifiers for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Package Inventory
// - External Id Enabled
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesExternalid(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/externalid", queryParams, body)
}

// PUT /packages/v2/item
// Updates the associated Item for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesItem(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/item", queryParams, body)
}

// PUT /packages/v2/labtests/required
// Updates the list of required lab test batches for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesLabtestsRequired(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/labtests/required", queryParams, body)
}

// PUT /packages/v2/note
// Updates notes associated with one or more packages for the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// - Manage Package Notes
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesNote(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/note", queryParams, body)
}

// PUT /packages/v2/location
// Updates the Location and Sublocation for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesLocation(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/location", queryParams, body)
}

// PUT /packages/v2/remediate
// Updates a list of Product remediations for packages at a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesRemediate(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/remediate", queryParams, body)
}

// PUT /packages/v2/tradesample/flag
// Flags or unflags one or more packages at the specified Facility as trade samples.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesTradeSampleFlag(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/tradesample/flag", queryParams, body)
}

// PUT /packages/v2/tradesample/unflag
// Removes the trade sample flag from one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesTradeSampleUnflag(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/tradesample/unflag", queryParams, body)
}

// PUT /packages/v2/usebydate
// Updates the use-by date for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdatePackagesUseByDate(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/packages/v2/usebydate", queryParams, body)
}


