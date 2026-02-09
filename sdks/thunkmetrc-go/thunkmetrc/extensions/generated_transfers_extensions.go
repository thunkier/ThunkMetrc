package extensions

import (
    "context"
    "strconv"
    "time"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/thunkmetrc/utils"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/services"
)



// IterateGetTransfersDeliveryPackageById iterates over all pages of GetTransfersDeliveryPackageById and executes handler for each item.
func IterateGetTransfersDeliveryPackageById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.DeliveryPackage) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersDeliveryPackageById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersDeliveryPackageRequiredlabtestbatchById iterates over all pages of GetTransfersDeliveryPackageRequiredlabtestbatchById and executes handler for each item.
func IterateGetTransfersDeliveryPackageRequiredlabtestbatchById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.DeliveryPackageRequiredlabtestbatch) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersDeliveryPackageRequiredlabtestbatchById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersDeliveryPackageWholesaleById iterates over all pages of GetTransfersDeliveryPackageWholesaleById and executes handler for each item.
func IterateGetTransfersDeliveryPackageWholesaleById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.DeliveryPackageWholesale) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersDeliveryPackageWholesaleById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersDeliveryTransporterById iterates over all pages of GetTransfersDeliveryTransporterById and executes handler for each item.
func IterateGetTransfersDeliveryTransporterById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.DeliveryTransporter) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersDeliveryTransporterById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersDeliveryTransporterDetailById iterates over all pages of GetTransfersDeliveryTransporterDetailById and executes handler for each item.
func IterateGetTransfersDeliveryTransporterDetailById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.DeliveryTransporterDetail) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersDeliveryTransporterDetailById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersHub iterates over all pages of GetTransfersHub and executes handler for each item.
func IterateGetTransfersHub(ctx context.Context, s *services.TransfersService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Hub) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersHub(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// SyncGetTransfersHub retrieves all items updated within the specified time window.
func SyncGetTransfersHub(ctx context.Context, s *services.TransfersService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Hub, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Hub) error) error {
        return IterateGetTransfersHub(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetTransfersHubParallel syncs across multiple targets in parallel.
func SyncGetTransfersHubParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Hub, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Hub, error) {
        return SyncGetTransfersHub(ctx, w.Transfers, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetIncomingTransfers iterates over all pages of GetIncomingTransfers and executes handler for each item.
func IterateGetIncomingTransfers(ctx context.Context, s *services.TransfersService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Transfer) error) error {
    page := 1
    for {
        
        response, err := s.GetIncomingTransfers(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// SyncGetIncomingTransfers retrieves all items updated within the specified time window.
func SyncGetIncomingTransfers(ctx context.Context, s *services.TransfersService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Transfer, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Transfer) error) error {
        return IterateGetIncomingTransfers(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetIncomingTransfersParallel syncs across multiple targets in parallel.
func SyncGetIncomingTransfersParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Transfer, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Transfer, error) {
        return SyncGetIncomingTransfers(ctx, w.Transfers, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetTransfersOutgoingTemplateDeliveryById iterates over all pages of GetTransfersOutgoingTemplateDeliveryById and executes handler for each item.
func IterateGetTransfersOutgoingTemplateDeliveryById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.TemplateDelivery) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersOutgoingTemplateDeliveryById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersOutgoingTemplateDeliveryPackageById iterates over all pages of GetTransfersOutgoingTemplateDeliveryPackageById and executes handler for each item.
func IterateGetTransfersOutgoingTemplateDeliveryPackageById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.TemplateDeliveryPackage) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersOutgoingTemplateDeliveryPackageById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersOutgoingTemplateDeliveryTransporterById iterates over all pages of GetTransfersOutgoingTemplateDeliveryTransporterById and executes handler for each item.
func IterateGetTransfersOutgoingTemplateDeliveryTransporterById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.TemplateDeliveryTransporter) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersOutgoingTemplateDeliveryTransporterById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersOutgoingTemplates iterates over all pages of GetTransfersOutgoingTemplates and executes handler for each item.
func IterateGetTransfersOutgoingTemplates(ctx context.Context, s *services.TransfersService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Template) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersOutgoingTemplates(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// SyncGetTransfersOutgoingTemplates retrieves all items updated within the specified time window.
func SyncGetTransfersOutgoingTemplates(ctx context.Context, s *services.TransfersService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Template, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Template) error) error {
        return IterateGetTransfersOutgoingTemplates(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetTransfersOutgoingTemplatesParallel syncs across multiple targets in parallel.
func SyncGetTransfersOutgoingTemplatesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Template, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Template, error) {
        return SyncGetTransfersOutgoingTemplates(ctx, w.Transfers, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetOutgoingTransfers iterates over all pages of GetOutgoingTransfers and executes handler for each item.
func IterateGetOutgoingTransfers(ctx context.Context, s *services.TransfersService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Transfer) error) error {
    page := 1
    for {
        
        response, err := s.GetOutgoingTransfers(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// SyncGetOutgoingTransfers retrieves all items updated within the specified time window.
func SyncGetOutgoingTransfers(ctx context.Context, s *services.TransfersService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Transfer, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Transfer) error) error {
        return IterateGetOutgoingTransfers(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetOutgoingTransfersParallel syncs across multiple targets in parallel.
func SyncGetOutgoingTransfersParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Transfer, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Transfer, error) {
        return SyncGetOutgoingTransfers(ctx, w.Transfers, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetRejectedTransfers iterates over all pages of GetRejectedTransfers and executes handler for each item.
func IterateGetRejectedTransfers(ctx context.Context, s *services.TransfersService, licenseNumber string, pageSize string, handler func(*models.Transfer) error) error {
    page := 1
    for {
        
        response, err := s.GetRejectedTransfers(ctx, licenseNumber, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersDeliveryById iterates over all pages of GetTransfersDeliveryById and executes handler for each item.
func IterateGetTransfersDeliveryById(ctx context.Context, s *services.TransfersService, id string, pageSize string, handler func(*models.TransfersDelivery) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersDeliveryById(ctx, id, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}

// IterateGetTransfersTypes iterates over all pages of GetTransfersTypes and executes handler for each item.
func IterateGetTransfersTypes(ctx context.Context, s *services.TransfersService, licenseNumber string, pageSize string, handler func(*models.TransfersType) error) error {
    page := 1
    for {
        
        response, err := s.GetTransfersTypes(ctx, licenseNumber, strconv.Itoa(page), pageSize)
        if err != nil {
            return err
        }
        
        // Handle response
        if len(response.Data) == 0 {
            break
        }
        for _, item := range response.Data {
            if err := handler(item); err != nil {
                return err
            }
        }
        if len(response.Data) < 20 { // Assuming default page size
            // TODO: Use metadata from response if available.
             break
        }
        
        page++
    }
    return nil
}


