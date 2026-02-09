package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateBrand (Items)
// POST {{baseUrl}}/items/v2/brand
//   licenseNumber (required): The License Number of the Facility for the Item Brands to create.
func (f *Fetcher) CreateBrand(body []models.ItemsBrandRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/items/v2/brand"
	
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
	return fetchOne[models.WriteResponse](f, "Items", "CreateBrand", "POST", url, body)
}

// CreateFile (Items)
// POST {{baseUrl}}/items/v2/file
//   licenseNumber (required): 
func (f *Fetcher) CreateFile(body []models.FileRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/items/v2/file"
	
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
	return fetchOne[models.WriteResponse](f, "Items", "CreateFile", "POST", url, body)
}

// CreateItems (Items)
// POST {{baseUrl}}/items/v2
//   licenseNumber (required): The License Number of the Facility for which to create new items.
func (f *Fetcher) CreateItems(body []models.ItemsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/items/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Items", "CreateItems", "POST", url, body)
}

// CreatePhoto (Items)
// POST {{baseUrl}}/items/v2/photo
//   licenseNumber (required): The License Number of the Facility for which to add an image.
func (f *Fetcher) CreatePhoto(body []models.PhotoRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/items/v2/photo"
	
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
	return fetchOne[models.WriteResponse](f, "Items", "CreatePhoto", "POST", url, body)
}

// DeleteBrandById (Items)
// DELETE {{baseUrl}}/items/v2/brand/{id}
//   licenseNumber (required): The License Number of the Facility for the Item Brand to delete.
func (f *Fetcher) DeleteBrandById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/items/v2/brand/{id}"
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
	return fetchOne[map[string]any](f, "Items", "DeleteBrandById", "DELETE", url, nil)
}

// DeleteItemsById (Items)
// DELETE {{baseUrl}}/items/v2/{id}
//   licenseNumber (required): The License Number of the Facility for the Item to delete.
func (f *Fetcher) DeleteItemsById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/items/v2/{id}"
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
	return fetchOne[map[string]any](f, "Items", "DeleteItemsById", "DELETE", url, nil)
}

// GetActiveItems (Items)
// GET {{baseUrl}}/items/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active items.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveItems(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Item, error) {
	url := "{{baseUrl}}/items/v2/active"
	
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
	return fetchList[models.Item](f, "Items", "GetActiveItems", "GET", url, nil)
}

// GetBrands (Items)
// GET {{baseUrl}}/items/v2/brands
//   licenseNumber (required): The License Number of the Facility for which to return the list of active item brands.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetBrands(licenseNumber string, pageNumber string, pageSize string) ([]models.Brand, error) {
	url := "{{baseUrl}}/items/v2/brands"
	
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
	return fetchList[models.Brand](f, "Items", "GetBrands", "GET", url, nil)
}

// GetCategories (Items)
// GET {{baseUrl}}/items/v2/categories
//   licenseNumber (required): If specified, the categories will be retrieved for the specified License Number. If not specified, all item categories will be returned.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetCategories(licenseNumber string, pageNumber string, pageSize string) ([]models.Category, error) {
	url := "{{baseUrl}}/items/v2/categories"
	
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
	return fetchList[models.Category](f, "Items", "GetCategories", "GET", url, nil)
}

// GetFileById (Items)
// GET {{baseUrl}}/items/v2/file/{id}
//   licenseNumber (required): The License Number of the Facility for which to return the file.
func (f *Fetcher) GetFileById(id string, licenseNumber string) (models.File, error) {
	url := "{{baseUrl}}/items/v2/file/{id}"
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
	return fetchOne[models.File](f, "Items", "GetFileById", "GET", url, nil)
}

// GetInactiveItems (Items)
// GET {{baseUrl}}/items/v2/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive items.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveItems(licenseNumber string, pageNumber string, pageSize string) ([]models.Item, error) {
	url := "{{baseUrl}}/items/v2/inactive"
	
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
	return fetchList[models.Item](f, "Items", "GetInactiveItems", "GET", url, nil)
}

// GetItemsById (Items)
// GET {{baseUrl}}/items/v2/{id}
//   licenseNumber (required): If specified, the Item will be validated against the specified License Number. If not specified, the Item will be validated against all of the User's current Facilities. Please note that if the Item is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetItemsById(id string, licenseNumber string) (models.Item, error) {
	url := "{{baseUrl}}/items/v2/{id}"
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
	return fetchOne[models.Item](f, "Items", "GetItemsById", "GET", url, nil)
}

// GetPhotoById (Items)
// GET {{baseUrl}}/items/v2/photo/{id}
//   licenseNumber (required): The License Number of the Facility for which to return the image.
func (f *Fetcher) GetPhotoById(id string, licenseNumber string) (models.Photo, error) {
	url := "{{baseUrl}}/items/v2/photo/{id}"
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
	return fetchOne[models.Photo](f, "Items", "GetPhotoById", "GET", url, nil)
}

// UpdateBrand (Items)
// PUT {{baseUrl}}/items/v2/brand
//   licenseNumber (required): The License Number of the Facility for the Item Brand updates.
func (f *Fetcher) UpdateBrand(body []models.ItemsBrandRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/items/v2/brand"
	
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
	return fetchOne[models.WriteResponse](f, "Items", "UpdateBrand", "PUT", url, body)
}

// UpdateItems (Items)
// PUT {{baseUrl}}/items/v2
//   licenseNumber (required): The License Number of the Facility for the Item updates.
func (f *Fetcher) UpdateItems(body []models.ItemsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/items/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Items", "UpdateItems", "PUT", url, body)
}
