package client

import (
	"strings"
	"fmt"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// GetEmployees (Employees)
// GET {{baseUrl}}/employees/v2
//   licenseNumber (required): The License Number of the Facility under which to get the employees.
//   pageNumber (optional): The number of the data page from which to return data.
//   pageSize (optional): The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.
func (f *Fetcher) GetEmployees(licenseNumber string, pageNumber string, pageSize string) ([]models.Employee, error) {
	url := "{{baseUrl}}/employees/v2"
	
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
	return fetchList[models.Employee](f, "Employees", "GetEmployees", "GET", url, nil)
}

// GetPermissions (Employees)
// GET {{baseUrl}}/employees/v2/permissions
//   licenseNumber (required): The License Number of the Facility under which to get the Employee.
//   employeeLicenseNumber (required): The License Number of the Employee under which to get the Permission.
func (f *Fetcher) GetPermissions(licenseNumber string, employeeLicenseNumber string) (map[string]any, error) {
	url := "{{baseUrl}}/employees/v2/permissions"
	
	queryParams := make([]string, 0)
	if licenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("licenseNumber=%s", licenseNumber))
	}
	if employeeLicenseNumber != "" {
		queryParams = append(queryParams, fmt.Sprintf("employeeLicenseNumber=%s", employeeLicenseNumber))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[map[string]any](f, "Employees", "GetPermissions", "GET", url, nil)
}
