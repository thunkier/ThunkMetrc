package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateDrivers (Transporters)
// POST {{baseUrl}}/transporters/v2/drivers
//   licenseNumber (required): The License Number of the Facility for which to create a Driver.
func (f *Fetcher) CreateDrivers(body []models.TransportersDriversRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transporters/v2/drivers"
	
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
	return fetchOne[models.WriteResponse](f, "Transporters", "CreateDrivers", "POST", url, body)
}

// CreateVehicles (Transporters)
// POST {{baseUrl}}/transporters/v2/vehicles
//   licenseNumber (required): The License Number of the Facility for which to create a Vehicle.
func (f *Fetcher) CreateVehicles(body []models.TransportersVehiclesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transporters/v2/vehicles"
	
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
	return fetchOne[models.WriteResponse](f, "Transporters", "CreateVehicles", "POST", url, body)
}

// DeleteDriverById (Transporters)
// DELETE {{baseUrl}}/transporters/v2/drivers/{id}
//   licenseNumber (required): The License Number of the Facility for which to archive a Driver.
func (f *Fetcher) DeleteDriverById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/transporters/v2/drivers/{id}"
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
	return fetchOne[map[string]any](f, "Transporters", "DeleteDriverById", "DELETE", url, nil)
}

// DeleteVehicleById (Transporters)
// DELETE {{baseUrl}}/transporters/v2/vehicles/{id}
//   licenseNumber (required): The License Number for which to archive a Vehicle
func (f *Fetcher) DeleteVehicleById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/transporters/v2/vehicles/{id}"
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
	return fetchOne[map[string]any](f, "Transporters", "DeleteVehicleById", "DELETE", url, nil)
}

// GetDriverById (Transporters)
// GET {{baseUrl}}/transporters/v2/drivers/{id}
//   licenseNumber (required): The License Number of the Facility for which to return a Driver.
func (f *Fetcher) GetDriverById(id string, licenseNumber string) (models.TransportersDriver, error) {
	url := "{{baseUrl}}/transporters/v2/drivers/{id}"
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
	return fetchOne[models.TransportersDriver](f, "Transporters", "GetDriverById", "GET", url, nil)
}

// GetDrivers (Transporters)
// GET {{baseUrl}}/transporters/v2/drivers
//   licenseNumber (required): The License Number of the Facility for which to return the list of drivers.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetDrivers(licenseNumber string, pageNumber string, pageSize string) ([]models.TransportersDriver, error) {
	url := "{{baseUrl}}/transporters/v2/drivers"
	
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
	return fetchList[models.TransportersDriver](f, "Transporters", "GetDrivers", "GET", url, nil)
}

// GetVehicleById (Transporters)
// GET {{baseUrl}}/transporters/v2/vehicles/{id}
//   licenseNumber (required): The License Number of the Facility for which to return a Vehicle.
func (f *Fetcher) GetVehicleById(id string, licenseNumber string) (models.TransportersVehicle, error) {
	url := "{{baseUrl}}/transporters/v2/vehicles/{id}"
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
	return fetchOne[models.TransportersVehicle](f, "Transporters", "GetVehicleById", "GET", url, nil)
}

// GetVehicles (Transporters)
// GET {{baseUrl}}/transporters/v2/vehicles
//   licenseNumber (required): The License Number of the Facility for which to return the list of vehicles.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetVehicles(licenseNumber string, pageNumber string, pageSize string) ([]models.TransportersVehicle, error) {
	url := "{{baseUrl}}/transporters/v2/vehicles"
	
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
	return fetchList[models.TransportersVehicle](f, "Transporters", "GetVehicles", "GET", url, nil)
}

// UpdateDrivers (Transporters)
// PUT {{baseUrl}}/transporters/v2/drivers
//   licenseNumber (required): The License Number of the Facility for which to update the list of drivers.
func (f *Fetcher) UpdateDrivers(body []models.TransportersDriversRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transporters/v2/drivers"
	
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
	return fetchOne[models.WriteResponse](f, "Transporters", "UpdateDrivers", "PUT", url, body)
}

// UpdateVehicles (Transporters)
// PUT {{baseUrl}}/transporters/v2/vehicles
//   licenseNumber (required): The License Number of the Facility for which to update the list of vehicles.
func (f *Fetcher) UpdateVehicles(body []models.TransportersVehiclesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transporters/v2/vehicles"
	
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
	return fetchOne[models.WriteResponse](f, "Transporters", "UpdateVehicles", "PUT", url, body)
}
