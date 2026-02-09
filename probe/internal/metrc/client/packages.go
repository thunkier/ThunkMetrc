package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateAdjustPackages (Packages)
// POST {{baseUrl}}/packages/v2/adjust
//   licenseNumber (required): The License Number of the Facility for which to record the list of adjustments.
func (f *Fetcher) CreateAdjustPackages(body []models.AdjustPackagesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/adjust"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "CreateAdjustPackages", "POST", url, body)
}

// CreatePackagesPackages (Packages)
// POST {{baseUrl}}/packages/v2
//   licenseNumber (required): The License Number of the Facility for which to record the list of new packages.
func (f *Fetcher) CreatePackagesPackages(body []models.PackagesPackagesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "CreatePackagesPackages", "POST", url, body)
}

// CreatePackagesPlantings (Packages)
// POST {{baseUrl}}/packages/v2/plantings
//   licenseNumber (required): The License Number of the Facility for which to record the list of new plantings.
func (f *Fetcher) CreatePackagesPlantings(body []models.PackagesPlantingsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/plantings"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "CreatePackagesPlantings", "POST", url, body)
}

// CreateTesting (Packages)
// POST {{baseUrl}}/packages/v2/testing
//   licenseNumber (required): The License Number of the Facility for which to record the list of new packages for testing.
func (f *Fetcher) CreateTesting(body []models.TestingRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/testing"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "CreateTesting", "POST", url, body)
}

// DeletePackagesById (Packages)
// DELETE {{baseUrl}}/packages/v2/{id}
//   licenseNumber (required): The License Number of the Facility for which to discontinue.
func (f *Fetcher) DeletePackagesById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/packages/v2/{id}"
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
	return fetchOne[map[string]any](f, "Packages", "DeletePackagesById", "DELETE", url, nil)
}

// FinishPackages (Packages)
// PUT {{baseUrl}}/packages/v2/finish
//   licenseNumber (required): The License Number of the Facility for which to update the list of finish packages.
func (f *Fetcher) FinishPackages(body []models.FinishPackagesRequest, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/packages/v2/finish"
	
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
	return fetchOne[map[string]any](f, "Packages", "FinishPackages", "PUT", url, body)
}

// FinishedgoodFlag (Packages)
// PUT {{baseUrl}}/packages/v2/finishedgood/flag
//   licenseNumber (required): The License Number of the Facility for which to flag finished goods.
func (f *Fetcher) FinishedgoodFlag(body []models.FinishedgoodFlagRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/finishedgood/flag"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "FinishedgoodFlag", "PUT", url, body)
}

// FinishedgoodUnflag (Packages)
// PUT {{baseUrl}}/packages/v2/finishedgood/unflag
//   licenseNumber (required): The License Number of the Facility for which to flag finished goods.
func (f *Fetcher) FinishedgoodUnflag(body []models.FinishedgoodUnflagRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/finishedgood/unflag"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "FinishedgoodUnflag", "PUT", url, body)
}

// GetActivePackages (Packages)
// GET {{baseUrl}}/packages/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active packages.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActivePackages(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.PackagesPackage, error) {
	url := "{{baseUrl}}/packages/v2/active"
	
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
	return fetchList[models.PackagesPackage](f, "Packages", "GetActivePackages", "GET", url, nil)
}

// GetAdjustReasons (Packages)
// GET {{baseUrl}}/packages/v2/adjust/reasons
//   licenseNumber (required): The License Number of the Facility for which to return the list of adjustment reasons.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetAdjustReasons(licenseNumber string, pageNumber string, pageSize string) ([]models.AdjustReason, error) {
	url := "{{baseUrl}}/packages/v2/adjust/reasons"
	
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
	return fetchList[models.AdjustReason](f, "Packages", "GetAdjustReasons", "GET", url, nil)
}

// GetAdjustments (Packages)
// GET {{baseUrl}}/packages/v2/adjustments
//   licenseNumber (required): The License Number of the Facility for which to return the list of Package Adjustments.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetAdjustments(licenseNumber string, pageNumber string, pageSize string) (map[string]any, error) {
	url := "{{baseUrl}}/packages/v2/adjustments"
	
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
	return fetchOne[map[string]any](f, "Packages", "GetAdjustments", "GET", url, nil)
}

