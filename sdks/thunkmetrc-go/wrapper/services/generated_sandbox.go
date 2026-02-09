package services

import (
	"context"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
)

type SandboxService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /sandbox/v2/integrator/setup
// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
// Parameters:
// - userKey (optional): Filter by userKey
// - body (request body)
        
func (s *SandboxService) CreateSandboxIntegratorSetup(ctx context.Context, body interface{}, userKey string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateSandboxIntegratorSetup(ctx, userKey, body)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}


