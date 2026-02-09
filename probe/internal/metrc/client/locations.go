package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateLocations (Locations)
// POST {{baseUrl}}/locations/v2
//   licenseNumber (required): The License Number of the Facility for which to record the list of locations.
func (f *Fetcher) CreateLocations(body []models.LocationsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/locations/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Locations", "CreateLocations", "POST", url, body)
}

// DeleteLocationsById (Locations)
// DELETE {{baseUrl}}/locations/v2/{id}
//   licenseNumber (required): The License Number of the Facility for which to delete the list of locations.
func (f *Fetcher) DeleteLocationsById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/locations/v2/{id}"
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
	return fetchOne[map[string]any](f, "Locations", "DeleteLocationsById", "DELETE", url, nil)
}

// GetActiveLocations (Locations)
// GET {{baseUrl}}/locations/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active locations.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveLocations(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.LocationsLocation, error) {
	url := "{{baseUrl}}/locations/v2/active"
	
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
	return fetchList[models.LocationsLocation](f, "Locations", "GetActiveLocations", "GET", url, nil)
}

// GetInactiveLocations (Locations)
// GET {{baseUrl}}/locations/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive locations.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveLocations(licenseNumber string, pageNumber string, pageSize string) ([]models.LocationsLocation, error) {
	url := "{{baseUrl}}/locations/v2/inactive"
	
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
	return fetchList[models.LocationsLocation](f, "Locations", "GetInactiveLocations", "GET", url, nil)
}

// GetLocationsById (Locations)
// GET {{baseUrl}}/locations/v2/{id}
//   licenseNumber (required): If specified, the Location will be validated against the specified License Number. If not specified, the Location will be validated against all of the User's current Facilities. Please note that if the Location is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetLocationsById(id string, licenseNumber string) (models.LocationsLocation, error) {
	url := "{{baseUrl}}/locations/v2/{id}"
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
	return fetchOne[models.LocationsLocation](f, "Locations", "GetLocationsById", "GET", url, nil)
}

// GetLocationsTypes (Locations)
// GET {{baseUrl}}/locations/v2/types
//   licenseNumber (required): The License Number of the Facility for which to return the list of active location types.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetLocationsTypes(licenseNumber string, pageNumber string, pageSize string) ([]models.LocationsType, error) {
	url := "{{baseUrl}}/locations/v2/types"
	
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
	return fetchList[models.LocationsType](f, "Locations", "GetLocationsTypes", "GET", url, nil)
}

// UpdateLocations (Locations)
// PUT {{baseUrl}}/locations/v2
//   licenseNumber (required): The License Number of the Facility for which to update the list of locations.
func (f *Fetcher) UpdateLocations(body []models.LocationsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/locations/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Locations", "UpdateLocations", "PUT", url, body)
}
