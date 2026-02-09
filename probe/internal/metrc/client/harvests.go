package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateHarvestsPackages (Harvests)
// POST {{baseUrl}}/harvests/v2/packages
//   licenseNumber (required): The License Number of the Facility for which to return the list of harvests packages.
func (f *Fetcher) CreateHarvestsPackages(body []models.HarvestsPackagesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/harvests/v2/packages"
	
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
	return fetchOne[models.WriteResponse](f, "Harvests", "CreateHarvestsPackages", "POST", url, body)
}

// CreateHarvestsWaste (Harvests)
// POST {{baseUrl}}/harvests/v2/waste
//   licenseNumber (required): The License Number of the Facility for which to return the list of harvests waste.
func (f *Fetcher) CreateHarvestsWaste(body []models.HarvestsWasteRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/harvests/v2/waste"
	
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
	return fetchOne[models.WriteResponse](f, "Harvests", "CreateHarvestsWaste", "POST", url, body)
}

// CreatePackagesTesting (Harvests)
// POST {{baseUrl}}/harvests/v2/packages/testing
//   licenseNumber (required): The License Number of the Facility for which to return the list of harvests packages.
func (f *Fetcher) CreatePackagesTesting(body []models.PackagesTestingRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/harvests/v2/packages/testing"
	
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
	return fetchOne[models.WriteResponse](f, "Harvests", "CreatePackagesTesting", "POST", url, body)
}

// DeleteWasteById (Harvests)
// DELETE {{baseUrl}}/harvests/v2/waste/{id}
//   licenseNumber (required): The License Number of the Facility for which to return the list of harvests waste.
func (f *Fetcher) DeleteWasteById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/harvests/v2/waste/{id}"
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
	return fetchOne[map[string]any](f, "Harvests", "DeleteWasteById", "DELETE", url, nil)
}

// FinishHarvests (Harvests)
// PUT {{baseUrl}}/harvests/v2/finish
//   licenseNumber (required): The License Number of the Facility for which to return the list of finished harvests.
func (f *Fetcher) FinishHarvests(body []models.FinishHarvestsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/harvests/v2/finish"
	
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
	return fetchOne[models.WriteResponse](f, "Harvests", "FinishHarvests", "PUT", url, body)
}

// GetActiveHarvests (Harvests)
// GET {{baseUrl}}/harvests/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active harvests.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveHarvests(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Harvest, error) {
	url := "{{baseUrl}}/harvests/v2/active"
	
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
	return fetchList[models.Harvest](f, "Harvests", "GetActiveHarvests", "GET", url, nil)
}

// GetHarvestsById (Harvests)
// GET {{baseUrl}}/harvests/v2/{id}
//   licenseNumber (required): If specified, the Harvest will be validated against the specified License Number. If not specified, the Harvest will be validated against all of the User's current Facilities. Please note that if the Harvest is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetHarvestsById(id string, licenseNumber string) (models.Harvest, error) {
	url := "{{baseUrl}}/harvests/v2/{id}"
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
	return fetchOne[models.Harvest](f, "Harvests", "GetHarvestsById", "GET", url, nil)
}

// GetHarvestsWaste (Harvests)
// GET {{baseUrl}}/harvests/v2/waste
//   licenseNumber (required): The License Number of the Facility for which to return the list of waste harvests.
//   harvestId (required): The harvestId is the unique identifier for each harvest.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetHarvestsWaste(licenseNumber string, harvestId string, pageNumber string, pageSize string) ([]models.HarvestsWaste, error) {
	url := "{{baseUrl}}/harvests/v2/waste"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}
	if harvestId != "" {
		queryParams = append(queryParams, fmt.Sprintf("harvestId=%s", harvestId))
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
	return fetchList[models.HarvestsWaste](f, "Harvests", "GetHarvestsWaste", "GET", url, nil)
}

// GetInactiveHarvests (Harvests)
// GET {{baseUrl}}/harvests/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive harvests.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveHarvests(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Harvest, error) {
	url := "{{baseUrl}}/harvests/v2/inactive"
	
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
	return fetchList[models.Harvest](f, "Harvests", "GetInactiveHarvests", "GET", url, nil)
}

// GetOnHoldHarvests (Harvests)
// GET {{baseUrl}}/harvests/v2/onhold
//   licenseNumber (required): The License Number of the Facility for which to return the list of harvests on hold.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetOnHoldHarvests(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Harvest, error) {
	url := "{{baseUrl}}/harvests/v2/onhold"
	
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
	return fetchList[models.Harvest](f, "Harvests", "GetOnHoldHarvests", "GET", url, nil)
}

// GetWasteTypes (Harvests)
// GET {{baseUrl}}/harvests/v2/waste/types
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetWasteTypes(pageNumber string, pageSize string) ([]models.WasteType, error) {
	url := "{{baseUrl}}/harvests/v2/waste/types"
	
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
	return fetchList[models.WasteType](f, "Harvests", "GetWasteTypes", "GET", url, nil)
}

// UnfinishHarvests (Harvests)
// PUT {{baseUrl}}/harvests/v2/unfinish
//   licenseNumber (required): The License Number of the Facility for which to return the list of unfinished harvests.
func (f *Fetcher) UnfinishHarvests(body []models.UnfinishHarvestsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/harvests/v2/unfinish"
	
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
	return fetchOne[models.WriteResponse](f, "Harvests", "UnfinishHarvests", "PUT", url, body)
}

// UpdateHarvestsLocation (Harvests)
// PUT {{baseUrl}}/harvests/v2/location
//   licenseNumber (required): The License Number of the Facility for which to return the list of harvests locations.
func (f *Fetcher) UpdateHarvestsLocation(body []models.HarvestsLocationRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/harvests/v2/location"
	
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
	return fetchOne[models.WriteResponse](f, "Harvests", "UpdateHarvestsLocation", "PUT", url, body)
}

// UpdateRename (Harvests)
// PUT {{baseUrl}}/harvests/v2/rename
//   licenseNumber (required): The License Number of the Facility for which to return the list of renamed harvests.
func (f *Fetcher) UpdateRename(body []models.RenameRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/harvests/v2/rename"
	
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
	return fetchOne[models.WriteResponse](f, "Harvests", "UpdateRename", "PUT", url, body)
}

// UpdateRestoreHarvestedPlants (Harvests)
// PUT {{baseUrl}}/harvests/v2/restore/harvestedplants
//   licenseNumber (required): The License Number of the Facility for which to return the list of harvests restored.
func (f *Fetcher) UpdateRestoreHarvestedPlants(body []models.RestoreHarvestedPlantsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/harvests/v2/restore/harvestedplants"
	
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
	return fetchOne[models.WriteResponse](f, "Harvests", "UpdateRestoreHarvestedPlants", "PUT", url, body)
}
