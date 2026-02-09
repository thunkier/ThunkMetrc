package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type SalesService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /sales/v2/deliveries
// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) CreateSalesDeliveries(ctx context.Context, body []*models.SalesCreateDeliveriesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateSalesDeliveries(ctx, licenseNumber, body)
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
// POST /sales/v2/deliveries/retailer/depart
// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) CreateSalesDeliveriesRetailerDepart(ctx context.Context, body []*models.SalesCreateDeliveriesRetailerDepartRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateSalesDeliveriesRetailerDepart(ctx, licenseNumber, body)
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
// POST /sales/v2/deliveries/retailer/end
// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) CreateSalesDeliveriesRetailerEnd(ctx context.Context, body []*models.SalesCreateDeliveriesRetailerEndRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateSalesDeliveriesRetailerEnd(ctx, licenseNumber, body)
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
// POST /sales/v2/deliveries/retailer/restock
// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) CreateSalesDeliveriesRetailerRestock(ctx context.Context, body []*models.SalesCreateDeliveriesRetailerRestockRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateSalesDeliveriesRetailerRestock(ctx, licenseNumber, body)
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
// POST /sales/v2/receipts
// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales (Write)
// - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
// - Industry/Facility Type/Advanced Sales
// - WebApi Sales Read Write State (All or WriteOnly)
// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) CreateSalesReceipts(ctx context.Context, body []*models.SalesCreateReceiptsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateSalesReceipts(ctx, licenseNumber, body)
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
// POST /sales/v2/deliveries/retailer
// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) CreateSalesDeliveriesRetailer(ctx context.Context, body []*models.SalesCreateSalesDeliveriesRetailerRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateSalesDeliveriesRetailer(ctx, licenseNumber, body)
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
// DELETE /sales/v2/deliveries/{id}
// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
// Permissions Required:
// - Manage Sales Delivery
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *SalesService) DeleteSalesDeliveryById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteSalesDeliveryById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// DELETE /sales/v2/deliveries/retailer/{id}
// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - Manage Retailer Delivery
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *SalesService) DeleteSalesDeliveryRetailerById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteSalesDeliveryRetailerById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// DELETE /sales/v2/receipts/{id}
// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
// Permissions Required:
// - Manage Sales
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *SalesService) DeleteSalesReceiptById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteSalesReceiptById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /sales/v2/deliveries/active
// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
// Permissions Required:
// - View Sales Delivery
// - Manage Sales Delivery
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *SalesService) GetSalesActiveDeliveries(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.ActiveDelivery], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesActiveDeliveries(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.ActiveDelivery]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.ActiveDelivery]: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/deliveries/retailer/active
// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
// Permissions Required:
// - View Retailer Delivery
// - Manage Retailer Delivery
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *SalesService) GetSalesActiveDeliveriesRetailer(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.ActiveDeliveriesRetailer], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesActiveDeliveriesRetailer(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.ActiveDeliveriesRetailer]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.ActiveDeliveriesRetailer]: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/receipts/active
// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
// Permissions Required:
// - View Sales
// - Manage Sales
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *SalesService) GetSalesActiveReceipts(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.ActiveReceipt], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesActiveReceipts(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.ActiveReceipt]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.ActiveReceipt]: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/counties
// Returns a list of counties available for sales deliveries.
// Parameters:
func (s *SalesService) GetSalesCounties(ctx context.Context, ) (*models.PaginatedResponse[*models.County], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesCounties(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.County]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.County]: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/customertypes
// Returns a list of customer types.
// Parameters:
func (s *SalesService) GetSalesCustomerTypes(ctx context.Context, ) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesCustomerTypes(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /sales/v2/deliveries/returnreasons
// Returns a list of return reasons for sales deliveries based on the provided License Number.
// Permissions Required:
// - Sales Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *SalesService) GetSalesDeliveriesReturnReasons(ctx context.Context, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.DeliveriesReturnReason], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesDeliveriesReturnReasons(ctx, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.DeliveriesReturnReason]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.DeliveriesReturnReason]: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/deliveries/retailer/{id}
// Retrieves a retailer delivery record by its ID, with an optional License Number.
// Permissions Required:
// - View Retailer Delivery
// - Manage Retailer Delivery
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *SalesService) GetSalesDeliveryRetailerById(ctx context.Context, id string, licenseNumber string) (*models.DeliveryRetailer, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesDeliveryRetailerById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.DeliveryRetailer
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.DeliveryRetailer: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/deliveries/inactive
// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
// Permissions Required:
// - View Sales Delivery
// - Manage Sales Delivery
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *SalesService) GetSalesInactiveDeliveries(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.InactiveDelivery], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesInactiveDeliveries(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.InactiveDelivery]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.InactiveDelivery]: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/deliveries/retailer/inactive
// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
// Permissions Required:
// - View Retailer Delivery
// - Manage Retailer Delivery
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *SalesService) GetSalesInactiveDeliveriesRetailer(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.InactiveDeliveriesRetailer], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesInactiveDeliveriesRetailer(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.InactiveDeliveriesRetailer]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.InactiveDeliveriesRetailer]: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/receipts/inactive
// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
// Permissions Required:
// - View Sales
// - Manage Sales
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *SalesService) GetSalesInactiveReceipts(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.InactiveReceipt], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesInactiveReceipts(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.InactiveReceipt]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.InactiveReceipt]: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/patientregistration/locations
// Returns a list of valid Patient registration locations for sales.
// Parameters:
func (s *SalesService) GetSalesPatientRegistrationLocations(ctx context.Context, ) ([]*models.PatientRegistrationLocation, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesPatientRegistrationLocations(ctx)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.PatientRegistrationLocation
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.PatientRegistrationLocation: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/paymenttypes
// Returns a list of available payment types for the specified License Number.
// Permissions Required:
// - View Sales Delivery
// - Manage Sales Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (s *SalesService) GetSalesPaymentTypes(ctx context.Context, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesPaymentTypes(ctx, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /sales/v2/receipts/{id}
// Retrieves a sales receipt by its Id, with an optional License Number.
// Permissions Required:
// - View Sales
// - Manage Sales
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *SalesService) GetSalesReceiptById(ctx context.Context, id string, licenseNumber string) (*models.ActiveReceipt, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesReceiptById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.ActiveReceipt
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.ActiveReceipt: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/receipts/external/{externalNumber}
// Retrieves a Sales Receipt by its external number, with an optional License Number.
// Permissions Required:
// - View Sales
// - Manage Sales
// Parameters:
// - externalNumber (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *SalesService) GetSalesReceiptsExternalByExternalNumber(ctx context.Context, externalNumber string, licenseNumber string) ([]*models.ReceiptsExternalByExternalNumber, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesReceiptsExternalByExternalNumber(ctx, externalNumber, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed []*models.ReceiptsExternalByExternalNumber
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to []*models.ReceiptsExternalByExternalNumber: %w", err)
    }
    return typed, nil
    
}
// GET /sales/v2/deliveries/{id}
// Retrieves a sales delivery record by its Id, with an optional License Number.
// Permissions Required:
// - View Sales Delivery
// - Manage Sales Delivery
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *SalesService) GetSalesDeliveryById(ctx context.Context, id string, licenseNumber string) (*models.SalesDelivery, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetSalesDeliveryById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.SalesDelivery
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.SalesDelivery: %w", err)
    }
    return typed, nil
    
}
// PUT /sales/v2/deliveries
// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - Manage Sales Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesDeliveries(ctx context.Context, body []*models.SalesUpdateDeliveriesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesDeliveries(ctx, licenseNumber, body)
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
// PUT /sales/v2/deliveries/complete
// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
// Permissions Required:
// - Manage Sales Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesDeliveriesComplete(ctx context.Context, body []*models.SalesUpdateDeliveriesCompleteRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesDeliveriesComplete(ctx, licenseNumber, body)
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
// PUT /sales/v2/deliveries/hub
// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - Manage Sales Delivery, Manage Sales Delivery Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesDeliveriesHub(ctx context.Context, body []*models.SalesUpdateDeliveriesHubRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesDeliveriesHub(ctx, licenseNumber, body)
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
// PUT /sales/v2/deliveries/hub/accept
// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
// Permissions Required:
// - Manage Sales Delivery Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesDeliveriesHubAccept(ctx context.Context, body []*models.SalesUpdateDeliveriesHubAcceptRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesDeliveriesHubAccept(ctx, licenseNumber, body)
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
// PUT /sales/v2/deliveries/hub/depart
// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
// Permissions Required:
// - Manage Sales Delivery Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesDeliveriesHubDepart(ctx context.Context, body []*models.SalesUpdateDeliveriesHubDepartRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesDeliveriesHubDepart(ctx, licenseNumber, body)
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
// PUT /sales/v2/deliveries/hub/verifyID
// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
// Permissions Required:
// - Manage Sales Delivery Hub
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesDeliveriesHubVerifyID(ctx context.Context, body []*models.SalesUpdateDeliveriesHubVerifyIDRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesDeliveriesHubVerifyID(ctx, licenseNumber, body)
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
// PUT /sales/v2/deliveries/retailer
// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
// - Industry/Facility Type/Retailer Delivery
// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
// - Manage Retailer Delivery
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesDeliveriesRetailer(ctx context.Context, body []*models.SalesUpdateDeliveriesRetailerRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesDeliveriesRetailer(ctx, licenseNumber, body)
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
// PUT /sales/v2/receipts
// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// Permissions Required:
// - Manage Sales
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesReceipts(ctx context.Context, body []*models.SalesUpdateReceiptsRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesReceipts(ctx, licenseNumber, body)
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
// PUT /sales/v2/receipts/finalize
// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
// Permissions Required:
// - Manage Sales
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesReceiptsFinalize(ctx context.Context, body []*models.SalesUpdateReceiptsFinalizeRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesReceiptsFinalize(ctx, licenseNumber, body)
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
// PUT /sales/v2/receipts/unfinalize
// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
// Permissions Required:
// - Manage Sales
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *SalesService) UpdateSalesReceiptsUnfinalize(ctx context.Context, body []*models.SalesUpdateReceiptsUnfinalizeRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateSalesReceiptsUnfinalize(ctx, licenseNumber, body)
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


