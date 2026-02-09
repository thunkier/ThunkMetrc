package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateRecord (LabTests)
// POST {{baseUrl}}/labtests/v2/record
//   licenseNumber (required): The License Number of the Facility recording the Lab Test.
func (f *Fetcher) CreateRecord(body []models.RecordRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/labtests/v2/record"
	
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
	return fetchOne[models.WriteResponse](f, "LabTests", "CreateRecord", "POST", url, body)
}

// GetBatches (LabTests)
// GET {{baseUrl}}/labtests/v2/batches
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetBatches(pageNumber string, pageSize string) ([]models.Batch, error) {
	url := "{{baseUrl}}/labtests/v2/batches"
	
	queryParams := make([]string, 0)
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
	return fetchList[models.Batch](f, "LabTests", "GetBatches", "GET", url, nil)
}

// GetLabTestDocumentById (LabTests)
// GET {{baseUrl}}/labtests/v2/labtestdocument/{id}
//   licenseNumber (required): LabTestDocumentFileId
func (f *Fetcher) GetLabTestDocumentById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/labtests/v2/labtestdocument/{id}"
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
	return fetchOne[map[string]any](f, "LabTests", "GetLabTestDocumentById", "GET", url, nil)
}

// GetLabTestsTypes (LabTests)
// GET {{baseUrl}}/labtests/v2/types
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetLabTestsTypes(pageNumber string, pageSize string) ([]models.LabTestsType, error) {
	url := "{{baseUrl}}/labtests/v2/types"
	
	queryParams := make([]string, 0)
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
	return fetchList[models.LabTestsType](f, "LabTests", "GetLabTestsTypes", "GET", url, nil)
}

// GetResults (LabTests)
// GET {{baseUrl}}/labtests/v2/results
//   packageId (required): 
//   licenseNumber (required): 
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetResults(packageId string, licenseNumber string, pageNumber string, pageSize string) ([]models.Result, error) {
	url := "{{baseUrl}}/labtests/v2/results"
	
	queryParams := make([]string, 0)
	if packageId != "" {
		queryParams = append(queryParams, fmt.Sprintf("packageId=%s", packageId))
	}
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
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
	return fetchList[models.Result](f, "LabTests", "GetResults", "GET", url, nil)
}

// GetStates (LabTests)
// GET {{baseUrl}}/labtests/v2/states
func (f *Fetcher) GetStates() (map[string]any, error) {
	url := "{{baseUrl}}/labtests/v2/states"
	return fetchOne[map[string]any](f, "LabTests", "GetStates", "GET", url, nil)
}

// UpdateLabTestDocument (LabTests)
// PUT {{baseUrl}}/labtests/v2/labtestdocument
//   licenseNumber (required): The License Number of the Facility updating the Lab Test document.
func (f *Fetcher) UpdateLabTestDocument(body []models.LabTestDocumentRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/labtests/v2/labtestdocument"
	
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
	return fetchOne[models.WriteResponse](f, "LabTests", "UpdateLabTestDocument", "PUT", url, body)
}

// UpdateResultsRelease (LabTests)
// PUT {{baseUrl}}/labtests/v2/results/release
//   licenseNumber (required): The License Number of the Facility releasing the Lab Test results.
func (f *Fetcher) UpdateResultsRelease(body []models.ResultsReleaseRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/labtests/v2/results/release"
	
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
	return fetchOne[models.WriteResponse](f, "LabTests", "UpdateResultsRelease", "PUT", url, body)
}
