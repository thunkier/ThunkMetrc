package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreateHubArrive (Transfers)
// POST {{baseUrl}}/transfers/v2/hub/arrive
//   licenseNumber (required): The License Number of the Facility for which to arrive the transfer.
func (f *Fetcher) CreateHubArrive(body []models.HubArriveRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transfers/v2/hub/arrive"
	
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
	return fetchOne[models.WriteResponse](f, "Transfers", "CreateHubArrive", "POST", url, body)
}

// CreateHubCheckin (Transfers)
// POST {{baseUrl}}/transfers/v2/hub/checkin
//   licenseNumber (required): The License Number of the Facility for which to arrive the transfer.
func (f *Fetcher) CreateHubCheckin(body []models.HubCheckinRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transfers/v2/hub/checkin"
	
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
	return fetchOne[models.WriteResponse](f, "Transfers", "CreateHubCheckin", "POST", url, body)
}

// CreateHubCheckout (Transfers)
// POST {{baseUrl}}/transfers/v2/hub/checkout
//   licenseNumber (required): The License Number of the Facility for which to arrive the transfer.
func (f *Fetcher) CreateHubCheckout(body []models.HubCheckoutRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transfers/v2/hub/checkout"
	
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
	return fetchOne[models.WriteResponse](f, "Transfers", "CreateHubCheckout", "POST", url, body)
}

// CreateHubDepart (Transfers)
// POST {{baseUrl}}/transfers/v2/hub/depart
//   licenseNumber (required): The License Number of the Facility for which to depart the transfer.
func (f *Fetcher) CreateHubDepart(body []models.HubDepartRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transfers/v2/hub/depart"
	
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
	return fetchOne[models.WriteResponse](f, "Transfers", "CreateHubDepart", "POST", url, body)
}

// CreateIncomingExternal (Transfers)
// POST {{baseUrl}}/transfers/v2/external/incoming
//   licenseNumber (required): The License Number for which to create a Shipment Plan.
func (f *Fetcher) CreateIncomingExternal(body []models.TransfersExternalRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transfers/v2/external/incoming"
	
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
	return fetchOne[models.WriteResponse](f, "Transfers", "CreateIncomingExternal", "POST", url, body)
}

// CreateOutgoingTemplates (Transfers)
// POST {{baseUrl}}/transfers/v2/templates/outgoing
//   licenseNumber (required): The License Number of the Facility for which to create a Shipment Template.
func (f *Fetcher) CreateOutgoingTemplates(body []models.TransfersTemplatesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transfers/v2/templates/outgoing"
	
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
	return fetchOne[models.WriteResponse](f, "Transfers", "CreateOutgoingTemplates", "POST", url, body)
}

// DeleteIncomingExternalById (Transfers)
// DELETE {{baseUrl}}/transfers/v2/external/incoming/{id}
//   licenseNumber (required): The License Number of the Facility for which to void the Transfer.
func (f *Fetcher) DeleteIncomingExternalById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/transfers/v2/external/incoming/{id}"
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
	return fetchOne[map[string]any](f, "Transfers", "DeleteIncomingExternalById", "DELETE", url, nil)
}

// DeleteOutgoingTemplateById (Transfers)
// DELETE {{baseUrl}}/transfers/v2/templates/outgoing/{id}
//   licenseNumber (required): The License Number of the Facility for which to archive the template.
func (f *Fetcher) DeleteOutgoingTemplateById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/transfers/v2/templates/outgoing/{id}"
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
	return fetchOne[map[string]any](f, "Transfers", "DeleteOutgoingTemplateById", "DELETE", url, nil)
}

// GetDeliveriesPackagesStates (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/packages/states
func (f *Fetcher) GetDeliveriesPackagesStates() (map[string]any, error) {
	url := "{{baseUrl}}/transfers/v2/deliveries/packages/states"
	return fetchOne[map[string]any](f, "Transfers", "GetDeliveriesPackagesStates", "GET", url, nil)
}

// GetDeliveryPackageById (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/{id}/packages
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetDeliveryPackageById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.DeliveryPackage], error) {
	url := "{{baseUrl}}/transfers/v2/deliveries/{id}/packages"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.DeliveryPackage](f, "Transfers", "GetDeliveryPackageById", "GET", url, nil)
}

