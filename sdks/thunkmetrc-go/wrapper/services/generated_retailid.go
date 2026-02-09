package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type RetailIdService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /retailid/v2/associate
// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
// - WebApi Retail ID Read Write State (All or WriteOnly)
// - Industry/View Packages
// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *RetailIdService) CreateRetailIdAssociate(ctx context.Context, body []*models.RetailIdCreateAssociateRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateRetailIdAssociate(ctx, licenseNumber, body)
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
// POST /retailid/v2/generate
// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
// - WebApi Retail ID Read Write State (All or WriteOnly)
// - Industry/View Packages
// - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *RetailIdService) CreateRetailIdGenerate(ctx context.Context, body *models.RetailIdCreateGenerateRequest, licenseNumber string) ([]*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateRetailIdGenerate(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.WriteResponse: %w", err)
    }
    return typed, nil
    
}
// POST /retailid/v2/merge
// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
// - WebApi Retail ID Read Write State (All or WriteOnly)
// - Key Value Settings/Retail ID Merge Packages Enabled
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *RetailIdService) CreateRetailIdMerge(ctx context.Context, body *models.RetailIdCreateMergeRequest, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateRetailIdMerge(ctx, licenseNumber, body)
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
// POST /retailid/v2/packages/info
// Retrieves Package information for given list of Package labels.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
// - WebApi Retail ID Read Write State (All or WriteOnly)
// - Industry/View Packages
// - Admin/Employees/Packages Page/Product Labels(Manage)
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *RetailIdService) CreateRetailIdPackagesInfo(ctx context.Context, body *models.RetailIdCreatePackagesInfoRequest, licenseNumber string) ([]*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateRetailIdPackagesInfo(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.WriteResponse: %w", err)
    }
    return typed, nil
    
}
// GET /retailid/v2/allotment
// Retrieves the available Retail Item ID quota for a facility.
// Permissions Required:
// - Download Product Labels
// - Manage Product Labels
// - Manage Tag Orders
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (s *RetailIdService) GetRetailIdAllotment(ctx context.Context, licenseNumber string) ([]*models.Allotment, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetRetailIdAllotment(ctx, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.Allotment
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.Allotment: %w", err)
    }
    return typed, nil
    
}
// GET /retailid/v2/receive/{label}
// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Manage RetailId
// - WebApi Retail ID Read Write State (All or ReadOnly)
// - Industry/View Packages
// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
// Parameters:
// - label (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *RetailIdService) GetRetailIdReceiveByLabel(ctx context.Context, label string, licenseNumber string) ([]*models.Receive, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetRetailIdReceiveByLabel(ctx, label, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.Receive
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.Receive: %w", err)
    }
    return typed, nil
    
}
// GET /retailid/v2/receive/qr/{shortCode}
// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Manage RetailId
// - WebApi Retail ID Read Write State (All or ReadOnly)
// - Industry/View Packages
// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
// Parameters:
// - shortCode (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *RetailIdService) GetRetailIdReceiveQrByShortCode(ctx context.Context, shortCode string, licenseNumber string) ([]*models.ReceiveQrByShortCode, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetRetailIdReceiveQrByShortCode(ctx, shortCode, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.ReceiveQrByShortCode
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.ReceiveQrByShortCode: %w", err)
    }
    return typed, nil
    
}


