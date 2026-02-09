package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateAdditivesByLocation (Plants)
// POST {{baseUrl}}/plants/v2/additives/bylocation
//   licenseNumber (required): The License Number of the Facility for which to record Plant additives.
func (f *Fetcher) CreateAdditivesByLocation(body []models.AdditivesByLocationRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/additives/bylocation"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "CreateAdditivesByLocation", "POST", url, body)
}

// CreateAdditivesByLocationUsingTemplate (Plants)
// POST {{baseUrl}}/plants/v2/additives/bylocation/usingtemplate
//   licenseNumber (required): The License Number of the Facility for which to record Plant additives.
func (f *Fetcher) CreateAdditivesByLocationUsingTemplate(body []models.AdditivesByLocationUsingTemplateRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/additives/bylocation/usingtemplate"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "CreateAdditivesByLocationUsingTemplate", "POST", url, body)
}

// CreateManicure (Plants)
// POST {{baseUrl}}/plants/v2/manicure
//   licenseNumber (required): The License Number of the Facility for which to record the list of plants manicured.
func (f *Fetcher) CreateManicure(body []models.ManicureRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/manicure"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "CreateManicure", "POST", url, body)
}

// CreatePlantBatchPackages (Plants)
// POST {{baseUrl}}/plants/v2/plantbatch/packages
//   licenseNumber (required): The License Number of the Facility for which to record the list of Plant Batch packages.
func (f *Fetcher) CreatePlantBatchPackages(body []models.PlantBatchPackagesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/plantbatch/packages"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "CreatePlantBatchPackages", "POST", url, body)
}

// CreatePlantsAdditives (Plants)
// POST {{baseUrl}}/plants/v2/additives
//   licenseNumber (required): The License Number of the Facility for which to record Plant additives.
func (f *Fetcher) CreatePlantsAdditives(body []models.AdditivesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/additives"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "CreatePlantsAdditives", "POST", url, body)
}

// CreatePlantsAdditivesUsingTemplate (Plants)
// POST {{baseUrl}}/plants/v2/additives/usingtemplate
//   licenseNumber (required): The License Number of the Facility for which to record Plant additives.
func (f *Fetcher) CreatePlantsAdditivesUsingTemplate(body []models.AdditivesUsingTemplateRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/additives/usingtemplate"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "CreatePlantsAdditivesUsingTemplate", "POST", url, body)
}

// CreatePlantsPlantings (Plants)
// POST {{baseUrl}}/plants/v2/plantings
//   licenseNumber (required): The License Number of the Facility for which to record the list of plant batches.
func (f *Fetcher) CreatePlantsPlantings(body []models.PlantsPlantingsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/plantings"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "CreatePlantsPlantings", "POST", url, body)
}

// CreatePlantsWaste (Plants)
// POST {{baseUrl}}/plants/v2/waste
//   licenseNumber (required): The License Number of the Facility for which to record waste.
func (f *Fetcher) CreatePlantsWaste(body []models.PlantsWasteRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/waste"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "CreatePlantsWaste", "POST", url, body)
}

// DeletePlants (Plants)
// DELETE {{baseUrl}}/plants/v2
//   licenseNumber (required): The License Number of the Facility that intends to destroy the plants.
func (f *Fetcher) DeletePlants(licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "DeletePlants", "DELETE", url, nil)
}

// GetAdditives (Plants)
// GET {{baseUrl}}/plants/v2/additives
//   licenseNumber (required): The License Number of the Facility for which to retrieve plant additives.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetAdditives(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Additive, error) {
	url := "{{baseUrl}}/plants/v2/additives"
	
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
	return fetchList[models.Additive](f, "Plants", "GetAdditives", "GET", url, nil)
}

// GetAdditivesTypes (Plants)
// GET {{baseUrl}}/plants/v2/additives/types
func (f *Fetcher) GetAdditivesTypes() (map[string]any, error) {
	url := "{{baseUrl}}/plants/v2/additives/types"
	return fetchOne[map[string]any](f, "Plants", "GetAdditivesTypes", "GET", url, nil)
}

// GetFloweringPlants (Plants)
// GET {{baseUrl}}/plants/v2/flowering
//   licenseNumber (required): The License Number of the Facility for which to return the list of flowering plants.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetFloweringPlants(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Plant, error) {
	url := "{{baseUrl}}/plants/v2/flowering"
	
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
	return fetchList[models.Plant](f, "Plants", "GetFloweringPlants", "GET", url, nil)
}

