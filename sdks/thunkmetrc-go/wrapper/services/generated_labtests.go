package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type LabTestsService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /labtests/v2/record
// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *LabTestsService) CreateLabTestsRecord(ctx context.Context, body []*models.LabTestsCreateRecordRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateLabTestsRecord(ctx, licenseNumber, body)
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
// GET /labtests/v2/batches
// Retrieves a list of Lab Test batches.
// Parameters:
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *LabTestsService) GetLabTestsBatches(ctx context.Context, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Batch], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetLabTestsBatches(ctx, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Batch]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Batch]: %w", err)
    }
    return typed, nil
    
}
// GET /labtests/v2/labtestdocument/{id}
// Retrieves a specific Lab Test result document by its Id for a given Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *LabTestsService) GetLabTestsLabTestDocumentById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetLabTestsLabTestDocumentById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /labtests/v2/types
// Returns a list of Lab Test types.
// Parameters:
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *LabTestsService) GetLabTestsTypes(ctx context.Context, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.LabTestsType], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetLabTestsTypes(ctx, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.LabTestsType]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.LabTestsType]: %w", err)
    }
    return typed, nil
    
}
// GET /labtests/v2/results
// Retrieves Lab Test results for a specified Package.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - packageId (optional): Filter by packageId
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *LabTestsService) GetLabTestsResults(ctx context.Context, licenseNumber string, packageId string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Result], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetLabTestsResults(ctx, licenseNumber, packageId, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Result]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Result]: %w", err)
    }
    return typed, nil
    
}
// GET /labtests/v2/states
// Returns a list of all lab testing states.
// Parameters:
func (s *LabTestsService) GetLabTestsStates(ctx context.Context, ) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetLabTestsStates(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /labtests/v2/labtestdocument
// Updates one or more documents for previously submitted lab tests.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *LabTestsService) UpdateLabTestsLabTestDocument(ctx context.Context, body []*models.LabTestsUpdateLabTestDocumentRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateLabTestsLabTestDocument(ctx, licenseNumber, body)
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
// PUT /labtests/v2/results/release
// Releases Lab Test results for one or more packages.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *LabTestsService) UpdateLabTestsResultsRelease(ctx context.Context, body []*models.LabTestsUpdateResultsReleaseRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateLabTestsResultsRelease(ctx, licenseNumber, body)
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


