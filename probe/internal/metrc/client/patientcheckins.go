package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreatePatientCheckIns (PatientCheckIns)
// POST {{baseUrl}}/patient-checkins/v2
//   licenseNumber (required): The License Number of the Facility for which to record the list of Patient check-ins.
func (f *Fetcher) CreatePatientCheckIns(body []models.PatientCheckInsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/patient-checkins/v2"
	
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
	return fetchOne[models.WriteResponse](f, "PatientCheckIns", "CreatePatientCheckIns", "POST", url, body)
}

// DeletePatientCheckInsById (PatientCheckIns)
// DELETE {{baseUrl}}/patient-checkins/v2/{id}
//   licenseNumber (required): The License Number of the Facility for which to archive Patient check-in.
func (f *Fetcher) DeletePatientCheckInsById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/patient-checkins/v2/{id}"
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
	return fetchOne[map[string]any](f, "PatientCheckIns", "DeletePatientCheckInsById", "DELETE", url, nil)
}

// GetLocations (PatientCheckIns)
// GET {{baseUrl}}/patient-checkins/v2/locations
func (f *Fetcher) GetLocations() ([]models.PatientCheckInsLocation, error) {
	url := "{{baseUrl}}/patient-checkins/v2/locations"
	return fetchList[models.PatientCheckInsLocation](f, "PatientCheckIns", "GetLocations", "GET", url, nil)
}

// GetPatientCheckIns (PatientCheckIns)
// GET {{baseUrl}}/patient-checkins/v2
//   licenseNumber (required): The License Number of the Facility for which to return the list of Patient check-ins.
//   checkinDateStart (optional): The last modified start timestamp
//   checkinDateEnd (optional): The last modified end timestamp
func (f *Fetcher) GetPatientCheckIns(licenseNumber string, checkinDateStart string, checkinDateEnd string) ([]models.PatientCheckIn, error) {
	url := "{{baseUrl}}/patient-checkins/v2"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}
	if checkinDateStart != "" {
		queryParams = append(queryParams, fmt.Sprintf("checkinDateStart=%s", checkinDateStart))
	}
	if checkinDateEnd != "" {
		queryParams = append(queryParams, fmt.Sprintf("checkinDateEnd=%s", checkinDateEnd))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchList[models.PatientCheckIn](f, "PatientCheckIns", "GetPatientCheckIns", "GET", url, nil)
}

// UpdatePatientCheckIns (PatientCheckIns)
// PUT {{baseUrl}}/patient-checkins/v2
//   licenseNumber (required): The License Number of the Facility for which to update the list of Patient check-ins.
func (f *Fetcher) UpdatePatientCheckIns(body []models.PatientCheckInsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/patient-checkins/v2"
	
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
	return fetchOne[models.WriteResponse](f, "PatientCheckIns", "UpdatePatientCheckIns", "PUT", url, body)
}
