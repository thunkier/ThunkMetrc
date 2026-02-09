package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateStrains (Strains)
// POST {{baseUrl}}/strains/v2
//   licenseNumber (required): The License Number of the Facility for which to record the list of strains.
func (f *Fetcher) CreateStrains(body []models.StrainsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/strains/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Strains", "CreateStrains", "POST", url, body)
}

// DeleteStrainsById (Strains)
// DELETE {{baseUrl}}/strains/v2/{id}
//   licenseNumber (required): The License Number of the Facility for which Strain to delete.
func (f *Fetcher) DeleteStrainsById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/strains/v2/{id}"
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
	return fetchOne[map[string]any](f, "Strains", "DeleteStrainsById", "DELETE", url, nil)
}

// GetActiveStrains (Strains)
// GET {{baseUrl}}/strains/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active strains.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveStrains(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Strain, error) {
	url := "{{baseUrl}}/strains/v2/active"
	
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
	return fetchList[models.Strain](f, "Strains", "GetActiveStrains", "GET", url, nil)
}

// GetInactiveStrains (Strains)
// GET {{baseUrl}}/strains/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive strains.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveStrains(licenseNumber string, pageNumber string, pageSize string) ([]models.Strain, error) {
	url := "{{baseUrl}}/strains/v2/inactive"
	
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
	return fetchList[models.Strain](f, "Strains", "GetInactiveStrains", "GET", url, nil)
}

// GetStrainsById (Strains)
// GET {{baseUrl}}/strains/v2/{id}
//   licenseNumber (required): If specified, the Strain will be validated against the specified License Number. If not specified, the Strain will be validated against all of the User's current Facilities. Please note that if the Strain is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetStrainsById(id string, licenseNumber string) (models.Strain, error) {
	url := "{{baseUrl}}/strains/v2/{id}"
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
	return fetchOne[models.Strain](f, "Strains", "GetStrainsById", "GET", url, nil)
}

// UpdateStrains (Strains)
// PUT {{baseUrl}}/strains/v2
//   licenseNumber (required): The License Number of the Facility for which to update the list of strains.
func (f *Fetcher) UpdateStrains(body []models.StrainsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/strains/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Strains", "UpdateStrains", "PUT", url, body)
}
