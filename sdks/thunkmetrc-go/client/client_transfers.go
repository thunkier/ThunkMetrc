package client

import (
	"context"
	"net/url"
)


// POST /transfers/v2/hub/arrive
// Arrive a transfer for a Facility.
// Permissions Required:
// - Manage Transfer Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateTransfersHubArrive(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/transfers/v2/hub/arrive", queryParams, body)
}

// POST /transfers/v2/hub/checkin
// CheckIn a transfer for a Facility.
// Permissions Required:
// - Manage Transfer Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateTransfersHubCheckin(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/transfers/v2/hub/checkin", queryParams, body)
}

// POST /transfers/v2/hub/checkout
// CheckOut a transfer for a Facility.
// Permissions Required:
// - Manage Transfer Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateTransfersHubCheckout(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/transfers/v2/hub/checkout", queryParams, body)
}

// POST /transfers/v2/hub/depart
// Depart a transfer for a Facility.
// Permissions Required:
// - Manage Transfer Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateTransfersHubDepart(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/transfers/v2/hub/depart", queryParams, body)
}

// POST /transfers/v2/external/incoming
// Creates external incoming shipment plans for a Facility.
// Permissions Required:
// - Manage Transfers
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateTransfersIncomingExternal(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/transfers/v2/external/incoming", queryParams, body)
}

// POST /transfers/v2/templates/outgoing
// Creates new transfer templates for a Facility.
// Permissions Required:
// - Manage Transfer Templates
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateTransfersOutgoingTemplates(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/transfers/v2/templates/outgoing", queryParams, body)
}

// DELETE /transfers/v2/external/incoming/{id}
// Voids an external incoming shipment plan for a Facility.
// Permissions Required:
// - Manage Transfers
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteTransfersIncomingExternalById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/transfers/v2/external/incoming/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE /transfers/v2/templates/outgoing/{id}
// Archives a transfer template for a Facility.
// Permissions Required:
// - Manage Transfer Templates
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteTransfersOutgoingTemplateById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/transfers/v2/templates/outgoing/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /transfers/v2/deliveries/packages/states
// Returns a list of available shipment Package states.
// Parameters:
func (c *MetrcClient) GetTransfersDeliveriesPackagesStates(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/transfers/v2/deliveries/packages/states", queryParams, nil)
}

// GET /transfers/v2/deliveries/{id}/packages
// Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersDeliveryPackageById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/deliveries/"+url.QueryEscape(id)+"/packages", queryParams, nil)
}

// GET /transfers/v2/deliveries/package/{id}/requiredlabtestbatches
// Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersDeliveryPackageRequiredlabtestbatchById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/deliveries/package/"+url.QueryEscape(id)+"/requiredlabtestbatches", queryParams, nil)
}

// GET /transfers/v2/deliveries/{id}/packages/wholesale
// Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersDeliveryPackageWholesaleById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/deliveries/"+url.QueryEscape(id)+"/packages/wholesale", queryParams, nil)
}

// GET /transfers/v2/deliveries/{id}/transporters
// Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersDeliveryTransporterById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/deliveries/"+url.QueryEscape(id)+"/transporters", queryParams, nil)
}

// GET /transfers/v2/deliveries/{id}/transporters/details
// Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersDeliveryTransporterDetailById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/deliveries/"+url.QueryEscape(id)+"/transporters/details", queryParams, nil)
}

// GET /transfers/v2/hub
// Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersHub(
	ctx context.Context,lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastModifiedEnd != "" {
		queryParams["lastModifiedEnd"] = lastModifiedEnd
	}
	if lastModifiedStart != "" {
		queryParams["lastModifiedStart"] = lastModifiedStart
	}
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/hub", queryParams, nil)
}

// GET /transfers/v2/incoming
// Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetIncomingTransfers(
	ctx context.Context,lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastModifiedEnd != "" {
		queryParams["lastModifiedEnd"] = lastModifiedEnd
	}
	if lastModifiedStart != "" {
		queryParams["lastModifiedStart"] = lastModifiedStart
	}
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/incoming", queryParams, nil)
}

