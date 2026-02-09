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



// IterateGetSalesActiveDeliveries iterates over all pages of GetSalesActiveDeliveries and executes handler for each item.
func IterateGetSalesActiveDeliveries(ctx context.Context, s *services.SalesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.ActiveDelivery) error) error {
    page := 1
    for {
        
        response, err := s.GetSalesActiveDeliveries(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetSalesActiveDeliveries retrieves all items updated within the specified time window.
func SyncGetSalesActiveDeliveries(ctx context.Context, s *services.SalesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.ActiveDelivery, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.ActiveDelivery) error) error {
        return IterateGetSalesActiveDeliveries(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetSalesActiveDeliveriesParallel syncs across multiple targets in parallel.
func SyncGetSalesActiveDeliveriesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.ActiveDelivery, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.ActiveDelivery, error) {
        return SyncGetSalesActiveDeliveries(ctx, w.Sales, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetSalesActiveDeliveriesRetailer iterates over all pages of GetSalesActiveDeliveriesRetailer and executes handler for each item.
func IterateGetSalesActiveDeliveriesRetailer(ctx context.Context, s *services.SalesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.ActiveDeliveriesRetailer) error) error {
    page := 1
    for {
        
        response, err := s.GetSalesActiveDeliveriesRetailer(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetSalesActiveDeliveriesRetailer retrieves all items updated within the specified time window.
func SyncGetSalesActiveDeliveriesRetailer(ctx context.Context, s *services.SalesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.ActiveDeliveriesRetailer, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.ActiveDeliveriesRetailer) error) error {
        return IterateGetSalesActiveDeliveriesRetailer(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetSalesActiveDeliveriesRetailerParallel syncs across multiple targets in parallel.
func SyncGetSalesActiveDeliveriesRetailerParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.ActiveDeliveriesRetailer, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.ActiveDeliveriesRetailer, error) {
        return SyncGetSalesActiveDeliveriesRetailer(ctx, w.Sales, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetSalesActiveReceipts iterates over all pages of GetSalesActiveReceipts and executes handler for each item.
func IterateGetSalesActiveReceipts(ctx context.Context, s *services.SalesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.ActiveReceipt) error) error {
    page := 1
    for {
        
        response, err := s.GetSalesActiveReceipts(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetSalesActiveReceipts retrieves all items updated within the specified time window.
func SyncGetSalesActiveReceipts(ctx context.Context, s *services.SalesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.ActiveReceipt, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.ActiveReceipt) error) error {
        return IterateGetSalesActiveReceipts(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetSalesActiveReceiptsParallel syncs across multiple targets in parallel.
func SyncGetSalesActiveReceiptsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.ActiveReceipt, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.ActiveReceipt, error) {
        return SyncGetSalesActiveReceipts(ctx, w.Sales, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetSalesDeliveriesReturnReasons iterates over all pages of GetSalesDeliveriesReturnReasons and executes handler for each item.
func IterateGetSalesDeliveriesReturnReasons(ctx context.Context, s *services.SalesService, licenseNumber string, pageSize string, handler func(*models.DeliveriesReturnReason) error) error {
    page := 1
    for {
        
        response, err := s.GetSalesDeliveriesReturnReasons(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetSalesInactiveDeliveries iterates over all pages of GetSalesInactiveDeliveries and executes handler for each item.
func IterateGetSalesInactiveDeliveries(ctx context.Context, s *services.SalesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.InactiveDelivery) error) error {
    page := 1
    for {
        
        response, err := s.GetSalesInactiveDeliveries(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetSalesInactiveDeliveries retrieves all items updated within the specified time window.
func SyncGetSalesInactiveDeliveries(ctx context.Context, s *services.SalesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.InactiveDelivery, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.InactiveDelivery) error) error {
        return IterateGetSalesInactiveDeliveries(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetSalesInactiveDeliveriesParallel syncs across multiple targets in parallel.
func SyncGetSalesInactiveDeliveriesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.InactiveDelivery, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.InactiveDelivery, error) {
        return SyncGetSalesInactiveDeliveries(ctx, w.Sales, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetSalesInactiveDeliveriesRetailer iterates over all pages of GetSalesInactiveDeliveriesRetailer and executes handler for each item.
func IterateGetSalesInactiveDeliveriesRetailer(ctx context.Context, s *services.SalesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.InactiveDeliveriesRetailer) error) error {
    page := 1
    for {
        
        response, err := s.GetSalesInactiveDeliveriesRetailer(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetSalesInactiveDeliveriesRetailer retrieves all items updated within the specified time window.
func SyncGetSalesInactiveDeliveriesRetailer(ctx context.Context, s *services.SalesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.InactiveDeliveriesRetailer, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.InactiveDeliveriesRetailer) error) error {
        return IterateGetSalesInactiveDeliveriesRetailer(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetSalesInactiveDeliveriesRetailerParallel syncs across multiple targets in parallel.
func SyncGetSalesInactiveDeliveriesRetailerParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.InactiveDeliveriesRetailer, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.InactiveDeliveriesRetailer, error) {
        return SyncGetSalesInactiveDeliveriesRetailer(ctx, w.Sales, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetSalesInactiveReceipts iterates over all pages of GetSalesInactiveReceipts and executes handler for each item.
func IterateGetSalesInactiveReceipts(ctx context.Context, s *services.SalesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.InactiveReceipt) error) error {
    page := 1
    for {
        
        response, err := s.GetSalesInactiveReceipts(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetSalesInactiveReceipts retrieves all items updated within the specified time window.
func SyncGetSalesInactiveReceipts(ctx context.Context, s *services.SalesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.InactiveReceipt, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.InactiveReceipt) error) error {
        return IterateGetSalesInactiveReceipts(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetSalesInactiveReceiptsParallel syncs across multiple targets in parallel.
func SyncGetSalesInactiveReceiptsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.InactiveReceipt, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.InactiveReceipt, error) {
        return SyncGetSalesInactiveReceipts(ctx, w.Sales, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}


