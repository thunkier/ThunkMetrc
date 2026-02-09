package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type CaregiversStatusService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// GET /caregivers/v2/status/{caregiverLicenseNumber}
// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
// Permissions Required:
// - Lookup Caregivers
// Parameters:
// - caregiverLicenseNumber (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *CaregiversStatusService) GetCaregiversStatusByCaregiverLicenseNumber(ctx context.Context, caregiverLicenseNumber string, licenseNumber string) ([]*models.CaregiversStatus, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetCaregiversStatusByCaregiverLicenseNumber(ctx, caregiverLicenseNumber, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.CaregiversStatus
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.CaregiversStatus: %w", err)
    }
    return typed, nil
    
}


