package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type WebhooksService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// DELETE /webhooks/v2/{subscriptionId}
// 
// Parameters:
// - subscriptionId (path param)
func (s *WebhooksService) DeleteWebhooksBySubscriptionId(ctx context.Context, subscriptionId string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteWebhooksBySubscriptionId(ctx, subscriptionId)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /webhooks/v2
// 
// Parameters:
func (s *WebhooksService) GetWebhooks(ctx context.Context, ) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetWebhooks(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /webhooks/v2/disable/{subscriptionId}
// 
// Parameters:
// - subscriptionId (path param)
// - body (request body)
        
func (s *WebhooksService) UpdateWebhooksDisableBySubscriptionId(ctx context.Context, subscriptionId string, body interface{}) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateWebhooksDisableBySubscriptionId(ctx, subscriptionId, body)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /webhooks/v2/enable/{subscriptionId}
// 
// Parameters:
// - subscriptionId (path param)
// - body (request body)
        
func (s *WebhooksService) UpdateWebhooksEnableBySubscriptionId(ctx context.Context, subscriptionId string, body interface{}) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateWebhooksEnableBySubscriptionId(ctx, subscriptionId, body)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /webhooks/v2
// 
// Parameters:
// - body (request body)
        
func (s *WebhooksService) UpdateWebhooks(ctx context.Context, body *models.WebhooksUpdateWebhooksRequest) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateWebhooks(ctx, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.WriteResponse: %w", err)
    }
    return typed, nil
    
}


