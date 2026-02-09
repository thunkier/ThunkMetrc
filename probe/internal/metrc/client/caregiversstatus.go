package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// GetCaregiversStatusByCaregiverLicenseNumber (CaregiversStatus)
// GET {{baseUrl}}/caregivers/v2/status/{caregiverLicenseNumber}
//   licenseNumber (required): The License Number of the Facility under which to get the Caregiver status.
func (f *Fetcher) GetCaregiversStatusByCaregiverLicenseNumber(caregiverLicenseNumber string, licenseNumber string) ([]models.CaregiversStatus, error) {
	url := "{{baseUrl}}/caregivers/v2/status/{caregiverLicenseNumber}"
	url = strings.ReplaceAll(url, "{"+"caregiverLicenseNumber"+"}", caregiverLicenseNumber)
	
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
	return fetchList[models.CaregiversStatus](f, "CaregiversStatus", "GetCaregiversStatusByCaregiverLicenseNumber", "GET", url, nil)
}
