package client

import (
	"context"
)


// POST /sandbox/v2/integrator/setup
// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
// Parameters:
// - userKey (optional): Filter by userKey
// - body (request body)
func (c *MetrcClient) CreateSandboxIntegratorSetup(
	ctx context.Context,userKey string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if userKey != "" {
		queryParams["userKey"] = userKey
	}

	return c.send(ctx, "POST", "/sandbox/v2/integrator/setup", queryParams, body)
}


