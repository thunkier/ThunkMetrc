package client

import (
	"context"
	"net/url"
)


// POST /labtests/v2/record
// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateLabTestsRecord(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/labtests/v2/record", queryParams, body)
}

// GET /labtests/v2/batches
// Retrieves a list of Lab Test batches.
// Parameters:
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetLabTestsBatches(
	ctx context.Context,pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/labtests/v2/batches", queryParams, nil)
}

// GET /labtests/v2/labtestdocument/{id}
// Retrieves a specific Lab Test result document by its Id for a given Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetLabTestsLabTestDocumentById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/labtests/v2/labtestdocument/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /labtests/v2/types
// Returns a list of Lab Test types.
// Parameters:
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetLabTestsTypes(
	ctx context.Context,pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/labtests/v2/types", queryParams, nil)
}

// GET /labtests/v2/results
// Retrieves Lab Test results for a specified Package.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - packageId (optional): Filter by packageId
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetLabTestsResults(
	ctx context.Context,licenseNumber string, packageId string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if packageId != "" {
		queryParams["packageId"] = packageId
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/labtests/v2/results", queryParams, nil)
}

// GET /labtests/v2/states
// Returns a list of all lab testing states.
// Parameters:
func (c *MetrcClient) GetLabTestsStates(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/labtests/v2/states", queryParams, nil)
}

// PUT /labtests/v2/labtestdocument
// Updates one or more documents for previously submitted lab tests.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateLabTestsLabTestDocument(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/labtests/v2/labtestdocument", queryParams, body)
}

// PUT /labtests/v2/results/release
// Releases Lab Test results for one or more packages.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateLabTestsResultsRelease(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/labtests/v2/results/release", queryParams, body)
}


