package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateAdjustProcessingJob (ProcessingJob)
// POST {{baseUrl}}/processing/v2/adjust
//   licenseNumber (required): The License Number of the Facility to adjust the Processing Job.
func (f *Fetcher) CreateAdjustProcessingJob(body []models.AdjustProcessingJobRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/processing/v2/adjust"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[models.WriteResponse](f, "ProcessingJob", "CreateAdjustProcessingJob", "POST", url, body)
}

// CreateJobTypes (ProcessingJob)
// POST {{baseUrl}}/processing/v2/jobtypes
//   licenseNumber (required): The License Number of the Facility of the Job Type.
func (f *Fetcher) CreateJobTypes(body []models.ProcessingJobJobTypesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/processing/v2/jobtypes"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[models.WriteResponse](f, "ProcessingJob", "CreateJobTypes", "POST", url, body)
}

// CreateProcessingJobPackages (ProcessingJob)
// POST {{baseUrl}}/processing/v2/createpackages
//   licenseNumber (required): The License Number of the Facility for which to create the packages.
func (f *Fetcher) CreateProcessingJobPackages(body []models.ProcessingJobPackagesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/processing/v2/createpackages"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[models.WriteResponse](f, "ProcessingJob", "CreateProcessingJobPackages", "POST", url, body)
}

// DeleteJobTypeById (ProcessingJob)
// DELETE {{baseUrl}}/processing/v2/jobtypes/{id}
//   licenseNumber (required): The License Number of the Facility of the Processing Job Type.
func (f *Fetcher) DeleteJobTypeById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/processing/v2/jobtypes/{id}"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[map[string]any](f, "ProcessingJob", "DeleteJobTypeById", "DELETE", url, nil)
}

// DeleteProcessingJobById (ProcessingJob)
// DELETE {{baseUrl}}/processing/v2/{id}
//   licenseNumber (required): The License Number of the Facility of the Processing Job.
func (f *Fetcher) DeleteProcessingJobById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/processing/v2/{id}"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[map[string]any](f, "ProcessingJob", "DeleteProcessingJobById", "DELETE", url, nil)
}

// FinishProcessingJob (ProcessingJob)
// PUT {{baseUrl}}/processing/v2/finish
//   licenseNumber (required): The License Number of the Facility to finish the Processing Job.
func (f *Fetcher) FinishProcessingJob(body []models.FinishProcessingJobRequest, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/processing/v2/finish"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[map[string]any](f, "ProcessingJob", "FinishProcessingJob", "PUT", url, body)
}

// GetActiveJobTypes (ProcessingJob)
// GET {{baseUrl}}/processing/v2/jobtypes/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active processing job types.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveJobTypes(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.ActiveJobType, error) {
	url := "{{baseUrl}}/processing/v2/jobtypes/active"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}
	if lastModifiedStart != "" {
		queryParams = append(queryParams, fmt.Sprintf("lastModifiedStart=%s", lastModifiedStart))
	}
	if lastModifiedEnd != "" {
		queryParams = append(queryParams, fmt.Sprintf("lastModifiedEnd=%s", lastModifiedEnd))
	}
	if pageNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("pageNumber=%s", pageNumber))
	}
	if pageSize != "" {
		queryParams = append(queryParams, fmt.Sprintf("pageSize=%s", pageSize))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchList[models.ActiveJobType](f, "ProcessingJob", "GetActiveJobTypes", "GET", url, nil)
}

// GetActiveProcessingJob (ProcessingJob)
// GET {{baseUrl}}/processing/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active processing jobs.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveProcessingJob(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.ProcessingJob, error) {
	url := "{{baseUrl}}/processing/v2/active"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}
	if lastModifiedStart != "" {
		queryParams = append(queryParams, fmt.Sprintf("lastModifiedStart=%s", lastModifiedStart))
	}
	if lastModifiedEnd != "" {
		queryParams = append(queryParams, fmt.Sprintf("lastModifiedEnd=%s", lastModifiedEnd))
	}
	if pageNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("pageNumber=%s", pageNumber))
	}
	if pageSize != "" {
		queryParams = append(queryParams, fmt.Sprintf("pageSize=%s", pageSize))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchList[models.ProcessingJob](f, "ProcessingJob", "GetActiveProcessingJob", "GET", url, nil)
}

