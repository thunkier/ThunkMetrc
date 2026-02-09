package client

import (
	"context"
)


// GET /facilities/v2
// This endpoint provides a list of facilities for which the authenticated user has access.
// Parameters:
func (c *MetrcClient) GetFacilities(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/facilities/v2", queryParams, nil)
}


