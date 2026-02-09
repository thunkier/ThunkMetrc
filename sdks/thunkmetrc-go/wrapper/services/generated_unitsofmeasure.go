package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type UnitsOfMeasureService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// GET /unitsofmeasure/v2/active
// Retrieves all active units of measure.
// Parameters:
func (s *UnitsOfMeasureService) GetActiveUnitsOfMeasure(ctx context.Context, ) (*models.PaginatedResponse[*models.UnitOfMeasure], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetActiveUnitsOfMeasure(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.UnitOfMeasure]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.UnitOfMeasure]: %w", err)
    }
    return typed, nil
    
}
// GET /unitsofmeasure/v2/inactive
// Retrieves all inactive units of measure.
// Parameters:
func (s *UnitsOfMeasureService) GetInactiveUnitsOfMeasure(ctx context.Context, ) (*models.PaginatedResponse[*models.UnitOfMeasure], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInactiveUnitsOfMeasure(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.UnitOfMeasure]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.UnitOfMeasure]: %w", err)
    }
    return typed, nil
    
}


