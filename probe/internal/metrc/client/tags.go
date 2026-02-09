package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// GetAvailablePackage (Tags)
// GET {{baseUrl}}/tags/v2/package/available
//   licenseNumber (required): The License Number for which to return available Package tags.
func (f *Fetcher) GetAvailablePackage(licenseNumber string) ([]models.Tag, error) {
	url := "{{baseUrl}}/tags/v2/package/available"
	
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
	return fetchList[models.Tag](f, "Tags", "GetAvailablePackage", "GET", url, nil)
}

// GetAvailablePlant (Tags)
// GET {{baseUrl}}/tags/v2/plant/available
//   licenseNumber (required): The License Number for which to return available tags.
func (f *Fetcher) GetAvailablePlant(licenseNumber string) ([]models.Tag, error) {
	url := "{{baseUrl}}/tags/v2/plant/available"
	
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
	return fetchList[models.Tag](f, "Tags", "GetAvailablePlant", "GET", url, nil)
}

// GetStagedTags (Tags)
// GET {{baseUrl}}/tags/v2/staged
//   licenseNumber (required): The License Number for which to return staged tags.
func (f *Fetcher) GetStagedTags(licenseNumber string) ([]models.Staged, error) {
	url := "{{baseUrl}}/tags/v2/staged"
	
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
	return fetchList[models.Staged](f, "Tags", "GetStagedTags", "GET", url, nil)
}
