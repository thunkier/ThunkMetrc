package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type PatientsStatusService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// GET /patients/v2/statuses/{patientLicenseNumber}
// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
// Permissions Required:
// - Lookup Patients
// Parameters:
// - patientLicenseNumber (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *PatientsStatusService) GetPatientsStatusesByPatientLicenseNumber(ctx context.Context, patientLicenseNumber string, licenseNumber string) ([]*models.PatientsStatus, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPatientsStatusesByPatientLicenseNumber(ctx, patientLicenseNumber, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.PatientsStatus
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.PatientsStatus: %w", err)
    }
    return typed, nil
    
}


