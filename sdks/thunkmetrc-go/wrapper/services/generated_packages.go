package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type PackagesService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /packages/v2/adjust
// Records a list of adjustments for packages at a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) CreateAdjustPackages(ctx context.Context, body []*models.PackagesCreateAdjustPackagesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateAdjustPackages(ctx, licenseNumber, body)
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
// POST /packages/v2
// Creates new packages for a specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) CreatePackagesPackages(ctx context.Context, body []*models.PackagesCreatePackagesPackagesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePackagesPackages(ctx, licenseNumber, body)
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
// POST /packages/v2/plantings
// Creates new plantings from packages for a specified Facility.
// Permissions Required:
// - View Immature Plants
// - Manage Immature Plants
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) CreatePackagesPlantings(ctx context.Context, body []*models.PackagesCreatePackagesPlantingsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePackagesPlantings(ctx, licenseNumber, body)
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
// POST /packages/v2/testing
// Creates new packages for testing for a specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) CreatePackagesTesting(ctx context.Context, body []*models.PackagesCreateTestingRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreatePackagesTesting(ctx, licenseNumber, body)
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
// DELETE /packages/v2/{id}
// Discontinues a Package at a specific Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *PackagesService) DeletePackagesById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeletePackagesById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /packages/v2/finish
// Updates a list of packages as finished for a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) FinishPackagesPackages(ctx context.Context, body []*models.PackagesFinishPackagesRequestItem, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.FinishPackagesPackages(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /packages/v2/finishedgood/flag
// Flags one or more Packages at the specified Facility as Finished Goods.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) FinishedgoodFlagPackages(ctx context.Context, body []*models.PackagesFinishedgoodFlagRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.FinishedgoodFlagPackages(ctx, licenseNumber, body)
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
// PUT /packages/v2/finishedgood/unflag
// Removes the Finished Good flag one or more Packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) FinishedgoodUnflagPackages(ctx context.Context, body []*models.PackagesFinishedgoodUnflagRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.FinishedgoodUnflagPackages(ctx, licenseNumber, body)
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
// GET /packages/v2/active
// Retrieves a list of active packages for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PackagesService) GetActivePackages(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PackagesPackage], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetActivePackages(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PackagesPackage]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PackagesPackage]: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/adjust/reasons
// Retrieves a list of adjustment reasons for packages at a specified Facility.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PackagesService) GetPackagesAdjustReasons(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.AdjustReason], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPackagesAdjustReasons(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.AdjustReason]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.AdjustReason]: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/adjustments
// Retrieves the Package Adjustments for a Facility
// Permissions Required:
// - View Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PackagesService) GetPackagesAdjustments(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Adjustment], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPackagesAdjustments(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Adjustment]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Adjustment]: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/intransit
// Retrieves a list of packages in transit for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PackagesService) GetInTransitPackages(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.InTransit], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInTransitPackages(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.InTransit]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.InTransit]: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/inactive
// Retrieves a list of inactive packages for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PackagesService) GetInactivePackages(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PackagesPackage], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInactivePackages(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PackagesPackage]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PackagesPackage]: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/labsamples
// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PackagesService) GetPackagesLabSamples(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PackagesPackage], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPackagesLabSamples(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PackagesPackage]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PackagesPackage]: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/onhold
// Retrieves a list of packages on hold for a specified Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PackagesService) GetOnHoldPackages(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PackagesPackage], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetOnHoldPackages(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PackagesPackage]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PackagesPackage]: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/{id}
// Retrieves a Package by its Id.
// Permissions Required:
// - View Packages
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *PackagesService) GetPackagesById(ctx context.Context, id string, licenseNumber string) (*models.PackagesPackage, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPackagesById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PackagesPackage
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PackagesPackage: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/{label}
// Retrieves a Package by its label.
// Permissions Required:
// - View Packages
// Parameters:
// - label (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *PackagesService) GetPackagesByLabel(ctx context.Context, label string, licenseNumber string) ([]*models.PackagesPackage, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPackagesByLabel(ctx, label, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.PackagesPackage
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.PackagesPackage: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/types
// Retrieves a list of available Package types.
// Parameters:
func (s *PackagesService) GetPackagesTypes(ctx context.Context, ) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPackagesTypes(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /packages/v2/{id}/source/harvests
// Retrieves the source harvests for a Package by its Id.
// Permissions Required:
// - View Package Source Harvests
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *PackagesService) GetPackagesSourceHarvestById(ctx context.Context, id string, licenseNumber string) (*models.SourceHarvest, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPackagesSourceHarvestById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.SourceHarvest
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.SourceHarvest: %w", err)
    }
    return typed, nil
    
}
// GET /packages/v2/transferred
// Retrieves a list of transferred packages for a specific Facility.
// Permissions Required:
// - View Packages
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *PackagesService) GetPackagesTransferred(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.PackagesPackage], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetPackagesTransferred(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.PackagesPackage]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.PackagesPackage]: %w", err)
    }
    return typed, nil
    
}
// PUT /packages/v2/unfinish
// Updates a list of packages as unfinished for a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UnfinishPackagesPackages(ctx context.Context, body []*models.PackagesUnfinishPackagesRequestItem, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UnfinishPackagesPackages(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /packages/v2/adjust
// Set the final quantity for a Package.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdateAdjustPackages(ctx context.Context, body []*models.PackagesUpdateAdjustPackagesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateAdjustPackages(ctx, licenseNumber, body)
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
// PUT /packages/v2/decontaminate
// Updates the Product decontaminate information for a list of packages at a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesDecontaminate(ctx context.Context, body []*models.PackagesUpdateDecontaminateRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesDecontaminate(ctx, licenseNumber, body)
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
// PUT /packages/v2/donation/flag
// Flags one or more packages for donation at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesDonationFlag(ctx context.Context, body []*models.PackagesUpdateDonationFlagRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesDonationFlag(ctx, licenseNumber, body)
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
// PUT /packages/v2/donation/unflag
// Removes the donation flag from one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesDonationUnflag(ctx context.Context, body []*models.PackagesUpdateDonationUnflagRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesDonationUnflag(ctx, licenseNumber, body)
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
// PUT /packages/v2/externalid
// Updates the external identifiers for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Package Inventory
// - External Id Enabled
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesExternalid(ctx context.Context, body []*models.PackagesUpdateExternalidRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesExternalid(ctx, licenseNumber, body)
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
// PUT /packages/v2/item
// Updates the associated Item for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesItem(ctx context.Context, body []*models.PackagesUpdateItemRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesItem(ctx, licenseNumber, body)
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
// PUT /packages/v2/labtests/required
// Updates the list of required lab test batches for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesLabtestsRequired(ctx context.Context, body []*models.PackagesUpdateLabtestsRequiredRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesLabtestsRequired(ctx, licenseNumber, body)
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
// PUT /packages/v2/note
// Updates notes associated with one or more packages for the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// - Manage Package Notes
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesNote(ctx context.Context, body []*models.PackagesUpdateNoteRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesNote(ctx, licenseNumber, body)
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
// PUT /packages/v2/location
// Updates the Location and Sublocation for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesLocation(ctx context.Context, body []*models.PackagesUpdatePackagesLocationRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesLocation(ctx, licenseNumber, body)
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
// PUT /packages/v2/remediate
// Updates a list of Product remediations for packages at a specific Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesRemediate(ctx context.Context, body []*models.PackagesUpdateRemediateRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesRemediate(ctx, licenseNumber, body)
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
// PUT /packages/v2/tradesample/flag
// Flags or unflags one or more packages at the specified Facility as trade samples.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesTradeSampleFlag(ctx context.Context, body []*models.PackagesUpdateTradeSampleFlagRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesTradeSampleFlag(ctx, licenseNumber, body)
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
// PUT /packages/v2/tradesample/unflag
// Removes the trade sample flag from one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Manage Packages Inventory
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesTradeSampleUnflag(ctx context.Context, body []*models.PackagesUpdateTradeSampleUnflagRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesTradeSampleUnflag(ctx, licenseNumber, body)
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
// PUT /packages/v2/usebydate
// Updates the use-by date for one or more packages at the specified Facility.
// Permissions Required:
// - View Packages
// - Create/Submit/Discontinue Packages
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *PackagesService) UpdatePackagesUseByDate(ctx context.Context, body []*models.PackagesUpdateUseByDateRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdatePackagesUseByDate(ctx, licenseNumber, body)
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


