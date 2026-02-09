package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type FacilitiesService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// GET /facilities/v2
// This endpoint provides a list of facilities for which the authenticated user has access.
// Parameters:
func (s *FacilitiesService) GetFacilities(ctx context.Context, ) ([]*models.Facility, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetFacilities(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.Facility
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.Facility: %w", err)
    }
    return typed, nil
    
}


