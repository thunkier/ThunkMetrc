package client

import (
	"context"
	"net/url"
)


// POST /sales/v2/deliveries
// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateSalesDeliveries(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/sales/v2/deliveries", queryParams, body)
}

// POST /sales/v2/deliveries/retailer/depart
// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateSalesDeliveriesRetailerDepart(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/sales/v2/deliveries/retailer/depart", queryParams, body)
}

// POST /sales/v2/deliveries/retailer/end
// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateSalesDeliveriesRetailerEnd(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/sales/v2/deliveries/retailer/end", queryParams, body)
}

// POST /sales/v2/deliveries/retailer/restock
// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateSalesDeliveriesRetailerRestock(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/sales/v2/deliveries/retailer/restock", queryParams, body)
}

// POST /sales/v2/receipts
// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales (Write)
// - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
// - Industry/Facility Type/Advanced Sales
// - WebApi Sales Read Write State (All or WriteOnly)
// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateSalesReceipts(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/sales/v2/receipts", queryParams, body)
}

// POST /sales/v2/deliveries/retailer
// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateSalesDeliveriesRetailer(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/sales/v2/deliveries/retailer", queryParams, body)
}

// DELETE /sales/v2/deliveries/{id}
// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
// Permissions Required:
// - Manage Sales Delivery
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteSalesDeliveryById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/sales/v2/deliveries/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE /sales/v2/deliveries/retailer/{id}
// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - Manage Retailer Delivery
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteSalesDeliveryRetailerById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/sales/v2/deliveries/retailer/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE /sales/v2/receipts/{id}
// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
// Permissions Required:
// - Manage Sales
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) DeleteSalesReceiptById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "DELETE", "/sales/v2/receipts/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /sales/v2/deliveries/active
// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
// Permissions Required:
// - View Sales Delivery
// - Manage Sales Delivery
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetSalesActiveDeliveries(
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

	return c.send(ctx, "GET", "/sales/v2/deliveries/active", queryParams, nil)
}

// GET /sales/v2/deliveries/retailer/active
// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
// Permissions Required:
// - View Retailer Delivery
// - Manage Retailer Delivery
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetSalesActiveDeliveriesRetailer(
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

	return c.send(ctx, "GET", "/sales/v2/deliveries/retailer/active", queryParams, nil)
}

// GET /sales/v2/receipts/active
// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
// Permissions Required:
// - View Sales
// - Manage Sales
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetSalesActiveReceipts(
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

	return c.send(ctx, "GET", "/sales/v2/receipts/active", queryParams, nil)
}

// GET /sales/v2/counties
// Returns a list of counties available for sales deliveries.
// Parameters:
func (c *MetrcClient) GetSalesCounties(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/sales/v2/counties", queryParams, nil)
}

// GET /sales/v2/customertypes
// Returns a list of customer types.
// Parameters:
func (c *MetrcClient) GetSalesCustomerTypes(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/sales/v2/customertypes", queryParams, nil)
}

// GET /sales/v2/deliveries/returnreasons
// Returns a list of return reasons for sales deliveries based on the provided License Number.
// Permissions Required:
// - Sales Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetSalesDeliveriesReturnReasons(
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

	return c.send(ctx, "GET", "/sales/v2/deliveries/returnreasons", queryParams, nil)
}

// GET /sales/v2/deliveries/retailer/{id}
// Retrieves a retailer delivery record by its ID, with an optional License Number.
// Permissions Required:
// - View Retailer Delivery
// - Manage Retailer Delivery
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetSalesDeliveryRetailerById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/sales/v2/deliveries/retailer/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /sales/v2/deliveries/inactive
// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
// Permissions Required:
// - View Sales Delivery
// - Manage Sales Delivery
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetSalesInactiveDeliveries(
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

	return c.send(ctx, "GET", "/sales/v2/deliveries/inactive", queryParams, nil)
}

