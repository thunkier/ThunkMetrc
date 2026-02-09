package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateAdjustPlantBatches (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/adjust
//   licenseNumber (required): The License Number of the Facility for which to record adjustments.
func (f *Fetcher) CreateAdjustPlantBatches(body []models.AdjustPlantBatchesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/adjust"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreateAdjustPlantBatches", "POST", url, body)
}

// CreateGrowthPhase (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/growthphase
//   licenseNumber (required): The License Number of the Facility for which to change the growth phase.
func (f *Fetcher) CreateGrowthPhase(body []models.PlantBatchesGrowthPhaseRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/growthphase"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreateGrowthPhase", "POST", url, body)
}

// CreatePackagesFromMotherPlant (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/packages/frommotherplant
//   licenseNumber (required): The License Number of the Facility for which to create packages from a mother plant.
func (f *Fetcher) CreatePackagesFromMotherPlant(body []models.PackagesFromMotherPlantRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/packages/frommotherplant"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreatePackagesFromMotherPlant", "POST", url, body)
}

// CreatePlantBatchesAdditives (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/additives
//   licenseNumber (required): The License Number of the Facility for which to record plant additives.
func (f *Fetcher) CreatePlantBatchesAdditives(body []models.PlantBatchesAdditivesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/additives"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreatePlantBatchesAdditives", "POST", url, body)
}

// CreatePlantBatchesAdditivesUsingTemplate (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/additives/usingtemplate
//   licenseNumber (required): The License Number of the Facility for which to record plant additives.
func (f *Fetcher) CreatePlantBatchesAdditivesUsingTemplate(body []models.PlantBatchesAdditivesUsingTemplateRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/additives/usingtemplate"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreatePlantBatchesAdditivesUsingTemplate", "POST", url, body)
}

// CreatePlantBatchesPackages (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/packages
//   licenseNumber (required): The License Number of the Facility for which to create a package from the Plant Batch.
func (f *Fetcher) CreatePlantBatchesPackages(body []models.PlantBatchesPackagesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/packages"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreatePlantBatchesPackages", "POST", url, body)
}

// CreatePlantBatchesPlantings (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/plantings
//   licenseNumber (required): The License Number of the Facility for which to create the plantings.
func (f *Fetcher) CreatePlantBatchesPlantings(body []models.PlantBatchesPlantingsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/plantings"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreatePlantBatchesPlantings", "POST", url, body)
}

// CreatePlantBatchesWaste (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/waste
//   licenseNumber (required): The License Number of the Facility for which to record waste of the Plant Batch.
func (f *Fetcher) CreatePlantBatchesWaste(body []models.PlantBatchesWasteRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/waste"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreatePlantBatchesWaste", "POST", url, body)
}

// CreateSplit (PlantBatches)
// POST {{baseUrl}}/plantbatches/v2/split
//   licenseNumber (required): The License Number of the Facility for which to split the plant batches.
func (f *Fetcher) CreateSplit(body []models.PlantBatchesSplitRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/split"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "CreateSplit", "POST", url, body)
}

// DeletePlantBatches (PlantBatches)
// DELETE {{baseUrl}}/plantbatches/v2
//   licenseNumber (required): The License Number of the Facility for which to destroy the Plant Batch.
func (f *Fetcher) DeletePlantBatches(licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "DeletePlantBatches", "DELETE", url, nil)
}

// GetActivePlantBatches (PlantBatches)
// GET {{baseUrl}}/plantbatches/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active plant batches.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActivePlantBatches(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.PlantBatch, error) {
	url := "{{baseUrl}}/plantbatches/v2/active"
	
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
	return fetchList[models.PlantBatch](f, "PlantBatches", "GetActivePlantBatches", "GET", url, nil)
}

