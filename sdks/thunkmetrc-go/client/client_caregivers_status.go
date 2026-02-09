package client

import (
	"context"
	"net/url"
)


// GET /caregivers/v2/status/{caregiverLicenseNumber}
// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
// Permissions Required:
// - Lookup Caregivers
// Parameters:
// - caregiverLicenseNumber (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetCaregiversStatusByCaregiverLicenseNumber(
	ctx context.Context,caregiverLicenseNumber string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/caregivers/v2/status/"+url.QueryEscape(caregiverLicenseNumber)+"", queryParams, nil)
}