// GetGrowthPhases (Plants)
// GET {{baseUrl}}/plants/v2/growthphases
//   licenseNumber (required): The License Number of the Facility for which to return the list of available Plant Growth Phases.
func (f *Fetcher) GetGrowthPhases(licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/plants/v2/growthphases"
	
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
	return fetchOne[map[string]any](f, "Plants", "GetGrowthPhases", "GET", url, nil)
}

// GetInactivePlants (Plants)
// GET {{baseUrl}}/plants/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive plants.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactivePlants(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Plant, error) {
	url := "{{baseUrl}}/plants/v2/inactive"
	
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
	return fetchList[models.Plant](f, "Plants", "GetInactivePlants", "GET", url, nil)
}

// GetMotherInactivePlants (Plants)
// GET {{baseUrl}}/plants/v2/mother/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive mother plants.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetMotherInactivePlants(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Mother, error) {
	url := "{{baseUrl}}/plants/v2/mother/inactive"
	
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
	return fetchList[models.Mother](f, "Plants", "GetMotherInactivePlants", "GET", url, nil)
}

// GetMotherOnHoldPlants (Plants)
// GET {{baseUrl}}/plants/v2/mother/onhold
//   licenseNumber (required): The License Number of the Facility for which to return the list of mother plants on hold.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetMotherOnHoldPlants(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Mother, error) {
	url := "{{baseUrl}}/plants/v2/mother/onhold"
	
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
	return fetchList[models.Mother](f, "Plants", "GetMotherOnHoldPlants", "GET", url, nil)
}

// GetMotherPlants (Plants)
// GET {{baseUrl}}/plants/v2/mother
//   licenseNumber (required): The License Number of the Facility for which to return the list of mother plants.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetMotherPlants(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Mother, error) {
	url := "{{baseUrl}}/plants/v2/mother"
	
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
	return fetchList[models.Mother](f, "Plants", "GetMotherPlants", "GET", url, nil)
}

// GetOnHoldPlants (Plants)
// GET {{baseUrl}}/plants/v2/onhold
//   licenseNumber (required): The License Number of the Facility for which to return the list of plants on hold.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetOnHoldPlants(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Plant, error) {
	url := "{{baseUrl}}/plants/v2/onhold"
	
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
	return fetchList[models.Plant](f, "Plants", "GetOnHoldPlants", "GET", url, nil)
}

// GetPlantsById (Plants)
// GET {{baseUrl}}/plants/v2/{id}
//   licenseNumber (required): If specified, the Plant will be validated against the specified License Number. If not specified, the Plant will be validated against all of the User's current Facilities. Please note that if the Plant is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetPlantsById(id string, licenseNumber string) (models.Plant, error) {
	url := "{{baseUrl}}/plants/v2/{id}"
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
	return fetchOne[models.Plant](f, "Plants", "GetPlantsById", "GET", url, nil)
}

// GetPlantsByLabel (Plants)
// GET {{baseUrl}}/plants/v2/{label}
//   licenseNumber (required): If specified, the Plant will be validated against the specified License Number. If not specified, the Plant will be validated against all of the User's current Facilities. Please note that if the Plant is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetPlantsByLabel(label string, licenseNumber string) ([]models.Plant, error) {
	url := "{{baseUrl}}/plants/v2/{label}"
	url = strings.ReplaceAll(url, "{"+"label"+"}", label)
	
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
	return fetchList[models.Plant](f, "Plants", "GetPlantsByLabel", "GET", url, nil)
}

// GetPlantsWaste (Plants)
// GET {{baseUrl}}/plants/v2/waste
//   licenseNumber (required): The License Number of the Facility for which to return the list of waste plants.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetPlantsWaste(licenseNumber string, pageNumber string, pageSize string) ([]models.PlantsWaste, error) {
	url := "{{baseUrl}}/plants/v2/waste"
	
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
	return fetchList[models.PlantsWaste](f, "Plants", "GetPlantsWaste", "GET", url, nil)
}

// GetPlantsWasteMethods (Plants)
// GET {{baseUrl}}/plants/v2/waste/methods/all
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetPlantsWasteMethods(pageNumber string, pageSize string) ([]models.WasteMethod, error) {
	url := "{{baseUrl}}/plants/v2/waste/methods/all"
	
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
	return fetchList[models.WasteMethod](f, "Plants", "GetPlantsWasteMethods", "GET", url, nil)
}

