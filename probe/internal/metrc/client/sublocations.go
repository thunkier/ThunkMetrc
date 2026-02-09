package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateSublocations (Sublocations)
// POST {{baseUrl}}/sublocations/v2
//   licenseNumber (required): The License Number of the Facility for which to record the list of sublocations.
func (f *Fetcher) CreateSublocations(body []models.SublocationsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sublocations/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Sublocations", "CreateSublocations", "POST", url, body)
}

// DeleteSublocationsById (Sublocations)
// DELETE {{baseUrl}}/sublocations/v2/{id}
//   licenseNumber (required): The License Number of the Facility for which to delete the list of sublocations.
func (f *Fetcher) DeleteSublocationsById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/sublocations/v2/{id}"
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
	return fetchOne[map[string]any](f, "Sublocations", "DeleteSublocationsById", "DELETE", url, nil)
}

// GetActiveSublocations (Sublocations)
// GET {{baseUrl}}/sublocations/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active sublocations.
//   lastModifiedStart (optional): The Last Modified start timestamp
//   lastModifiedEnd (optional): The Last Modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveSublocations(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Sublocation, error) {
	url := "{{baseUrl}}/sublocations/v2/active"
	
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
	return fetchList[models.Sublocation](f, "Sublocations", "GetActiveSublocations", "GET", url, nil)
}

// GetInactiveSublocations (Sublocations)
// GET {{baseUrl}}/sublocations/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive sublocations.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveSublocations(licenseNumber string, pageNumber string, pageSize string) ([]models.Sublocation, error) {
	url := "{{baseUrl}}/sublocations/v2/inactive"
	
	queryParams := make([]string, 0)
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
	return fetchList[models.Sublocation](f, "Sublocations", "GetInactiveSublocations", "GET", url, nil)
}

// GetSublocationsById (Sublocations)
// GET {{baseUrl}}/sublocations/v2/{id}
//   licenseNumber (required): If specified, the Sublocation will be validated against the specified License Number. If not specified, the Sublocation will be validated against all of the User's current Facilities. Please note that if the Sublocation is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetSublocationsById(id string, licenseNumber string) (models.Sublocation, error) {
	url := "{{baseUrl}}/sublocations/v2/{id}"
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
	return fetchOne[models.Sublocation](f, "Sublocations", "GetSublocationsById", "GET", url, nil)
}

// UpdateSublocations (Sublocations)
// PUT {{baseUrl}}/sublocations/v2
//   licenseNumber (required): The License Number of the Facility for which to update the list of sublocations.
func (f *Fetcher) UpdateSublocations(body []models.SublocationsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sublocations/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Sublocations", "UpdateSublocations", "PUT", url, body)
}