// GetInactivePlantBatches (PlantBatches)
// GET {{baseUrl}}/plantbatches/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive plant batches.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactivePlantBatches(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.PlantBatch, error) {
	url := "{{baseUrl}}/plantbatches/v2/inactive"
	
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
	return fetchList[models.PlantBatch](f, "PlantBatches", "GetInactivePlantBatches", "GET", url, nil)
}

// GetPlantBatchesById (PlantBatches)
// GET {{baseUrl}}/plantbatches/v2/{id}
//   licenseNumber (required): If specified, the Plant Batch will be validated against the specified License Number. If not specified, the Plant Batch will be validated against all of the User's current Facilities. Please note that if the Plant Batch is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetPlantBatchesById(id string, licenseNumber string) (models.PlantBatch, error) {
	url := "{{baseUrl}}/plantbatches/v2/{id}"
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
	return fetchOne[models.PlantBatch](f, "PlantBatches", "GetPlantBatchesById", "GET", url, nil)
}

// GetPlantBatchesTypes (PlantBatches)
// GET {{baseUrl}}/plantbatches/v2/types
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetPlantBatchesTypes(pageNumber string, pageSize string) ([]models.PlantBatchesType, error) {
	url := "{{baseUrl}}/plantbatches/v2/types"
	
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
	return fetchList[models.PlantBatchesType](f, "PlantBatches", "GetPlantBatchesTypes", "GET", url, nil)
}

// GetPlantBatchesWaste (PlantBatches)
// GET {{baseUrl}}/plantbatches/v2/waste
//   licenseNumber (required): The License Number of the Facility for which to return the list of waste plants.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetPlantBatchesWaste(licenseNumber string, pageNumber string, pageSize string) ([]models.PlantBatchesWaste, error) {
	url := "{{baseUrl}}/plantbatches/v2/waste"
	
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
	return fetchList[models.PlantBatchesWaste](f, "PlantBatches", "GetPlantBatchesWaste", "GET", url, nil)
}

// GetPlantBatchesWasteReasons (PlantBatches)
// GET {{baseUrl}}/plantbatches/v2/waste/reasons
//   licenseNumber (required): The License Number of the Facility for which to return a list of waste reasons.
func (f *Fetcher) GetPlantBatchesWasteReasons(licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/plantbatches/v2/waste/reasons"
	
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
	return fetchOne[map[string]any](f, "PlantBatches", "GetPlantBatchesWasteReasons", "GET", url, nil)
}

// UpdateName (PlantBatches)
// PUT {{baseUrl}}/plantbatches/v2/name
//   licenseNumber (required): The License Number of the Facility for which to rename the Plant Batch.
func (f *Fetcher) UpdateName(body []models.NameRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/name"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "UpdateName", "PUT", url, body)
}

// UpdatePlantBatchesLocation (PlantBatches)
// PUT {{baseUrl}}/plantbatches/v2/location
//   licenseNumber (required): The License Number of the Facility for which to change the Location of the Plant Batch.
func (f *Fetcher) UpdatePlantBatchesLocation(body []models.PlantBatchesLocationRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/location"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "UpdatePlantBatchesLocation", "PUT", url, body)
}

// UpdatePlantBatchesStrain (PlantBatches)
// PUT {{baseUrl}}/plantbatches/v2/strain
//   licenseNumber (required): The License Number of the Facility for which to change the Strain of the Plant Batch.
func (f *Fetcher) UpdatePlantBatchesStrain(body []models.PlantBatchesStrainRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/strain"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "UpdatePlantBatchesStrain", "PUT", url, body)
}

// UpdatePlantBatchesTag (PlantBatches)
// PUT {{baseUrl}}/plantbatches/v2/tag
//   licenseNumber (required): The License Number of the Facility for which to replace Plant Batch tags.
func (f *Fetcher) UpdatePlantBatchesTag(body []models.PlantBatchesTagRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/plantbatches/v2/tag"
	
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
	return fetchOne[models.WriteResponse](f, "PlantBatches", "UpdatePlantBatchesTag", "PUT", url, body)
}
