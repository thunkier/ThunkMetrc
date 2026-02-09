package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type AdditivesTemplatesService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /additivestemplates/v2
// Creates new additive templates for a specified Facility.
// Permissions Required:
// - Manage Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *AdditivesTemplatesService) CreateAdditivesTemplates(ctx context.Context, body []*models.AdditivesTemplatesCreateAdditivesTemplatesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateAdditivesTemplates(ctx, licenseNumber, body)
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
// GET /additivestemplates/v2/active
// Retrieves a list of active additive templates for a specified Facility.
// Permissions Required:
// - Manage Additives
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *AdditivesTemplatesService) GetActiveAdditivesTemplates(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.AdditivesTemplate], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetActiveAdditivesTemplates(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.AdditivesTemplate]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.AdditivesTemplate]: %w", err)
    }
    return typed, nil
    
}
// GET /additivestemplates/v2/{id}
// Retrieves an Additive Template by its Id.
// Permissions Required:
// - Manage Additives
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *AdditivesTemplatesService) GetAdditivesTemplatesById(ctx context.Context, id string, licenseNumber string) (*models.AdditivesTemplate, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetAdditivesTemplatesById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.AdditivesTemplate
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.AdditivesTemplate: %w", err)
    }
    return typed, nil
    
}
// GET /additivestemplates/v2/inactive
// Retrieves a list of inactive additive templates for a specified Facility.
// Permissions Required:
// - Manage Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *AdditivesTemplatesService) GetInactiveAdditivesTemplates(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.AdditivesTemplate], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInactiveAdditivesTemplates(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.AdditivesTemplate]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.AdditivesTemplate]: %w", err)
    }
    return typed, nil
    
}
// PUT /additivestemplates/v2
// Updates existing additive templates for a specified Facility.
// Permissions Required:
// - Manage Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *AdditivesTemplatesService) UpdateAdditivesTemplates(ctx context.Context, body []*models.AdditivesTemplatesUpdateAdditivesTemplatesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateAdditivesTemplates(ctx, licenseNumber, body)
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


