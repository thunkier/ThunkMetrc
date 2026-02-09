package client

import (
	"strings"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

// DeleteWebhooksBySubscriptionId (Webhooks)
// DELETE {{baseUrl}}/webhooks/v2/{subscriptionId}
func (f *Fetcher) DeleteWebhooksBySubscriptionId(subscriptionId string) (map[string]any, error) {
	url := "{{baseUrl}}/webhooks/v2/{subscriptionId}"
	url = strings.ReplaceAll(url, "{"+"subscriptionId"+"}", subscriptionId)
	return fetchOne[map[string]any](f, "Webhooks", "DeleteWebhooksBySubscriptionId", "DELETE", url, nil)
}

// GetWebhooks (Webhooks)
// GET {{baseUrl}}/webhooks/v2
func (f *Fetcher) GetWebhooks() (map[string]any, error) {
	url := "{{baseUrl}}/webhooks/v2"
	return fetchOne[map[string]any](f, "Webhooks", "GetWebhooks", "GET", url, nil)
}

// UpdateDisableBySubscriptionId (Webhooks)
// PUT {{baseUrl}}/webhooks/v2/disable/{subscriptionId}
func (f *Fetcher) UpdateDisableBySubscriptionId(subscriptionId string) (map[string]any, error) {
	url := "{{baseUrl}}/webhooks/v2/disable/{subscriptionId}"
	url = strings.ReplaceAll(url, "{"+"subscriptionId"+"}", subscriptionId)
	return fetchOne[map[string]any](f, "Webhooks", "UpdateDisableBySubscriptionId", "PUT", url, nil)
}

// UpdateEnableBySubscriptionId (Webhooks)
// PUT {{baseUrl}}/webhooks/v2/enable/{subscriptionId}
func (f *Fetcher) UpdateEnableBySubscriptionId(subscriptionId string) (map[string]any, error) {
	url := "{{baseUrl}}/webhooks/v2/enable/{subscriptionId}"
	url = strings.ReplaceAll(url, "{"+"subscriptionId"+"}", subscriptionId)
	return fetchOne[map[string]any](f, "Webhooks", "UpdateEnableBySubscriptionId", "PUT", url, nil)
}

// UpdateWebhooks (Webhooks)
// PUT {{baseUrl}}/webhooks/v2
func (f *Fetcher) UpdateWebhooks(body models.WebhooksRequest) (models.WriteResponse, error) {
	url := "{{baseUrl}}/webhooks/v2"
	return fetchOne[models.WriteResponse](f, "Webhooks", "UpdateWebhooks", "PUT", url, body)
}
