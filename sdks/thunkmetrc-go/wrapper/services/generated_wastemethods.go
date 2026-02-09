package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type WasteMethodsService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// GET /wastemethods/v2
// Retrieves all available waste methods.
// Parameters:
func (s *WasteMethodsService) GetWasteMethodsWasteMethods(ctx context.Context, ) ([]*models.WasteMethod, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetWasteMethodsWasteMethods(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.WasteMethod
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.WasteMethod: %w", err)
    }
    return typed, nil
    
}