// GET /sales/v2/deliveries/retailer/inactive
// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
// Permissions Required:
// - View Retailer Delivery
// - Manage Retailer Delivery
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetSalesInactiveDeliveriesRetailer(
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

	return c.send(ctx, "GET", "/sales/v2/deliveries/retailer/inactive", queryParams, nil)
}

// GET /sales/v2/receipts/inactive
// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
// Permissions Required:
// - View Sales
// - Manage Sales
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (c *MetrcClient) GetSalesInactiveReceipts(
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

	return c.send(ctx, "GET", "/sales/v2/receipts/inactive", queryParams, nil)
}

// GET /sales/v2/patientregistration/locations
// Returns a list of valid Patient registration locations for sales.
// Parameters:
func (c *MetrcClient) GetSalesPatientRegistrationLocations(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/sales/v2/patientregistration/locations", queryParams, nil)
}

// GET /sales/v2/paymenttypes
// Returns a list of available payment types for the specified License Number.
// Permissions Required:
// - View Sales Delivery
// - Manage Sales Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetSalesPaymentTypes(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/sales/v2/paymenttypes", queryParams, nil)
}

// GET /sales/v2/receipts/{id}
// Retrieves a sales receipt by its Id, with an optional License Number.
// Permissions Required:
// - View Sales
// - Manage Sales
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetSalesReceiptById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/sales/v2/receipts/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET /sales/v2/receipts/external/{externalNumber}
// Retrieves a Sales Receipt by its external number, with an optional License Number.
// Permissions Required:
// - View Sales
// - Manage Sales
// Parameters:
// - externalNumber (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetSalesReceiptsExternalByExternalNumber(
	ctx context.Context,externalNumber string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/sales/v2/receipts/external/"+url.QueryEscape(externalNumber)+"", queryParams, nil)
}

// GET /sales/v2/deliveries/{id}
// Retrieves a sales delivery record by its Id, with an optional License Number.
// Permissions Required:
// - View Sales Delivery
// - Manage Sales Delivery
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetSalesDeliveryById(
	ctx context.Context,id string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/sales/v2/deliveries/"+url.QueryEscape(id)+"", queryParams, nil)
}

// PUT /sales/v2/deliveries
// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - Manage Sales Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesDeliveries(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/deliveries", queryParams, body)
}

// PUT /sales/v2/deliveries/complete
// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
// Permissions Required:
// - Manage Sales Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesDeliveriesComplete(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/deliveries/complete", queryParams, body)
}

// PUT /sales/v2/deliveries/hub
// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - Manage Sales Delivery, Manage Sales Delivery Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesDeliveriesHub(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/deliveries/hub", queryParams, body)
}

// PUT /sales/v2/deliveries/hub/accept
// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
// Permissions Required:
// - Manage Sales Delivery Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesDeliveriesHubAccept(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/deliveries/hub/accept", queryParams, body)
}

// PUT /sales/v2/deliveries/hub/depart
// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
// Permissions Required:
// - Manage Sales Delivery Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesDeliveriesHubDepart(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/deliveries/hub/depart", queryParams, body)
}

// PUT /sales/v2/deliveries/hub/verifyID
// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
// Permissions Required:
// - Manage Sales Delivery Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesDeliveriesHubVerifyID(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/deliveries/hub/verifyID", queryParams, body)
}

// PUT /sales/v2/deliveries/retailer
// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesDeliveriesRetailer(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/deliveries/retailer", queryParams, body)
}

// PUT /sales/v2/receipts
// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - Manage Sales
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesReceipts(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/receipts", queryParams, body)
}

// PUT /sales/v2/receipts/finalize
// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
// Permissions Required:
// - Manage Sales
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesReceiptsFinalize(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/receipts/finalize", queryParams, body)
}

// PUT /sales/v2/receipts/unfinalize
// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
// Permissions Required:
// - Manage Sales
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) UpdateSalesReceiptsUnfinalize(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "PUT", "/sales/v2/receipts/unfinalize", queryParams, body)
}


