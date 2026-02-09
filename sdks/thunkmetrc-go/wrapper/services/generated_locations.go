package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type LocationsService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /locations/v2
// Creates new locations for a specified Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *LocationsService) CreateLocations(ctx context.Context, body []*models.LocationsCreateLocationsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateLocations(ctx, licenseNumber, body)
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
// DELETE /locations/v2/{id}
// Archives a specified Location, identified by its Id, for a Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *LocationsService) DeleteLocationsById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteLocationsById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /locations/v2/active
// Retrieves a list of active locations for a specified Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *LocationsService) GetActiveLocations(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.LocationsLocation], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetActiveLocations(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.LocationsLocation]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.LocationsLocation]: %w", err)
    }
    return typed, nil
    
}
// GET /locations/v2/inactive
// Retrieves a list of inactive locations for a specified Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *LocationsService) GetInactiveLocations(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.LocationsLocation], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInactiveLocations(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.LocationsLocation]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.LocationsLocation]: %w", err)
    }
    return typed, nil
    
}
// GET /locations/v2/{id}
// Retrieves a Location by its Id.
// Permissions Required:
// - Manage Locations
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *LocationsService) GetLocationsById(ctx context.Context, id string, licenseNumber string) (*models.LocationsLocation, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetLocationsById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.LocationsLocation
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.LocationsLocation: %w", err)
    }
    return typed, nil
    
}
// GET /locations/v2/types
// Retrieves a list of active location types for a specified Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *LocationsService) GetLocationsTypes(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.LocationsType], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetLocationsTypes(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.LocationsType]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.LocationsType]: %w", err)
    }
    return typed, nil
    
}
// PUT /locations/v2
// Updates existing locations for a specified Facility.
// Permissions Required:
// - Manage Locations
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *LocationsService) UpdateLocations(ctx context.Context, body []*models.LocationsUpdateLocationsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateLocations(ctx, licenseNumber, body)
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


