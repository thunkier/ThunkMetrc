package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateAssociate (RetailId)
// POST {{baseUrl}}/retailid/v2/associate
//   licenseNumber (required): The License Number of the Facility for which to associate the packages and QR codes provided.
func (f *Fetcher) CreateAssociate(body []models.AssociateRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/retailid/v2/associate"
	
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
	return fetchOne[models.WriteResponse](f, "RetailId", "CreateAssociate", "POST", url, body)
}

// CreateGenerate (RetailId)
// POST {{baseUrl}}/retailid/v2/generate
//   licenseNumber (required): The License Number of the Facility for which to associate the packages and QR codes provided.
func (f *Fetcher) CreateGenerate(body models.GenerateRequest, licenseNumber string) ([]models.WriteResponse, error) {
	url := "{{baseUrl}}/retailid/v2/generate"
	
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
	return fetchList[models.WriteResponse](f, "RetailId", "CreateGenerate", "POST", url, body)
}

// CreateMerge (RetailId)
// POST {{baseUrl}}/retailid/v2/merge
//   licenseNumber (required): The License Number of the Facility.
func (f *Fetcher) CreateMerge(body models.RetailIdMergeRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/retailid/v2/merge"
	
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
	return fetchOne[models.WriteResponse](f, "RetailId", "CreateMerge", "POST", url, body)
}

// CreatePackagesInfo (RetailId)
// POST {{baseUrl}}/retailid/v2/packages/info
//   licenseNumber (required): The License Number of the Facility
func (f *Fetcher) CreatePackagesInfo(body models.PackagesInfoRequest, licenseNumber string) ([]models.WriteResponse, error) {
	url := "{{baseUrl}}/retailid/v2/packages/info"
	
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
	return fetchList[models.WriteResponse](f, "RetailId", "CreatePackagesInfo", "POST", url, body)
}

// GetAllotment (RetailId)
// GET {{baseUrl}}/retailid/v2/allotment
//   licenseNumber (required): The License Number of the Facility for which to retrieve available Retail Item ID quota.
func (f *Fetcher) GetAllotment(licenseNumber string) ([]models.Allotment, error) {
	url := "{{baseUrl}}/retailid/v2/allotment"
	
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
	return fetchList[models.Allotment](f, "RetailId", "GetAllotment", "GET", url, nil)
}

// GetReceiveByLabel (RetailId)
// GET {{baseUrl}}/retailid/v2/receive/{label}
//   licenseNumber (required): If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities.
func (f *Fetcher) GetReceiveByLabel(label string, licenseNumber string) ([]models.Receive, error) {
	url := "{{baseUrl}}/retailid/v2/receive/{label}"
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
	return fetchList[models.Receive](f, "RetailId", "GetReceiveByLabel", "GET", url, nil)
}

// GetReceiveQrByShortCode (RetailId)
// GET {{baseUrl}}/retailid/v2/receive/qr/{shortCode}
//   licenseNumber (required): The License Number of the Facility
func (f *Fetcher) GetReceiveQrByShortCode(shortCode string, licenseNumber string) ([]models.ReceiveQrByShortCode, error) {
	url := "{{baseUrl}}/retailid/v2/receive/qr/{shortCode}"
	url = strings.ReplaceAll(url, "{"+"shortCode"+"}", shortCode)
	
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
	return fetchList[models.ReceiveQrByShortCode](f, "RetailId", "GetReceiveQrByShortCode", "GET", url, nil)
}