// GetInTransitPackages (Packages)
// GET {{baseUrl}}/packages/v2/intransit
//   licenseNumber (required): The License Number of the Facility for which to return the list of packages in transit.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInTransitPackages(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.InTransit, error) {
	url := "{{baseUrl}}/packages/v2/intransit"
	
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
	return fetchList[models.InTransit](f, "Packages", "GetInTransitPackages", "GET", url, nil)
}

// GetInactivePackages (Packages)
// GET {{baseUrl}}/packages/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive packages.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactivePackages(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.PackagesPackage, error) {
	url := "{{baseUrl}}/packages/v2/inactive"
	
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
	return fetchList[models.PackagesPackage](f, "Packages", "GetInactivePackages", "GET", url, nil)
}

// GetLabSamples (Packages)
// GET {{baseUrl}}/packages/v2/labsamples
//   licenseNumber (required): The License Number of the Facility for which to return the list of lab sample packages that have been created/sent for testing.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetLabSamples(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.PackagesPackage, error) {
	url := "{{baseUrl}}/packages/v2/labsamples"
	
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
	return fetchList[models.PackagesPackage](f, "Packages", "GetLabSamples", "GET", url, nil)
}

// GetOnHoldPackages (Packages)
// GET {{baseUrl}}/packages/v2/onhold
//   licenseNumber (required): The License Number of the Facility for which to return the list of packages on hold.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetOnHoldPackages(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.PackagesPackage, error) {
	url := "{{baseUrl}}/packages/v2/onhold"
	
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
	return fetchList[models.PackagesPackage](f, "Packages", "GetOnHoldPackages", "GET", url, nil)
}

// GetPackagesById (Packages)
// GET {{baseUrl}}/packages/v2/{id}
//   licenseNumber (required): If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetPackagesById(id string, licenseNumber string) (models.PackagesPackage, error) {
	url := "{{baseUrl}}/packages/v2/{id}"
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
	return fetchOne[models.PackagesPackage](f, "Packages", "GetPackagesById", "GET", url, nil)
}

// GetPackagesByLabel (Packages)
// GET {{baseUrl}}/packages/v2/{label}
//   licenseNumber (required): If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetPackagesByLabel(label string, licenseNumber string) ([]models.PackagesPackage, error) {
	url := "{{baseUrl}}/packages/v2/{label}"
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
	return fetchList[models.PackagesPackage](f, "Packages", "GetPackagesByLabel", "GET", url, nil)
}

// GetPackagesTypes (Packages)
// GET {{baseUrl}}/packages/v2/types
func (f *Fetcher) GetPackagesTypes() (map[string]any, error) {
	url := "{{baseUrl}}/packages/v2/types"
	return fetchOne[map[string]any](f, "Packages", "GetPackagesTypes", "GET", url, nil)
}

// GetSourceHarvestById (Packages)
// GET {{baseUrl}}/packages/v2/{id}/source/harvests
//   licenseNumber (required): If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetSourceHarvestById(id string, licenseNumber string) (models.SourceHarvest, error) {
	url := "{{baseUrl}}/packages/v2/{id}/source/harvests"
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
	return fetchOne[models.SourceHarvest](f, "Packages", "GetSourceHarvestById", "GET", url, nil)
}

// GetTransferred (Packages)
// GET {{baseUrl}}/packages/v2/transferred
//   licenseNumber (required): The License Number of the Facility for which to return the list of transferred packages.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetTransferred(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.PackagesPackage, error) {
	url := "{{baseUrl}}/packages/v2/transferred"
	
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
	return fetchList[models.PackagesPackage](f, "Packages", "GetTransferred", "GET", url, nil)
}

// UnfinishPackages (Packages)
// PUT {{baseUrl}}/packages/v2/unfinish
//   licenseNumber (required): The License Number of the Facility for which to update list of unfinish packages.
func (f *Fetcher) UnfinishPackages(body []models.UnfinishPackagesRequest, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/packages/v2/unfinish"
	
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
	return fetchOne[map[string]any](f, "Packages", "UnfinishPackages", "PUT", url, body)
}

// UpdateAdjustPackages (Packages)
// PUT {{baseUrl}}/packages/v2/adjust
//   licenseNumber (required): The License Number of the Facility for which to record the list of adjustments.
func (f *Fetcher) UpdateAdjustPackages(body []models.AdjustPackagesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/adjust"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateAdjustPackages", "PUT", url, body)
}