// GET /transfers/v2/manifest/{id}/html
// Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetTransfersManifestHtmlById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/transfers/v2/manifest/"+url.QueryEscape(id)+"/html", queryParams, nil)
}

// GET /transfers/v2/manifest/{id}/pdf
// Get Transfer Manifest PDF for a given Transfer Id
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetTransfersManifestPdfById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/transfers/v2/manifest/"+url.QueryEscape(id)+"/pdf", queryParams, nil)
}

// GET /transfers/v2/templates/outgoing/{id}/deliveries
// Retrieves a list of deliveries associated with a specific transfer template.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersOutgoingTemplateDeliveryById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/templates/outgoing/"+url.QueryEscape(id)+"/deliveries", queryParams, nil)
}

// GET /transfers/v2/templates/outgoing/deliveries/{id}/packages
// Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersOutgoingTemplateDeliveryPackageById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/templates/outgoing/deliveries/"+url.QueryEscape(id)+"/packages", queryParams, nil)
}

// GET /transfers/v2/templates/outgoing/deliveries/{id}/transporters
// Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersOutgoingTemplateDeliveryTransporterById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/templates/outgoing/deliveries/"+url.QueryEscape(id)+"/transporters", queryParams, nil)
}

// GET /transfers/v2/templates/outgoing/deliveries/{id}/transporters/details
// Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
func (c *MetrcClient) GetTransfersOutgoingTemplateDeliveryTransporterDetailById(
	ctx context.Context,id string, ) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/transfers/v2/templates/outgoing/deliveries/"+url.QueryEscape(id)+"/transporters/details", queryParams, nil)
}

// GET /transfers/v2/templates/outgoing
// Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersOutgoingTemplates(
	ctx context.Context,lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastModifiedEnd != "" {
		queryParams["lastModifiedEnd"] = lastModifiedEnd
	}
	if lastModifiedStart != "" {
		queryParams["lastModifiedStart"] = lastModifiedStart
	}
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/templates/outgoing", queryParams, nil)
}

// GET /transfers/v2/outgoing
// Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetOutgoingTransfers(
	ctx context.Context,lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastModifiedEnd != "" {
		queryParams["lastModifiedEnd"] = lastModifiedEnd
	}
	if lastModifiedStart != "" {
		queryParams["lastModifiedStart"] = lastModifiedStart
	}
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/outgoing", queryParams, nil)
}

// GET /transfers/v2/rejected
// Retrieves a list of shipments with rejected packages for a Facility.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetRejectedTransfers(
	ctx context.Context,licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/rejected", queryParams, nil)
}

// GET /transfers/v2/{id}/deliveries
// Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersDeliveryById(
	ctx context.Context,id string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/"+url.QueryEscape(id)+"/deliveries", queryParams, nil)
}

// GET /transfers/v2/types
// Retrieves a list of available transfer types for a Facility based on its license number.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetTransfersTypes(
	ctx context.Context,licenseNumber string, pageNumber string, pageSize string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}
	if pageNumber != "" {
		queryParams["pageNumber"] = pageNumber
	}
	if pageSize != "" {
		queryParams["pageSize"] = pageSize
	}

	return c.send(ctx, "GET", "/transfers/v2/types", queryParams, nil)
}

// PUT /transfers/v2/external/incoming
// Updates external incoming shipment plans for a Facility.
// Permissions Required:
// - Manage Transfers
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateTransfersIncomingExternal(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/transfers/v2/external/incoming", queryParams, body)
}

// PUT /transfers/v2/templates/outgoing
// Updates existing transfer templates for a Facility.
// Permissions Required:
// - Manage Transfer Templates
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateTransfersOutgoingTemplates(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/transfers/v2/templates/outgoing", queryParams, body)
}


