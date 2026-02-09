package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type EmployeesService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// GET /employees/v2
// Retrieves a list of employees for a specified Facility.
// Permissions Required:
// - Manage Employees
// - View Employees
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *EmployeesService) GetEmployees(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Employee], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetEmployees(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Employee]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Employee]: %w", err)
    }
    return typed, nil
    
}
// GET /employees/v2/permissions
// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
// Permissions Required:
// - Manage Employees
// Parameters:
// - employeeLicenseNumber (optional): Filter by employeeLicenseNumber
// - licenseNumber (optional): Filter by licenseNumber
func (s *EmployeesService) GetEmployeesPermissions(ctx context.Context, employeeLicenseNumber string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetEmployeesPermissions(ctx, employeeLicenseNumber, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}