// GetDeliveryPackageRequiredlabtestbatchById (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/package/{id}/requiredlabtestbatches
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetDeliveryPackageRequiredlabtestbatchById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.DeliveryPackageRequiredlabtestbatch], error) {
	url := "{{baseUrl}}/transfers/v2/deliveries/package/{id}/requiredlabtestbatches"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.DeliveryPackageRequiredlabtestbatch](f, "Transfers", "GetDeliveryPackageRequiredlabtestbatchById", "GET", url, nil)
}

// GetDeliveryPackageWholesaleById (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/{id}/packages/wholesale
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetDeliveryPackageWholesaleById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.DeliveryPackageWholesale], error) {
	url := "{{baseUrl}}/transfers/v2/deliveries/{id}/packages/wholesale"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.DeliveryPackageWholesale](f, "Transfers", "GetDeliveryPackageWholesaleById", "GET", url, nil)
}

// GetDeliveryTransporterById (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/{id}/transporters
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetDeliveryTransporterById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.DeliveryTransporter], error) {
	url := "{{baseUrl}}/transfers/v2/deliveries/{id}/transporters"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.DeliveryTransporter](f, "Transfers", "GetDeliveryTransporterById", "GET", url, nil)
}

// GetDeliveryTransporterDetailById (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/{id}/transporters/details
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetDeliveryTransporterDetailById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.DeliveryTransporterDetail], error) {
	url := "{{baseUrl}}/transfers/v2/deliveries/{id}/transporters/details"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.DeliveryTransporterDetail](f, "Transfers", "GetDeliveryTransporterDetailById", "GET", url, nil)
}

// GetHub (Transfers)
// GET {{baseUrl}}/transfers/v2/hub
//   licenseNumber (required): The License Number of the Facility for which to return shipment plans.
//   lastModifiedStart (optional): The last modified start timestamp. If specified, also specifying any of the estimated arrival date parameters will result in an error.
//   lastModifiedEnd (optional): The last modified end timestamp. If specified, also specifying any of the estimated arrival date parameters will result in an error.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetHub(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Hub, error) {
	url := "{{baseUrl}}/transfers/v2/hub"
	
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
	return fetchList[models.Hub](f, "Transfers", "GetHub", "GET", url, nil)
}

// GetIncomingTransfers (Transfers)
// GET {{baseUrl}}/transfers/v2/incoming
//   licenseNumber (required): The License Number of the Facility for which to return incoming transfers.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetIncomingTransfers(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Transfer, error) {
	url := "{{baseUrl}}/transfers/v2/incoming"
	
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
	return fetchList[models.Transfer](f, "Transfers", "GetIncomingTransfers", "GET", url, nil)
}

// GetManifestHtmlById (Transfers)
// GET {{baseUrl}}/transfers/v2/manifest/{id}/html
//   licenseNumber (required): 
func (f *Fetcher) GetManifestHtmlById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/transfers/v2/manifest/{id}/html"
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
	return fetchOne[map[string]any](f, "Transfers", "GetManifestHtmlById", "GET", url, nil)
}

// GetManifestPdfById (Transfers)
// GET {{baseUrl}}/transfers/v2/manifest/{id}/pdf
//   licenseNumber (required): 
func (f *Fetcher) GetManifestPdfById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/transfers/v2/manifest/{id}/pdf"
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
	return fetchOne[map[string]any](f, "Transfers", "GetManifestPdfById", "GET", url, nil)
}

// GetOutgoingTemplateDeliveryById (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing/{id}/deliveries
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetOutgoingTemplateDeliveryById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.TemplateDelivery], error) {
	url := "{{baseUrl}}/transfers/v2/templates/outgoing/{id}/deliveries"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.TemplateDelivery](f, "Transfers", "GetOutgoingTemplateDeliveryById", "GET", url, nil)
}

// GetOutgoingTemplateDeliveryPackageById (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/packages
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetOutgoingTemplateDeliveryPackageById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.TemplateDeliveryPackage], error) {
	url := "{{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/packages"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.TemplateDeliveryPackage](f, "Transfers", "GetOutgoingTemplateDeliveryPackageById", "GET", url, nil)
}