// GetPlantsWasteReasons (Plants)
// GET {{baseUrl}}/plants/v2/waste/reasons
//   licenseNumber (required): The License Number of the Facility for which to return the list of waste reasons.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetPlantsWasteReasons(licenseNumber string, pageNumber string, pageSize string) ([]models.WasteReason, error) {
	url := "{{baseUrl}}/plants/v2/waste/reasons"
	
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
	return fetchList[models.WasteReason](f, "Plants", "GetPlantsWasteReasons", "GET", url, nil)
}

// GetVegetativePlants (Plants)
// GET {{baseUrl}}/plants/v2/vegetative
//   licenseNumber (required): The License Number of the Facility for which to return the list of vegetating plants.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetVegetativePlants(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Plant, error) {
	url := "{{baseUrl}}/plants/v2/vegetative"
	
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
	return fetchList[models.Plant](f, "Plants", "GetVegetativePlants", "GET", url, nil)
}

// GetWasteById (Plants)
// GET {{baseUrl}}/plants/v2/waste/{id}/plant
//   licenseNumber (required): The License Number of the Facility for which to return the list of waste plants.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetWasteById(id string, licenseNumber string, pageNumber string, pageSize string) (PaginatedResponse[models.PlantsWaste], error) {
	url := "{{baseUrl}}/plants/v2/waste/{id}/plant"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.PlantsWaste](f, "Plants", "GetWasteById", "GET", url, nil)
}

// GetWastePackageById (Plants)
// GET {{baseUrl}}/plants/v2/waste/{id}/package
//   licenseNumber (required): The License Number of the Facility for which to return the list of waste plants.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetWastePackageById(id string, licenseNumber string, pageNumber string, pageSize string) (PaginatedResponse[models.WastePackage], error) {
	url := "{{baseUrl}}/plants/v2/waste/{id}/package"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.WastePackage](f, "Plants", "GetWastePackageById", "GET", url, nil)
}

// UpdateAdjustPlants (Plants)
// PUT {{baseUrl}}/plants/v2/adjust
//   licenseNumber (required): The License Number of the Facility for which to update the Plant count.
func (f *Fetcher) UpdateAdjustPlants(body []models.AdjustPlantsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/adjust"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "UpdateAdjustPlants", "PUT", url, body)
}

// UpdateGrowthPhase (Plants)
// PUT {{baseUrl}}/plants/v2/growthphase
//   licenseNumber (required): The License Number of the Facility for which to update the list of plants growth phase.
func (f *Fetcher) UpdateGrowthPhase(body []models.PlantsGrowthPhaseRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/growthphase"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "UpdateGrowthPhase", "PUT", url, body)
}

// UpdateHarvest (Plants)
// PUT {{baseUrl}}/plants/v2/harvest
//   licenseNumber (required): The License Number of the Facility for which to update the list of Plant harvests.
func (f *Fetcher) UpdateHarvest(body []models.HarvestRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/harvest"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "UpdateHarvest", "PUT", url, body)
}

// UpdateMerge (Plants)
// PUT {{baseUrl}}/plants/v2/merge
//   licenseNumber (required): The License Number of the Facility for which to update the list of plants merged.
func (f *Fetcher) UpdateMerge(body []models.PlantsMergeRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/merge"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "UpdateMerge", "PUT", url, body)
}

// UpdatePlantsLocation (Plants)
// PUT {{baseUrl}}/plants/v2/location
//   licenseNumber (required): The License Number of the Facility for which to update the list of plants moved.
func (f *Fetcher) UpdatePlantsLocation(body []models.PlantsLocationRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/location"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "UpdatePlantsLocation", "PUT", url, body)
}

// UpdatePlantsStrain (Plants)
// PUT {{baseUrl}}/plants/v2/strain
//   licenseNumber (required): The License Number of the Facility for which to update the list of Plant strains.
func (f *Fetcher) UpdatePlantsStrain(body []models.StrainRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/strain"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "UpdatePlantsStrain", "PUT", url, body)
}

// UpdatePlantsTag (Plants)
// PUT {{baseUrl}}/plants/v2/tag
//   licenseNumber (required): The License Number of the Facility for which to upate the list of Plant tags.
func (f *Fetcher) UpdatePlantsTag(body []models.TagRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/tag"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "UpdatePlantsTag", "PUT", url, body)
}

// UpdateSplit (Plants)
// PUT {{baseUrl}}/plants/v2/split
//   licenseNumber (required): The License Number of the Facility for which to update Plant splits.
func (f *Fetcher) UpdateSplit(body []models.PlantsSplitRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plants/v2/split"
	
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
	return fetchOne[models.WriteResponse](f, "Plants", "UpdateSplit", "PUT", url, body)
}
