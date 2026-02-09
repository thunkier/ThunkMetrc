package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type TagsService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// GET /tags/v2/package/available
// Returns a list of available package tags. NOTE: This is a premium endpoint.
// Permissions Required:
// - Manage Tags
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (s *TagsService) GetTagsAvailablePackage(ctx context.Context, licenseNumber string) ([]*models.Tag, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTagsAvailablePackage(ctx, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.Tag
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.Tag: %w", err)
    }
    return typed, nil
    
}
// GET /tags/v2/plant/available
// Returns a list of available plant tags. NOTE: This is a premium endpoint.
// Permissions Required:
// - Manage Tags
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (s *TagsService) GetTagsAvailablePlant(ctx context.Context, licenseNumber string) ([]*models.Tag, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTagsAvailablePlant(ctx, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.Tag
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.Tag: %w", err)
    }
    return typed, nil
    
}
// GET /tags/v2/staged
// Returns a list of staged tags. NOTE: This is a premium endpoint.
// Permissions Required:
// - Manage Tags
// - RetailId.AllowPackageStaging Key Value enabled
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (s *TagsService) GetStagedTags(ctx context.Context, licenseNumber string) ([]*models.Staged, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetStagedTags(ctx, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.Staged
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.Staged: %w", err)
    }
    return typed, nil
    
}


