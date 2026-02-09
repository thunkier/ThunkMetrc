package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type TransfersService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /transfers/v2/hub/arrive
// Arrive a transfer for a Facility.
// Permissions Required:
// - Manage Transfer Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransfersService) CreateTransfersHubArrive(ctx context.Context, body []*models.TransfersCreateHubArriveRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateTransfersHubArrive(ctx, licenseNumber, body)
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
// POST /transfers/v2/hub/checkin
// CheckIn a transfer for a Facility.
// Permissions Required:
// - Manage Transfer Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransfersService) CreateTransfersHubCheckin(ctx context.Context, body []*models.TransfersCreateHubCheckinRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateTransfersHubCheckin(ctx, licenseNumber, body)
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
// POST /transfers/v2/hub/checkout
// CheckOut a transfer for a Facility.
// Permissions Required:
// - Manage Transfer Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransfersService) CreateTransfersHubCheckout(ctx context.Context, body []*models.TransfersCreateHubCheckoutRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateTransfersHubCheckout(ctx, licenseNumber, body)
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
// POST /transfers/v2/hub/depart
// Depart a transfer for a Facility.
// Permissions Required:
// - Manage Transfer Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransfersService) CreateTransfersHubDepart(ctx context.Context, body []*models.TransfersCreateHubDepartRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateTransfersHubDepart(ctx, licenseNumber, body)
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
// POST /transfers/v2/external/incoming
// Creates external incoming shipment plans for a Facility.
// Permissions Required:
// - Manage Transfers
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransfersService) CreateTransfersIncomingExternal(ctx context.Context, body []*models.TransfersCreateIncomingExternalRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateTransfersIncomingExternal(ctx, licenseNumber, body)
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
// POST /transfers/v2/templates/outgoing
// Creates new transfer templates for a Facility.
// Permissions Required:
// - Manage Transfer Templates
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransfersService) CreateTransfersOutgoingTemplates(ctx context.Context, body []*models.TransfersCreateOutgoingTemplatesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateTransfersOutgoingTemplates(ctx, licenseNumber, body)
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
// DELETE /transfers/v2/external/incoming/{id}
// Voids an external incoming shipment plan for a Facility.
// Permissions Required:
// - Manage Transfers
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *TransfersService) DeleteTransfersIncomingExternalById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteTransfersIncomingExternalById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// DELETE /transfers/v2/templates/outgoing/{id}
// Archives a transfer template for a Facility.
// Permissions Required:
// - Manage Transfer Templates
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *TransfersService) DeleteTransfersOutgoingTemplateById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteTransfersOutgoingTemplateById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /transfers/v2/deliveries/packages/states
// Returns a list of available shipment Package states.
// Parameters:
func (s *TransfersService) GetTransfersDeliveriesPackagesStates(ctx context.Context, ) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersDeliveriesPackagesStates(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /transfers/v2/deliveries/{id}/packages
// Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersDeliveryPackageById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.DeliveryPackage], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersDeliveryPackageById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.DeliveryPackage]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.DeliveryPackage]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/deliveries/package/{id}/requiredlabtestbatches
// Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersDeliveryPackageRequiredlabtestbatchById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.DeliveryPackageRequiredlabtestbatch], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersDeliveryPackageRequiredlabtestbatchById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.DeliveryPackageRequiredlabtestbatch]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.DeliveryPackageRequiredlabtestbatch]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/deliveries/{id}/packages/wholesale
// Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersDeliveryPackageWholesaleById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.DeliveryPackageWholesale], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersDeliveryPackageWholesaleById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.DeliveryPackageWholesale]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.DeliveryPackageWholesale]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/deliveries/{id}/transporters
// Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersDeliveryTransporterById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.DeliveryTransporter], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersDeliveryTransporterById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.DeliveryTransporter]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.DeliveryTransporter]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/deliveries/{id}/transporters/details
// Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersDeliveryTransporterDetailById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.DeliveryTransporterDetail], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersDeliveryTransporterDetailById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.DeliveryTransporterDetail]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.DeliveryTransporterDetail]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/hub
// Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersHub(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Hub], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersHub(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Hub]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Hub]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/incoming
// Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetIncomingTransfers(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Transfer], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetIncomingTransfers(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Transfer]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Transfer]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/manifest/{id}/html
// Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *TransfersService) GetTransfersManifestHtmlById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersManifestHtmlById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /transfers/v2/manifest/{id}/pdf
// Get Transfer Manifest PDF for a given Transfer Id
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *TransfersService) GetTransfersManifestPdfById(ctx context.Context, id string, licenseNumber string) (*models.ManifestPdf, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersManifestPdfById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.ManifestPdf
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.ManifestPdf: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/templates/outgoing/{id}/deliveries
// Retrieves a list of deliveries associated with a specific transfer template.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersOutgoingTemplateDeliveryById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.TemplateDelivery], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersOutgoingTemplateDeliveryById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.TemplateDelivery]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.TemplateDelivery]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/templates/outgoing/deliveries/{id}/packages
// Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersOutgoingTemplateDeliveryPackageById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.TemplateDeliveryPackage], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersOutgoingTemplateDeliveryPackageById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.TemplateDeliveryPackage]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.TemplateDeliveryPackage]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/templates/outgoing/deliveries/{id}/transporters
// Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersOutgoingTemplateDeliveryTransporterById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.TemplateDeliveryTransporter], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersOutgoingTemplateDeliveryTransporterById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.TemplateDeliveryTransporter]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.TemplateDeliveryTransporter]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/templates/outgoing/deliveries/{id}/transporters/details
// Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - id (path param)
func (s *TransfersService) GetTransfersOutgoingTemplateDeliveryTransporterDetailById(ctx context.Context, id string) (*models.TemplateDeliveryTransporterDetail, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersOutgoingTemplateDeliveryTransporterDetailById(ctx, id)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.TemplateDeliveryTransporterDetail
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.TemplateDeliveryTransporterDetail: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/templates/outgoing
// Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Transfer Templates
// - View Transfer Templates
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersOutgoingTemplates(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Template], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersOutgoingTemplates(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Template]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Template]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/outgoing
// Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetOutgoingTransfers(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Transfer], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetOutgoingTransfers(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Transfer]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Transfer]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/rejected
// Retrieves a list of shipments with rejected packages for a Facility.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetRejectedTransfers(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.Transfer], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetRejectedTransfers(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.Transfer]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.Transfer]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/{id}/deliveries
// Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
// Permissions Required:
// - Manage Transfers
// - View Transfers
// Parameters:
// - id (path param)
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersDeliveryById(ctx context.Context, id string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.TransfersDelivery], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersDeliveryById(ctx, id, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.TransfersDelivery]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.TransfersDelivery]: %w", err)
    }
    return typed, nil
    
}
// GET /transfers/v2/types
// Retrieves a list of available transfer types for a Facility based on its license number.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *TransfersService) GetTransfersTypes(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.TransfersType], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetTransfersTypes(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.TransfersType]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.TransfersType]: %w", err)
    }
    return typed, nil
    
}
// PUT /transfers/v2/external/incoming
// Updates external incoming shipment plans for a Facility.
// Permissions Required:
// - Manage Transfers
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransfersService) UpdateTransfersIncomingExternal(ctx context.Context, body []*models.TransfersUpdateIncomingExternalRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateTransfersIncomingExternal(ctx, licenseNumber, body)
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
// PUT /transfers/v2/templates/outgoing
// Updates existing transfer templates for a Facility.
// Permissions Required:
// - Manage Transfer Templates
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *TransfersService) UpdateTransfersOutgoingTemplates(ctx context.Context, body []*models.TransfersUpdateOutgoingTemplatesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateTransfersOutgoingTemplates(ctx, licenseNumber, body)
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