// GetInactiveJobTypes (ProcessingJob)
// GET {{baseUrl}}/processing/v2/jobtypes/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive processing job types.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveJobTypes(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.InactiveJobType, error) {
	url := "{{baseUrl}}/processing/v2/jobtypes/inactive"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}
	if lastModifiedStart != "" {
		queryParams = append(queryParams, fmt.Sprintf("lastModifiedStart=%s", lastModifiedStart))
	}
	if lastModifiedEnd != "" {
		queryParams = append(queryParams, fmt.Sprintf("lastModifiedEnd=%s", lastModifiedEnd))
	}
	if pageNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("pageNumber=%s", pageNumber))
	}
	if pageSize != "" {
		queryParams = append(queryParams, fmt.Sprintf("pageSize=%s", pageSize))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchList[models.InactiveJobType](f, "ProcessingJob", "GetInactiveJobTypes", "GET", url, nil)
}

// GetInactiveProcessingJob (ProcessingJob)
// GET {{baseUrl}}/processing/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive processing jobs.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveProcessingJob(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.ProcessingJob, error) {
	url := "{{baseUrl}}/processing/v2/inactive"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}
	if lastModifiedStart != "" {
		queryParams = append(queryParams, fmt.Sprintf("lastModifiedStart=%s", lastModifiedStart))
	}
	if lastModifiedEnd != "" {
		queryParams = append(queryParams, fmt.Sprintf("lastModifiedEnd=%s", lastModifiedEnd))
	}
	if pageNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("pageNumber=%s", pageNumber))
	}
	if pageSize != "" {
		queryParams = append(queryParams, fmt.Sprintf("pageSize=%s", pageSize))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchList[models.ProcessingJob](f, "ProcessingJob", "GetInactiveProcessingJob", "GET", url, nil)
}

// GetJobTypesAttributes (ProcessingJob)
// GET {{baseUrl}}/processing/v2/jobtypes/attributes
//   licenseNumber (required): The License Number of the Facility for which to return the list of active processing job type attributes.
func (f *Fetcher) GetJobTypesAttributes(licenseNumber string) ([]models.JobTypesAttribute, error) {
	url := "{{baseUrl}}/processing/v2/jobtypes/attributes"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchList[models.JobTypesAttribute](f, "ProcessingJob", "GetJobTypesAttributes", "GET", url, nil)
}

// GetJobTypesCategories (ProcessingJob)
// GET {{baseUrl}}/processing/v2/jobtypes/categories
//   licenseNumber (required): The License Number of the Facility for which to return the list of processing job type categories.
func (f *Fetcher) GetJobTypesCategories(licenseNumber string) ([]models.JobTypesCategory, error) {
	url := "{{baseUrl}}/processing/v2/jobtypes/categories"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchList[models.JobTypesCategory](f, "ProcessingJob", "GetJobTypesCategories", "GET", url, nil)
}

// GetProcessingJobById (ProcessingJob)
// GET {{baseUrl}}/processing/v2/{id}
//   licenseNumber (required): If specified, the processing job will be validated against the specified License Number. If not specified, the processing job will be validated against all of the User's current Facilities. Please note that if the processing job is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetProcessingJobById(id string, licenseNumber string) (models.ProcessingJob, error) {
	url := "{{baseUrl}}/processing/v2/{id}"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[models.ProcessingJob](f, "ProcessingJob", "GetProcessingJobById", "GET", url, nil)
}

// StartProcessingJob (ProcessingJob)
// POST {{baseUrl}}/processing/v2/start
//   licenseNumber (required): The License Number of the Facility to start the Processing Job.
func (f *Fetcher) StartProcessingJob(body []models.StartProcessingJobRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/processing/v2/start"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[models.WriteResponse](f, "ProcessingJob", "StartProcessingJob", "POST", url, body)
}

// UnfinishProcessingJob (ProcessingJob)
// PUT {{baseUrl}}/processing/v2/unfinish
//   licenseNumber (required): The License Number of the Facility to unfinish the Processing Job.
func (f *Fetcher) UnfinishProcessingJob(body []models.UnfinishProcessingJobRequest, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/processing/v2/unfinish"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[map[string]any](f, "ProcessingJob", "UnfinishProcessingJob", "PUT", url, body)
}

// UpdateJobTypes (ProcessingJob)
// PUT {{baseUrl}}/processing/v2/jobtypes
//   licenseNumber (required): The License Number of the Facility for which to update the processing job types.
func (f *Fetcher) UpdateJobTypes(body []models.ProcessingJobJobTypesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/processing/v2/jobtypes"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[models.WriteResponse](f, "ProcessingJob", "UpdateJobTypes", "PUT", url, body)
}
