package client

import (
	"context"
	"net/url"
)


// DELETE /webhooks/v2/{subscriptionId}
// 
// Parameters:
// - subscriptionId (path param)
func (c *MetrcClient) DeleteWebhooksBySubscriptionId(
	ctx context.Context,subscriptionId string, ) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "DELETE", "/webhooks/v2/"+url.QueryEscape(subscriptionId)+"", queryParams, nil)
}

// GET /webhooks/v2
// 
// Parameters:
func (c *MetrcClient) GetWebhooks(
	ctx context.Context,) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "GET", "/webhooks/v2", queryParams, nil)
}

// PUT /webhooks/v2/disable/{subscriptionId}
// 
// Parameters:
// - subscriptionId (path param)
// - body (request body)
func (c *MetrcClient) UpdateWebhooksDisableBySubscriptionId(
	ctx context.Context,subscriptionId string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "PUT", "/webhooks/v2/disable/"+url.QueryEscape(subscriptionId)+"", queryParams, body)
}

// PUT /webhooks/v2/enable/{subscriptionId}
// 
// Parameters:
// - subscriptionId (path param)
// - body (request body)
func (c *MetrcClient) UpdateWebhooksEnableBySubscriptionId(
	ctx context.Context,subscriptionId string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "PUT", "/webhooks/v2/enable/"+url.QueryEscape(subscriptionId)+"", queryParams, body)
}

// PUT /webhooks/v2
// 
// Parameters:
// - body (request body)
func (c *MetrcClient) UpdateWebhooks(
	ctx context.Context,body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)

	return c.send(ctx, "PUT", "/webhooks/v2", queryParams, body)
}


