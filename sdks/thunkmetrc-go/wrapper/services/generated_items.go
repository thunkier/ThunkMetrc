package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type ItemsService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /items/v2/brand
// Creates one or more new item brands for the specified Facility identified by the License Number.
// Permissions Required:
// - Manage Items
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ItemsService) CreateItemsBrand(ctx context.Context, body []*models.ItemsCreateBrandRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateItemsBrand(ctx, licenseNumber, body)
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
// POST /items/v2/file
// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
// Permissions Required:
// - Manage Items
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ItemsService) CreateItemsFile(ctx context.Context, body []*models.ItemsCreateFileRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateItemsFile(ctx, licenseNumber, body)
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
// POST /items/v2
// Permissions Required:
// - Manage Items
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ItemsService) CreateItems(ctx context.Context, body []*models.ItemsCreateItemsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateItems(ctx, licenseNumber, body)
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
// POST /items/v2/photo
// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
// Permissions Required:
// - Manage Items
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ItemsService) CreateItemsPhoto(ctx context.Context, body []*models.ItemsCreatePhotoRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateItemsPhoto(ctx, licenseNumber, body)
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
// DELETE /items/v2/brand/{id}
// Archives the specified Item Brand by Id for the given Facility License Number.
// Permissions Required:
// - Manage Items
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *ItemsService) DeleteItemsBrandById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteItemsBrandById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// DELETE /items/v2/{id}
// Archives the specified Product by Id for the given Facility License Number.
// Permissions Required:
// - Manage Items
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *ItemsService) DeleteItemsById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteItemsById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /items/v2/active
// Returns a list of active items for the specified Facility.
// Permissions Required:
// - Manage Items
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *ItemsService) GetActiveItems(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Item], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetActiveItems(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Item]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Item]: %w", err)
    }
    return typed, nil
    
}
// GET /items/v2/brands
// Retrieves a list of active item brands for the specified Facility.
// Permissions Required:
// - Manage Items
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *ItemsService) GetItemsBrands(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Brand], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetItemsBrands(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Brand]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Brand]: %w", err)
    }
    return typed, nil
    
}
// GET /items/v2/categories
// Retrieves a list of item categories.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *ItemsService) GetItemsCategories(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Category], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetItemsCategories(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Category]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Category]: %w", err)
    }
    return typed, nil
    
}
// GET /items/v2/file/{id}
// Retrieves a file by its Id for the specified Facility.
// Permissions Required:
// - Manage Items
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *ItemsService) GetItemsFileById(ctx context.Context, id string, licenseNumber string) (*models.File, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetItemsFileById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.File
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.File: %w", err)
    }
    return typed, nil
    
}
// GET /items/v2/inactive
// Retrieves a list of inactive items for the specified Facility.
// Permissions Required:
// - Manage Items
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *ItemsService) GetInactiveItems(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Item], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInactiveItems(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Item]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Item]: %w", err)
    }
    return typed, nil
    
}
// GET /items/v2/{id}
// Retrieves detailed information about a specific Item by Id.
// Permissions Required:
// - Manage Items
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *ItemsService) GetItemsById(ctx context.Context, id string, licenseNumber string) (*models.Item, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetItemsById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.Item
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.Item: %w", err)
    }
    return typed, nil
    
}
// GET /items/v2/photo/{id}
// Retrieves an image by its Id for the specified Facility.
// Permissions Required:
// - Manage Items
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *ItemsService) GetItemsPhotoById(ctx context.Context, id string, licenseNumber string) (*models.Photo, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetItemsPhotoById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.Photo
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.Photo: %w", err)
    }
    return typed, nil
    
}
// PUT /items/v2/brand
// Updates one or more existing item brands for the specified Facility.
// Permissions Required:
// - Manage Items
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ItemsService) UpdateItemsBrand(ctx context.Context, body []*models.ItemsUpdateBrandRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateItemsBrand(ctx, licenseNumber, body)
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
// PUT /items/v2
// Updates one or more existing products for the specified Facility.
// Permissions Required:
// - Manage Items
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ItemsService) UpdateItems(ctx context.Context, body []*models.ItemsUpdateItemsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateItems(ctx, licenseNumber, body)
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


