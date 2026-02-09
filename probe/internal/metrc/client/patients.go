package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// CreatePatients (Patients)
// POST {{baseUrl}}/patients/v2
//   licenseNumber (required): The License Number of the Facility for which to record the list of patients.
func (f *Fetcher) CreatePatients(body []models.PatientsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/patients/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Patients", "CreatePatients", "POST", url, body)
}

// DeletePatientsById (Patients)
// DELETE {{baseUrl}}/patients/v2/{id}
//   licenseNumber (required): The License Number of the Facility for which to delete the Patient.
func (f *Fetcher) DeletePatientsById(id string, licenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/patients/v2/{id}"
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
	return fetchOne[map[string]any](f, "Patients", "DeletePatientsById", "DELETE", url, nil)
}

// GetActivePatients (Patients)
// GET {{baseUrl}}/patients/v2/active
//   licenseNumber (required): The License Number of the Facility for which to return the list of active patients.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetActivePatients(licenseNumber string, pageNumber string, pageSize string) ([]models.Patient, error) {
	url := "{{baseUrl}}/patients/v2/active"
	
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
	return fetchList[models.Patient](f, "Patients", "GetActivePatients", "GET", url, nil)
}

// GetPatientsById (Patients)
// GET {{baseUrl}}/patients/v2/{id}
//   licenseNumber (required): If specified, the Patient will be validated against the specified License Number. If not specified, the Patient will be validated against all of the User's current Facilities. Please note that if the Patient is not valid for the specified License Number, a 401 Unauthorized status will be returned.
func (f *Fetcher) GetPatientsById(id string, licenseNumber string) (models.Patient, error) {
	url := "{{baseUrl}}/patients/v2/{id}"
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
	return fetchOne[models.Patient](f, "Patients", "GetPatientsById", "GET", url, nil)
}

// UpdatePatients (Patients)
// PUT {{baseUrl}}/patients/v2
//   licenseNumber (required): The License Number of the Facility for which to update the list of patients.
func (f *Fetcher) UpdatePatients(body []models.PatientsRequest, licenseNumber string) (models.WriteResponse, error) {
	url := "{{baseUrl}}/patients/v2"
	
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
	return fetchOne[models.WriteResponse](f, "Patients", "UpdatePatients", "PUT", url, body)
}