// UpdateDecontaminate (Packages)
// PUT {{baseUrl}}/packages/v2/decontaminate
//   licenseNumber (required): The License Number of the Facility for which to update the list of product decontaminations.
func (f *Fetcher) UpdateDecontaminate(body []models.DecontaminateRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/decontaminate"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateDecontaminate", "PUT", url, body)
}

// UpdateDonationFlag (Packages)
// PUT {{baseUrl}}/packages/v2/donation/flag
//   licenseNumber (required): The License Number of the Facility for which to update list of flagged donations.
func (f *Fetcher) UpdateDonationFlag(body []models.DonationFlagRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/donation/flag"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateDonationFlag", "PUT", url, body)
}

// UpdateDonationUnflag (Packages)
// PUT {{baseUrl}}/packages/v2/donation/unflag
//   licenseNumber (required): The License Number of the Facility for which to update list of unflaged donations.
func (f *Fetcher) UpdateDonationUnflag(body []models.DonationUnflagRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/donation/unflag"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateDonationUnflag", "PUT", url, body)
}

// UpdateExternalid (Packages)
// PUT {{baseUrl}}/packages/v2/externalid
//   licenseNumber (required): The License Number of the Facility for which to update list of change external Ids.
func (f *Fetcher) UpdateExternalid(body []models.ExternalidRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/externalid"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateExternalid", "PUT", url, body)
}

// UpdateItem (Packages)
// PUT {{baseUrl}}/packages/v2/item
//   licenseNumber (required): The License Number of the Facility for which to update list of changed items.
func (f *Fetcher) UpdateItem(body []models.ItemRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/item"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateItem", "PUT", url, body)
}

// UpdateLabtestsRequired (Packages)
// PUT {{baseUrl}}/packages/v2/labtests/required
//   licenseNumber (required): The License Number of the Facility for which to update the list of required lab test batches.
func (f *Fetcher) UpdateLabtestsRequired(body []models.LabtestsRequiredRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/labtests/required"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateLabtestsRequired", "PUT", url, body)
}

// UpdateNote (Packages)
// PUT {{baseUrl}}/packages/v2/note
//   licenseNumber (required): The License Number of the Facility for which to update list of change notes.
func (f *Fetcher) UpdateNote(body []models.NoteRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/note"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateNote", "PUT", url, body)
}

// UpdatePackagesLocation (Packages)
// PUT {{baseUrl}}/packages/v2/location
//   licenseNumber (required): The License Number of the Facility for which to update the list of change locations.
func (f *Fetcher) UpdatePackagesLocation(body []models.PackagesLocationRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/location"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdatePackagesLocation", "PUT", url, body)
}

// UpdateRemediate (Packages)
// PUT {{baseUrl}}/packages/v2/remediate
//   licenseNumber (required): The License Number of the Facility for which to update the list of product remediations.
func (f *Fetcher) UpdateRemediate(body []models.RemediateRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/remediate"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateRemediate", "PUT", url, body)
}

// UpdateTradeSampleFlag (Packages)
// PUT {{baseUrl}}/packages/v2/tradesample/flag
//   licenseNumber (required): The License Number of the Facility for which to update trade sample flags.
func (f *Fetcher) UpdateTradeSampleFlag(body []models.TradeSampleFlagRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/tradesample/flag"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateTradeSampleFlag", "PUT", url, body)
}

// UpdateTradeSampleUnflag (Packages)
// PUT {{baseUrl}}/packages/v2/tradesample/unflag
//   licenseNumber (required): The License Number of the Facility for which to update trade sample unflag.
func (f *Fetcher) UpdateTradeSampleUnflag(body []models.TradeSampleUnflagRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/tradesample/unflag"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateTradeSampleUnflag", "PUT", url, body)
}

// UpdateUseByDate (Packages)
// PUT {{baseUrl}}/packages/v2/usebydate
//   licenseNumber (required): The License Number of the Facility for which to update list of changed items.
func (f *Fetcher) UpdateUseByDate(body []models.UseByDateRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/packages/v2/usebydate"
	
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
	return fetchOne[models.WriteResponse](f, "Packages", "UpdateUseByDate", "PUT", url, body)
}
