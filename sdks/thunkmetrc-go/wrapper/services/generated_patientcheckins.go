package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type PatientCheckInsService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /patient-checkins/v2
// Records patient check-ins for a specified Facility.
// Permissions Required:
// - ManagePatientsCheckIns
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PatientCheckInsService) CreatePatientCheckIns(ctx context.Context, body []*models.PatientCheckInsCreatePatientCheckInsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePatientCheckIns(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.WriteResponse: %w", err)
    }
    return typed, nil
    
}
// DELETE /patient-checkins/v2/{id}
// Archives a Patient Check-In, identified by its Id, for a specified Facility.
// Permissions Required:
// - ManagePatientsCheckIns
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *PatientCheckInsService) DeletePatientCheckInsById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeletePatientCheckInsById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /patient-checkins/v2/locations
// Retrieves a list of Patient Check-In locations.
// Parameters:
func (s *PatientCheckInsService) GetPatientCheckInsLocations(ctx context.Context, ) ([]*models.PatientCheckInsLocation, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPatientCheckInsLocations(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.PatientCheckInsLocation
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.PatientCheckInsLocation: %w", err)
    }
    return typed, nil
    
}
// GET /patient-checkins/v2
// Retrieves a list of patient check-ins for a specified Facility.
// Permissions Required:
// - ManagePatientsCheckIns
// Parameters:
// - checkinDateEnd (optional): Filter by checkinDateEnd
// - checkinDateStart (optional): Filter by checkinDateStart
// - licenseNumber (optional): Filter by licenseNumber
func (s *PatientCheckInsService) GetPatientCheckIns(ctx context.Context, checkinDateEnd string, checkinDateStart string, licenseNumber string) (*models.PaginatedResponse[*models.PatientCheckIn], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPatientCheckIns(ctx, checkinDateEnd, checkinDateStart, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PatientCheckIn]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PatientCheckIn]: %w", err)
    }
    return typed, nil
    
}
// PUT /patient-checkins/v2
// Updates patient check-ins for a specified Facility.
// Permissions Required:
// - ManagePatientsCheckIns
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PatientCheckInsService) UpdatePatientCheckIns(ctx context.Context, body []*models.PatientCheckInsUpdatePatientCheckInsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePatientCheckIns(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.WriteResponse: %w", err)
    }
    return typed, nil
    
}


