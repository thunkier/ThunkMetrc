package client

import (
	"context"
	"net/url"
)


// POST /processing/v2/adjust
// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateAdjustProcessingJob(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/processing/v2/adjust", queryParams, body)
}

// POST /processing/v2/jobtypes
// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateProcessingJobJobTypes(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/processing/v2/jobtypes", queryParams, body)
}

// POST /processing/v2/createpackages
// Creates packages from processing jobs at a Facility, including optional location and note assignments.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateProcessingJobPackages(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/processing/v2/createpackages", queryParams, body)
}

// DELETE /processing/v2/jobtypes/{id}
// Archives a Processing Job Type at a Facility, making it inactive for future use.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteProcessingJobJobTypeById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/processing/v2/jobtypes/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE /processing/v2/{id}
// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteProcessingJobById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/processing/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// PUT /processing/v2/finish
// Completes processing jobs at a Facility by recording final notes and waste measurements.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) FinishProcessingJobProcessingJob(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/processing/v2/finish", queryParams, body)
}

// GET /processing/v2/jobtypes/active
// Retrieves a list of all active processing job types defined within a Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetProcessingJobActiveJobTypes(
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

	return c.send(ctx, "GET", "/processing/v2/jobtypes/active", queryParams, nil)
}

// GET /processing/v2/active
// Retrieves active processing jobs at a specified Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetActiveProcessingJob(
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

	return c.send(ctx, "GET", "/processing/v2/active", queryParams, nil)
}

// GET /processing/v2/jobtypes/inactive
// Retrieves a list of all inactive processing job types defined within a Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetProcessingJobInactiveJobTypes(
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

	return c.send(ctx, "GET", "/processing/v2/jobtypes/inactive", queryParams, nil)
}

// GET /processing/v2/inactive
// Retrieves inactive processing jobs at a specified Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetInactiveProcessingJob(
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

	return c.send(ctx, "GET", "/processing/v2/inactive", queryParams, nil)
}

// GET /processing/v2/jobtypes/attributes
// Retrieves all processing job attributes available for a Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetProcessingJobJobTypesAttributes(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/processing/v2/jobtypes/attributes", queryParams, nil)
}

// GET /processing/v2/jobtypes/categories
// Retrieves all processing job categories available for a specified Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetProcessingJobJobTypesCategories(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/processing/v2/jobtypes/categories", queryParams, nil)
}

// GET /processing/v2/{id}
// Retrieves a ProcessingJob by Id.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetProcessingJobById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/processing/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// POST /processing/v2/start
// Initiates new processing jobs at a Facility, including job details and associated packages.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) StartProcessingJobProcessingJob(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/processing/v2/start", queryParams, body)
}

// PUT /processing/v2/unfinish
// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UnfinishProcessingJobProcessingJob(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/processing/v2/unfinish", queryParams, body)
}

// PUT /processing/v2/jobtypes
// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateProcessingJobJobTypes(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/processing/v2/jobtypes", queryParams, body)
}


