package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type StrainsService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /strains/v2
// Creates new strain records for a specified Facility.
// Permissions Required:
// - Manage Strains
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *StrainsService) CreateStrains(ctx context.Context, body []*models.StrainsCreateStrainsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateStrains(ctx, licenseNumber, body)
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
// DELETE /strains/v2/{id}
// Archives an existing strain record for a Facility
// Permissions Required:
// - Manage Strains
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *StrainsService) DeleteStrainsById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteStrainsById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /strains/v2/active
// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Strains
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *StrainsService) GetActiveStrains(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Strain], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetActiveStrains(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Strain]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Strain]: %w", err)
    }
    return typed, nil
    
}
// GET /strains/v2/inactive
// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Strains
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *StrainsService) GetInactiveStrains(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Strain], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInactiveStrains(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Strain]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Strain]: %w", err)
    }
    return typed, nil
    
}
// GET /strains/v2/{id}
// Retrieves a Strain record by its Id, with an optional license number.
// Permissions Required:
// - Manage Strains
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *StrainsService) GetStrainsById(ctx context.Context, id string, licenseNumber string) (*models.Strain, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetStrainsById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.Strain
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.Strain: %w", err)
    }
    return typed, nil
    
}
// PUT /strains/v2
// Updates existing strain records for a specified Facility.
// Permissions Required:
// - Manage Strains
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *StrainsService) UpdateStrains(ctx context.Context, body []*models.StrainsUpdateStrainsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateStrains(ctx, licenseNumber, body)
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


