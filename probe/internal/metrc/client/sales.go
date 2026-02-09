package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateDeliveries (Sales)
// POST {{baseUrl}}/sales/v2/deliveries
//   licenseNumber (required): The License Number of the Facility for which to record deliveries.
func (f *Fetcher) CreateDeliveries(body []models.SalesDeliveriesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "CreateDeliveries", "POST", url, body)
}

// CreateDeliveriesRetailerDepart (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer/depart
//   licenseNumber (required): The License Number of the Facility for which to record depart delivery.
func (f *Fetcher) CreateDeliveriesRetailerDepart(body []models.DeliveriesRetailerDepartRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer/depart"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "CreateDeliveriesRetailerDepart", "POST", url, body)
}

// CreateDeliveriesRetailerEnd (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer/end
//   licenseNumber (required): The License Number of the Facility for which to record end retailer deliveries.
func (f *Fetcher) CreateDeliveriesRetailerEnd(body []models.DeliveriesRetailerEndRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer/end"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "CreateDeliveriesRetailerEnd", "POST", url, body)
}

// CreateDeliveriesRetailerRestock (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer/restock
//   licenseNumber (required): The License Number of the Facility for which to record restock retailer deliveries.
func (f *Fetcher) CreateDeliveriesRetailerRestock(body []models.DeliveriesRetailerRestockRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer/restock"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "CreateDeliveriesRetailerRestock", "POST", url, body)
}

// CreateReceipts (Sales)
// POST {{baseUrl}}/sales/v2/receipts
//   licenseNumber (required): The License Number of the Facility for which to record receipts.
func (f *Fetcher) CreateReceipts(body []models.SalesReceiptsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/receipts"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "CreateReceipts", "POST", url, body)
}

// CreateSalesDeliveriesRetailer (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer
//   licenseNumber (required): The License Number of the Facility for which to record retailer deliveries.
func (f *Fetcher) CreateSalesDeliveriesRetailer(body []models.SalesDeliveriesRetailerRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "CreateSalesDeliveriesRetailer", "POST", url, body)
}

// DeleteDeliveryById (Sales)
// DELETE {{baseUrl}}/sales/v2/deliveries/{id}
//   licenseNumber (required): The License Number of the Facility for which to void delivery.
func (f *Fetcher) DeleteDeliveryById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/{id}"
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
	return fetchOne[map[string]any](f, "Sales", "DeleteDeliveryById", "DELETE", url, nil)
}

// DeleteDeliveryRetailerById (Sales)
// DELETE {{baseUrl}}/sales/v2/deliveries/retailer/{id}
//   licenseNumber (required): The License Number of the Facility for which to void retailer delivery.
func (f *Fetcher) DeleteDeliveryRetailerById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer/{id}"
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
	return fetchOne[map[string]any](f, "Sales", "DeleteDeliveryRetailerById", "DELETE", url, nil)
}

// DeleteReceiptById (Sales)
// DELETE {{baseUrl}}/sales/v2/receipts/{id}
//   licenseNumber (required): The License Number of the Facility for which to archive receipt.
func (f *Fetcher) DeleteReceiptById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/sales/v2/receipts/{id}"
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
	return fetchOne[map[string]any](f, "Sales", "DeleteReceiptById", "DELETE", url, nil)
}

// GetActiveDeliveries (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active deliveries.
//   lastModifiedStart (optional): The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.
//   lastModifiedEnd (optional): The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveDeliveries(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.ActiveDelivery, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/active"
	
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
	return fetchList[models.ActiveDelivery](f, "Sales", "GetActiveDeliveries", "GET", url, nil)
}

// GetActiveDeliveriesRetailer (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/retailer/active
//   licenseNumber (required): The License Number of the Facility for which to return active retailer deliveries.
//   lastModifiedStart (optional): The last modified start timestamp.
//   lastModifiedEnd (optional): The last modified end timestamp.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveDeliveriesRetailer(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.ActiveDeliveriesRetailer, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer/active"
	
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
	return fetchList[models.ActiveDeliveriesRetailer](f, "Sales", "GetActiveDeliveriesRetailer", "GET", url, nil)
}

// GetActiveReceipts (Sales)
// GET {{baseUrl}}/sales/v2/receipts/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active receipts.
//   lastModifiedStart (optional): The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.
//   lastModifiedEnd (optional): The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActiveReceipts(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.ActiveReceipt, error) {
	url := "{{baseUrl}}/sales/v2/receipts/active"
	
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
	return fetchList[models.ActiveReceipt](f, "Sales", "GetActiveReceipts", "GET", url, nil)
}

