package client

import (
	"context"
)


// GET /unitsofmeasure/v2/active
// Retrieves all active units of measure.
// Parameters:
func (c *MetrcClient) GetActiveUnitsOfMeasure(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/unitsofmeasure/v2/active", queryParams, nil)
}

// GET /unitsofmeasure/v2/inactive
// Retrieves all inactive units of measure.
// Parameters:
func (c *MetrcClient) GetInactiveUnitsOfMeasure(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/unitsofmeasure/v2/inactive", queryParams, nil)
}


