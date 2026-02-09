package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateAdditivesTemplates (AdditivesTemplates)
// POST {{baseUrl}}/additivestemplates/v2
//   licenseNumber (required): The License Number of the Facility for which to record the list of additives.
func (f *Fetcher) CreateAdditivesTemplates(body []models.AdditivesTemplatesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/additivestemplates/v2"
	
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
	return fetchOne[models.WriteResponse](f, "AdditivesTemplates", "CreateAdditivesTemplates", "POST", url, body)
}

// GetActiveAdditivesTemplates (AdditivesTemplates)
// GET {{baseUrl}}/additivestemplates/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active additives.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveAdditivesTemplates(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.AdditivesTemplate, error) {
	url := "{{baseUrl}}/additivestemplates/v2/active"
	
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
	return fetchList[models.AdditivesTemplate](f, "AdditivesTemplates", "GetActiveAdditivesTemplates", "GET", url, nil)
}

// GetAdditivesTemplatesById (AdditivesTemplates)
// GET {{baseUrl}}/additivestemplates/v2/{id}
//   licenseNumber (required): If specified, the Additive will be validated against the specified License Number. If not specified, the Additive will be validated against all of the User's current Facilities. Please note that if the Additive is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetAdditivesTemplatesById(id string, licenseNumber string) (models.AdditivesTemplate, error) {
	url := "{{baseUrl}}/additivestemplates/v2/{id}"
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
	return fetchOne[models.AdditivesTemplate](f, "AdditivesTemplates", "GetAdditivesTemplatesById", "GET", url, nil)
}

// GetInactiveAdditivesTemplates (AdditivesTemplates)
// GET {{baseUrl}}/additivestemplates/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive additives.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveAdditivesTemplates(licenseNumber string, pageNumber string, pageSize string) ([]models.AdditivesTemplate, error) {
	url := "{{baseUrl}}/additivestemplates/v2/inactive"
	
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
	return fetchList[models.AdditivesTemplate](f, "AdditivesTemplates", "GetInactiveAdditivesTemplates", "GET", url, nil)
}

// UpdateAdditivesTemplates (AdditivesTemplates)
// PUT {{baseUrl}}/additivestemplates/v2
//   licenseNumber (required): The License Number of the Facility for which to update the list of additives.
func (f *Fetcher) UpdateAdditivesTemplates(body []models.AdditivesTemplatesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/additivestemplates/v2"
	
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
	return fetchOne[models.WriteResponse](f, "AdditivesTemplates", "UpdateAdditivesTemplates", "PUT", url, body)
}