// GetOutgoingTemplateDeliveryTransporterById (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/transporters
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetOutgoingTemplateDeliveryTransporterById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.TemplateDeliveryTransporter], error) {
	url := "{{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/transporters"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.TemplateDeliveryTransporter](f, "Transfers", "GetOutgoingTemplateDeliveryTransporterById", "GET", url, nil)
}

// GetOutgoingTemplateDeliveryTransporterDetailById (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/transporters/details
func (f *Fetcher) GetOutgoingTemplateDeliveryTransporterDetailById(id string) (models.TemplateDeliveryTransporterDetail, error) {
	url := "{{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/transporters/details"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	return fetchOne[models.TemplateDeliveryTransporterDetail](f, "Transfers", "GetOutgoingTemplateDeliveryTransporterDetailById", "GET", url, nil)
}

// GetOutgoingTemplates (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing
//   licenseNumber (required): The License Number of the Facility for which to return outgoing shipment templates.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetOutgoingTemplates(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Template, error) {
	url := "{{baseUrl}}/transfers/v2/templates/outgoing"
	
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
	return fetchList[models.Template](f, "Transfers", "GetOutgoingTemplates", "GET", url, nil)
}

// GetOutgoingTransfers (Transfers)
// GET {{baseUrl}}/transfers/v2/outgoing
//   licenseNumber (required): The License Number of the Facility for which to return outgoing transfers.
//   lastModifiedStart (optional): The last modified start timestamp
//   lastModifiedEnd (optional): The last modified end timestamp
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetOutgoingTransfers(licenseNumber string, lastModifiedStart string, lastModifiedEnd string, pageNumber string, pageSize string) ([]models.Transfer, error) {
	url := "{{baseUrl}}/transfers/v2/outgoing"
	
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
	return fetchList[models.Transfer](f, "Transfers", "GetOutgoingTransfers", "GET", url, nil)
}

// GetRejectedTransfers (Transfers)
// GET {{baseUrl}}/transfers/v2/rejected
//   licenseNumber (required): The License Number of the Facility for which to return rejected transfers.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetRejectedTransfers(licenseNumber string, pageNumber string, pageSize string) ([]models.Transfer, error) {
	url := "{{baseUrl}}/transfers/v2/rejected"
	
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
	return fetchList[models.Transfer](f, "Transfers", "GetRejectedTransfers", "GET", url, nil)
}

// GetTransfersDeliveryById (Transfers)
// GET {{baseUrl}}/transfers/v2/{id}/deliveries
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetTransfersDeliveryById(id string, pageNumber string, pageSize string) (PaginatedResponse[models.TransfersDelivery], error) {
	url := "{{baseUrl}}/transfers/v2/{id}/deliveries"
	url = strings.ReplaceAll(url, "{"+"id"+"}", id)
	
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
	return fetchPaginated[models.TransfersDelivery](f, "Transfers", "GetTransfersDeliveryById", "GET", url, nil)
}

// GetTransfersTypes (Transfers)
// GET {{baseUrl}}/transfers/v2/types
//   licenseNumber (required): The License Number of the Facility for which to return transfer types.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetTransfersTypes(licenseNumber string, pageNumber string, pageSize string) ([]models.TransfersType, error) {
	url := "{{baseUrl}}/transfers/v2/types"
	
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
	return fetchList[models.TransfersType](f, "Transfers", "GetTransfersTypes", "GET", url, nil)
}

// UpdateIncomingExternal (Transfers)
// PUT {{baseUrl}}/transfers/v2/external/incoming
//   licenseNumber (required): The License Number of the Facility for which to update external incoming shipment plans.
func (f *Fetcher) UpdateIncomingExternal(body []models.TransfersExternalRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transfers/v2/external/incoming"
	
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
	return fetchOne[models.WriteResponse](f, "Transfers", "UpdateIncomingExternal", "PUT", url, body)
}

// UpdateOutgoingTemplates (Transfers)
// PUT {{baseUrl}}/transfers/v2/templates/outgoing
//   licenseNumber (required): The License Number of the Facility for which to update shipment plan templates.
func (f *Fetcher) UpdateOutgoingTemplates(body []models.TransfersTemplatesRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/transfers/v2/templates/outgoing"
	
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
	return fetchOne[models.WriteResponse](f, "Transfers", "UpdateOutgoingTemplates", "PUT", url, body)
}
