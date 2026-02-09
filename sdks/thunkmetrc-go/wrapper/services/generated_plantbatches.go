package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type PlantBatchesService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /plantbatches/v2/adjust
// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreateAdjustPlantBatches(ctx context.Context, body []*models.PlantBatchesCreateAdjustPlantBatchesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateAdjustPlantBatches(ctx, licenseNumber, body)
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
// POST /plantbatches/v2/growthphase
// Updates the growth phase of plants at a specified Facility based on tracking information.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreatePlantBatchesGrowthPhase(ctx context.Context, body []*models.PlantBatchesCreateGrowthPhaseRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePlantBatchesGrowthPhase(ctx, licenseNumber, body)
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
// POST /plantbatches/v2/packages/frommotherplant
// Creates packages from mother plants at the specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreatePlantBatchesPackagesFromMotherPlant(ctx context.Context, body []*models.PlantBatchesCreatePackagesFromMotherPlantRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePlantBatchesPackagesFromMotherPlant(ctx, licenseNumber, body)
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
// POST /plantbatches/v2/additives
// Records Additive usage details for plant batches at a specific Facility.
// Permissions Required:
// - Manage Plants Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreatePlantBatchesAdditives(ctx context.Context, body []*models.PlantBatchesCreatePlantBatchesAdditivesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePlantBatchesAdditives(ctx, licenseNumber, body)
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
// POST /plantbatches/v2/additives/usingtemplate
// Records Additive usage for plant batches at a Facility using predefined additive templates.
// Permissions Required:
// - Manage Plants Additives
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreatePlantBatchesAdditivesUsingTemplate(ctx context.Context, body []*models.PlantBatchesCreatePlantBatchesAdditivesUsingTemplateRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePlantBatchesAdditivesUsingTemplate(ctx, licenseNumber, body)
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
// POST /plantbatches/v2/packages
// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreatePlantBatchesPackages(ctx context.Context, body []*models.PlantBatchesCreatePlantBatchesPackagesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePlantBatchesPackages(ctx, licenseNumber, body)
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
// POST /plantbatches/v2/plantings
// Creates new plantings for a Facility by generating plant batches based on provided planting details.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreatePlantBatchesPlantings(ctx context.Context, body []*models.PlantBatchesCreatePlantBatchesPlantingsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePlantBatchesPlantings(ctx, licenseNumber, body)
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
// POST /plantbatches/v2/waste
// Records waste information for plant batches based on the submitted data for the specified Facility.
// Permissions Required:
// - Manage Plants Waste
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreatePlantBatchesWaste(ctx context.Context, body []*models.PlantBatchesCreatePlantBatchesWasteRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePlantBatchesWaste(ctx, licenseNumber, body)
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
// POST /plantbatches/v2/split
// Splits an existing Plant Batch into multiple groups at the specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) CreatePlantBatchesSplit(ctx context.Context, body []*models.PlantBatchesCreateSplitRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePlantBatchesSplit(ctx, licenseNumber, body)
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
// DELETE /plantbatches/v2
// Completes the destruction of plant batches based on the provided input data.
// Permissions Required:
// - View Immature Plants
// - Destroy Immature Plants
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
        
func (s *PlantBatchesService) DeletePlantBatches(ctx context.Context, body []*models.PlantBatchesDeletePlantBatchesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeletePlantBatches(ctx, licenseNumber)
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
// GET /plantbatches/v2/active
// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
// Permissions Required:
// - View Immature Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PlantBatchesService) GetActivePlantBatches(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PlantBatch], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetActivePlantBatches(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PlantBatch]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PlantBatch]: %w", err)
    }
    return typed, nil
    
}
// GET /plantbatches/v2/inactive
// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
// Permissions Required:
// - View Immature Plants
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PlantBatchesService) GetInactivePlantBatches(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PlantBatch], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInactivePlantBatches(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PlantBatch]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PlantBatch]: %w", err)
    }
    return typed, nil
    
}
// GET /plantbatches/v2/{id}
// Retrieves a Plant Batch by Id.
// Permissions Required:
// - View Immature Plants
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *PlantBatchesService) GetPlantBatchesById(ctx context.Context, id string, licenseNumber string) (*models.PlantBatch, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPlantBatchesById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PlantBatch
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PlantBatch: %w", err)
    }
    return typed, nil
    
}
// GET /plantbatches/v2/types
// Retrieves a list of plant batch types.
// Parameters:
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PlantBatchesService) GetPlantBatchesTypes(ctx context.Context, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PlantBatchesType], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPlantBatchesTypes(ctx, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PlantBatchesType]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PlantBatchesType]: %w", err)
    }
    return typed, nil
    
}
// GET /plantbatches/v2/waste
// Retrieves waste details associated with plant batches at a specified Facility.
// Permissions Required:
// - View Plants Waste
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PlantBatchesService) GetPlantBatchesWaste(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PlantBatchesWaste], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPlantBatchesWaste(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PlantBatchesWaste]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PlantBatchesWaste]: %w", err)
    }
    return typed, nil
    
}
// GET /plantbatches/v2/waste/reasons
// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (s *PlantBatchesService) GetPlantBatchesWasteReasons(ctx context.Context, licenseNumber string) ([]*models.PlantBatchesWasteReason, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPlantBatchesWasteReasons(ctx, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.PlantBatchesWasteReason
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.PlantBatchesWasteReason: %w", err)
    }
    return typed, nil
    
}
// PUT /plantbatches/v2/name
// Renames plant batches at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) UpdatePlantBatchesName(ctx context.Context, body []*models.PlantBatchesUpdateNameRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePlantBatchesName(ctx, licenseNumber, body)
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
// PUT /plantbatches/v2/location
// Moves one or more plant batches to new locations with in a specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) UpdatePlantBatchesLocation(ctx context.Context, body []*models.PlantBatchesUpdatePlantBatchesLocationRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePlantBatchesLocation(ctx, licenseNumber, body)
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
// PUT /plantbatches/v2/strain
// Changes the strain of plant batches at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) UpdatePlantBatchesStrain(ctx context.Context, body []*models.PlantBatchesUpdatePlantBatchesStrainRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePlantBatchesStrain(ctx, licenseNumber, body)
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
// PUT /plantbatches/v2/tag
// Replaces tags for plant batches at a specified Facility.
// Permissions Required:
// - View Veg/Flower Plants
// - Manage Veg/Flower Plants Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PlantBatchesService) UpdatePlantBatchesTag(ctx context.Context, body []*models.PlantBatchesUpdatePlantBatchesTagRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePlantBatchesTag(ctx, licenseNumber, body)
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