// GetCounties (Sales)
// GET {{baseUrl}}/sales/v2/counties
func (f *Fetcher) GetCounties() ([]models.County, error) {
	url := "{{baseUrl}}/sales/v2/counties"
	return fetchList[models.County](f, "Sales", "GetCounties", "GET", url, nil)
}

// GetCustomerTypes (Sales)
// GET {{baseUrl}}/sales/v2/customertypes
func (f *Fetcher) GetCustomerTypes() (map[string]any, error) {
	url := "{{baseUrl}}/sales/v2/customertypes"
	return fetchOne[map[string]any](f, "Sales", "GetCustomerTypes", "GET", url, nil)
}

// GetDeliveriesReturnReasons (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/returnreasons
//   licenseNumber (required): The License Number of the Facility for which to return the list of delivery return reasons.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetDeliveriesReturnReasons(licenseNumber string, pageNumber string, pageSize string) ([]models.DeliveriesReturnReason, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/returnreasons"
	
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
	return fetchList[models.DeliveriesReturnReason](f, "Sales", "GetDeliveriesReturnReasons", "GET", url, nil)
}

// GetDeliveryRetailerById (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/retailer/{id}
//   licenseNumber (required): If specified, the Sales Delivery will be validated against the specified License Number. If not specified, the Sales Delivery will be validated against all of the User's current Facilities. Please note that if the Sales Delivery is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetDeliveryRetailerById(id string, licenseNumber string) (models.DeliveryRetailer, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer/{id}"
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
	return fetchOne[models.DeliveryRetailer](f, "Sales", "GetDeliveryRetailerById", "GET", url, nil)
}

// GetInactiveDeliveries (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive deliveries.
//   lastModifiedStart (optional): The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.
//   lastModifiedEnd (optional): The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveDeliveries(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.InactiveDelivery, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/inactive"
	
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
	return fetchList[models.InactiveDelivery](f, "Sales", "GetInactiveDeliveries", "GET", url, nil)
}

// GetInactiveDeliveriesRetailer (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/retailer/inactive
//   licenseNumber (required): The License Number of the Facility for which to return inactive retailer deliveries.
//   lastModifiedStart (optional): The last modified start timestamp.
//   lastModifiedEnd (optional): The last modified end timestamp.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveDeliveriesRetailer(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.InactiveDeliveriesRetailer, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer/inactive"
	
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
	return fetchList[models.InactiveDeliveriesRetailer](f, "Sales", "GetInactiveDeliveriesRetailer", "GET", url, nil)
}

// GetInactiveReceipts (Sales)
// GET {{baseUrl}}/sales/v2/receipts/inactive
//   licenseNumber (required): The License Number of the Facility for which to return the list of inactive receipts.
//   lastModifiedStart (optional): The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.
//   lastModifiedEnd (optional): The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetInactiveReceipts(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.InactiveReceipt, error) {
	url := "{{baseUrl}}/sales/v2/receipts/inactive"
	
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
	return fetchList[models.InactiveReceipt](f, "Sales", "GetInactiveReceipts", "GET", url, nil)
}

// GetPatientRegistrationLocations (Sales)
// GET {{baseUrl}}/sales/v2/patientregistration/locations
func (f *Fetcher) GetPatientRegistrationLocations() ([]models.PatientRegistrationLocation, error) {
	url := "{{baseUrl}}/sales/v2/patientregistration/locations"
	return fetchList[models.PatientRegistrationLocation](f, "Sales", "GetPatientRegistrationLocations", "GET", url, nil)
}

// GetPaymentTypes (Sales)
// GET {{baseUrl}}/sales/v2/paymenttypes
//   licenseNumber (required): The License Number of the Facility for which to return the list of payment types.
func (f *Fetcher) GetPaymentTypes(licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/sales/v2/paymenttypes"
	
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
	return fetchOne[map[string]any](f, "Sales", "GetPaymentTypes", "GET", url, nil)
}

// GetReceiptById (Sales)
// GET {{baseUrl}}/sales/v2/receipts/{id}
//   licenseNumber (required): If specified, the Sales Receipt will be validated against the specified License Number. If not specified, the Sales Receipt will be validated against all of the User's current Facilities. Please note that if the Sales Receipt is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetReceiptById(id string, licenseNumber string) (models.InactiveReceipt, error) {
	url := "{{baseUrl}}/sales/v2/receipts/{id}"
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
	return fetchOne[models.InactiveReceipt](f, "Sales", "GetReceiptById", "GET", url, nil)
}

