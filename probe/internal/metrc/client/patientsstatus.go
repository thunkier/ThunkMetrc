package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// GetPatientsStatusesByPatientLicenseNumber (PatientsStatus)
// GET {{baseUrl}}/patients/v2/statuses/{patientLicenseNumber}
//   licenseNumber (required): The License Number of the Facility under which to get the Patient Status.
func (f *Fetcher) GetPatientsStatusesByPatientLicenseNumber(patientLicenseNumber string, licenseNumber string) ([]models.PatientsStatus, error) {
	url := "{{baseUrl}}/patients/v2/statuses/{patientLicenseNumber}"
	url = strings.ReplaceAll(url, "{"+"patientLicenseNumber"+"}", patientLicenseNumber)
	
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
	return fetchList[models.PatientsStatus](f, "PatientsStatus", "GetPatientsStatusesByPatientLicenseNumber", "GET", url, nil)
}
