package client

import (
	"context"
	"net/url"
)


// POST /retailid/v2/associate
// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
// - WebApi Retail ID Read Write State (All or WriteOnly)
// - Industry/View Packages
// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateRetailIdAssociate(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/retailid/v2/associate", queryParams, body)
}

// POST /retailid/v2/generate
// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
// - WebApi Retail ID Read Write State (All or WriteOnly)
// - Industry/View Packages
// - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateRetailIdGenerate(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/retailid/v2/generate", queryParams, body)
}

// POST /retailid/v2/merge
// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
// - WebApi Retail ID Read Write State (All or WriteOnly)
// - Key Value Settings/Retail ID Merge Packages Enabled
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateRetailIdMerge(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/retailid/v2/merge", queryParams, body)
}

// POST /retailid/v2/packages/info
// Retrieves Package information for given list of Package labels.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
// - WebApi Retail ID Read Write State (All or WriteOnly)
// - Industry/View Packages
// - Admin/Employees/Packages Page/Product Labels(Manage)
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
func (c *MetrcClient) CreateRetailIdPackagesInfo(
	ctx context.Context,licenseNumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "POST", "/retailid/v2/packages/info", queryParams, body)
}

// GET /retailid/v2/allotment
// Retrieves the available Retail Item ID quota for a facility.
// Permissions Required:
// - Download Product Labels
// - Manage Product Labels
// - Manage Tag Orders
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetRetailIdAllotment(
	ctx context.Context,licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/retailid/v2/allotment", queryParams, nil)
}

// GET /retailid/v2/receive/{label}
// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Manage RetailId
// - WebApi Retail ID Read Write State (All or ReadOnly)
// - Industry/View Packages
// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
// Parameters:
// - label (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetRetailIdReceiveByLabel(
	ctx context.Context,label string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/retailid/v2/receive/"+url.QueryEscape(label)+"", queryParams, nil)
}

// GET /retailid/v2/receive/qr/{shortCode}
// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Manage RetailId
// - WebApi Retail ID Read Write State (All or ReadOnly)
// - Industry/View Packages
// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
// Parameters:
// - shortCode (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetRetailIdReceiveQrByShortCode(
	ctx context.Context,shortCode string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/retailid/v2/receive/qr/"+url.QueryEscape(shortCode)+"", queryParams, nil)
}