// GetReceiptsExternalByExternalNumber (Sales)
// GET {{baseUrl}}/sales/v2/receipts/external/{externalNumber}
//   licenseNumber (required): If specified, the External Sales Receipt Number will be validated against the specified License Number. If not specified, the External Sales Receipt Number will be validated against all of the User's current Facilities. Please note that if the External Sales Receipt Number is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetReceiptsExternalByExternalNumber(externalNumber string, licenseNumber string) ([]models.ReceiptsExternalByExternalNumber, error) {
	url := "{{baseUrl}}/sales/v2/receipts/external/{externalNumber}"
	url = strings.ReplaceAll(url, "{"+"externalNumber"+"}", externalNumber)
	
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
	return fetchList[models.ReceiptsExternalByExternalNumber](f, "Sales", "GetReceiptsExternalByExternalNumber", "GET", url, nil)
}

// GetSalesDeliveryById (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/{id}
//   licenseNumber (required): If specified, the Sales Delivery will be validated against the specified License Number. If not specified, the Sales Delivery will be validated against all of the User's current Facilities. Please note that if the Sales Delivery is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetSalesDeliveryById(id string, licenseNumber string) (models.SalesDelivery, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/{id}"
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
	return fetchOne[models.SalesDelivery](f, "Sales", "GetSalesDeliveryById", "GET", url, nil)
}

// UpdateDeliveries (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries
//   licenseNumber (required): The License Number of the Facility for which to update deliveries.
func (f *Fetcher) UpdateDeliveries(body []models.SalesDeliveriesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateDeliveries", "PUT", url, body)
}

// UpdateDeliveriesComplete (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/complete
//   licenseNumber (required): The License Number of the Facility for which to update deliveries completed.
func (f *Fetcher) UpdateDeliveriesComplete(body []models.DeliveriesCompleteRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/complete"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateDeliveriesComplete", "PUT", url, body)
}

// UpdateDeliveriesHub (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/hub
//   licenseNumber (required): The License Number of the Facility for which to update hub transporters.
func (f *Fetcher) UpdateDeliveriesHub(body []models.DeliveriesHubRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/hub"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateDeliveriesHub", "PUT", url, body)
}

// UpdateDeliveriesHubAccept (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/hub/accept
//   licenseNumber (required): The License Number of the Facility for which to update hub deliveries accepted.
func (f *Fetcher) UpdateDeliveriesHubAccept(body []models.DeliveriesHubAcceptRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/hub/accept"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateDeliveriesHubAccept", "PUT", url, body)
}

// UpdateDeliveriesHubDepart (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/hub/depart
//   licenseNumber (required): The License Number of the Facility for which to update hub delivery departures.
func (f *Fetcher) UpdateDeliveriesHubDepart(body []models.DeliveriesHubDepartRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/hub/depart"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateDeliveriesHubDepart", "PUT", url, body)
}

// UpdateDeliveriesHubVerifyID (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/hub/verifyID
//   licenseNumber (required): The License Number of the Facility for which to update hub delivery Id's.
func (f *Fetcher) UpdateDeliveriesHubVerifyID(body []models.DeliveriesHubVerifyIDRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/hub/verifyID"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateDeliveriesHubVerifyID", "PUT", url, body)
}

// UpdateDeliveriesRetailer (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/retailer
//   licenseNumber (required): The License Number of the Facility for which to update retailer deliveries.
func (f *Fetcher) UpdateDeliveriesRetailer(body []models.SalesDeliveriesRetailerRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/deliveries/retailer"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateDeliveriesRetailer", "PUT", url, body)
}

// UpdateReceipts (Sales)
// PUT {{baseUrl}}/sales/v2/receipts
//   licenseNumber (required): The License Number of the Facility for which to update receipts.
func (f *Fetcher) UpdateReceipts(body []models.SalesReceiptsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/receipts"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateReceipts", "PUT", url, body)
}

// UpdateReceiptsFinalize (Sales)
// PUT {{baseUrl}}/sales/v2/receipts/finalize
//   licenseNumber (required): The License Number of the Facility for which to update finalized receipts.
func (f *Fetcher) UpdateReceiptsFinalize(body []models.ReceiptsFinalizeRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/receipts/finalize"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateReceiptsFinalize", "PUT", url, body)
}

// UpdateReceiptsUnfinalize (Sales)
// PUT {{baseUrl}}/sales/v2/receipts/unfinalize
//   licenseNumber (required): The License Number of the Facility for which to update unfinalized receipts.
func (f *Fetcher) UpdateReceiptsUnfinalize(body []models.ReceiptsUnfinalizeRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/sales/v2/receipts/unfinalize"
	
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
	return fetchOne[models.WriteResponse](f, "Sales", "UpdateReceiptsUnfinalize", "PUT", url, body)
}
