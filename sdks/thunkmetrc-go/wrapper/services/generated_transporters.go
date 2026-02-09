package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type TransportersService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /transporters/v2/drivers
// Creates new driver records for a Facility.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransportersService) CreateTransportersDrivers(ctx context.Context, body []*models.TransportersCreateDriversRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateTransportersDrivers(ctx, licenseNumber, body)
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
// POST /transporters/v2/vehicles
// Creates new vehicle records for a Facility.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransportersService) CreateTransportersVehicles(ctx context.Context, body []*models.TransportersCreateVehiclesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateTransportersVehicles(ctx, licenseNumber, body)
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
// DELETE /transporters/v2/drivers/{id}
// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *TransportersService) DeleteTransportersDriverById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteTransportersDriverById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// DELETE /transporters/v2/vehicles/{id}
// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *TransportersService) DeleteTransportersVehicleById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteTransportersVehicleById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /transporters/v2/drivers/{id}
// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
// Permissions Required:
// - Transporters
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *TransportersService) GetTransportersDriverById(ctx context.Context, id string, licenseNumber string) (*models.TransportersDriver, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransportersDriverById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.TransportersDriver
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.TransportersDriver: %w", err)
    }
    return typed, nil
    
}
// GET /transporters/v2/drivers
// Retrieves a list of drivers for a Facility.
// Permissions Required:
// - Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransportersService) GetTransportersDrivers(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.TransportersDriver], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransportersDrivers(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.TransportersDriver]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.TransportersDriver]: %w", err)
    }
    return typed, nil
    
}
// GET /transporters/v2/vehicles/{id}
// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
// Permissions Required:
// - Transporters
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *TransportersService) GetTransportersVehicleById(ctx context.Context, id string, licenseNumber string) (*models.TransportersVehicle, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransportersVehicleById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.TransportersVehicle
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.TransportersVehicle: %w", err)
    }
    return typed, nil
    
}
// GET /transporters/v2/vehicles
// Retrieves a list of vehicles for a Facility.
// Permissions Required:
// - Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransportersService) GetTransportersVehicles(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.TransportersVehicle], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransportersVehicles(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.TransportersVehicle]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.TransportersVehicle]: %w", err)
    }
    return typed, nil
    
}
// PUT /transporters/v2/drivers
// Updates existing driver records for a Facility.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransportersService) UpdateTransportersDrivers(ctx context.Context, body []*models.TransportersUpdateDriversRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateTransportersDrivers(ctx, licenseNumber, body)
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
// PUT /transporters/v2/vehicles
// Updates existing vehicle records for a facility.
// Permissions Required:
// - Manage Transporters
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransportersService) UpdateTransportersVehicles(ctx context.Context, body []*models.TransportersUpdateVehiclesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateTransportersVehicles(ctx, licenseNumber, body)
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


