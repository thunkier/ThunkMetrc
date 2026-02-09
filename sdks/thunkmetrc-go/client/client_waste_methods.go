package client

import (
	"context"
)


// GET /wastemethods/v2
// Retrieves all available waste methods.
// Parameters:
func (c *MetrcClient) GetWasteMethodsWasteMethods(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/wastemethods/v2", queryParams, nil)
}


