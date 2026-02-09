package client

import (
	"context"
	"net/url"
)


// GET /patients/v2/statuses/{patientLicenseNumber}
// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
// Permissions Required:
// - Lookup Patients
// Parameters:
// - patientLicenseNumber (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (c *MetrcClient) GetPatientsStatusesByPatientLicenseNumber(
	ctx context.Context,patientLicenseNumber string, licenseNumber string, ) (interface{}, error) {
	queryParams := make(map[string]string)
	if licenseNumber != "" {
		queryParams["licenseNumber"] = licenseNumber
	}

	return c.send(ctx, "GET", "/patients/v2/statuses/"+url.QueryEscape(patientLicenseNumber)+"", queryParams, nil)
}


