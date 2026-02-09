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



// IterateGetActivePackages iterates over all pages of GetActivePackages and executes handler for each item.
func IterateGetActivePackages(ctx context.Context, s *services.PackagesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.PackagesPackage) error) error {
    page := 1
    for {
        
        response, err := s.GetActivePackages(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetActivePackages retrieves all items updated within the specified time window.
func SyncGetActivePackages(ctx context.Context, s *services.PackagesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.PackagesPackage, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.PackagesPackage) error) error {
        return IterateGetActivePackages(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetActivePackagesParallel syncs across multiple targets in parallel.
func SyncGetActivePackagesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.PackagesPackage, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.PackagesPackage, error) {
        return SyncGetActivePackages(ctx, w.Packages, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetPackagesAdjustReasons iterates over all pages of GetPackagesAdjustReasons and executes handler for each item.
func IterateGetPackagesAdjustReasons(ctx context.Context, s *services.PackagesService, licenseNumber string, pageSize string, handler func(*models.AdjustReason) error) error {
    page := 1
    for {
        
        response, err := s.GetPackagesAdjustReasons(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetPackagesAdjustments iterates over all pages of GetPackagesAdjustments and executes handler for each item.
func IterateGetPackagesAdjustments(ctx context.Context, s *services.PackagesService, licenseNumber string, pageSize string, handler func(*models.Adjustment) error) error {
    page := 1
    for {
        
        response, err := s.GetPackagesAdjustments(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetInTransitPackages iterates over all pages of GetInTransitPackages and executes handler for each item.
func IterateGetInTransitPackages(ctx context.Context, s *services.PackagesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.InTransit) error) error {
    page := 1
    for {
        
        response, err := s.GetInTransitPackages(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetInTransitPackages retrieves all items updated within the specified time window.
func SyncGetInTransitPackages(ctx context.Context, s *services.PackagesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.InTransit, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.InTransit) error) error {
        return IterateGetInTransitPackages(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetInTransitPackagesParallel syncs across multiple targets in parallel.
func SyncGetInTransitPackagesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.InTransit, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.InTransit, error) {
        return SyncGetInTransitPackages(ctx, w.Packages, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetInactivePackages iterates over all pages of GetInactivePackages and executes handler for each item.
func IterateGetInactivePackages(ctx context.Context, s *services.PackagesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.PackagesPackage) error) error {
    page := 1
    for {
        
        response, err := s.GetInactivePackages(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetInactivePackages retrieves all items updated within the specified time window.
func SyncGetInactivePackages(ctx context.Context, s *services.PackagesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.PackagesPackage, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.PackagesPackage) error) error {
        return IterateGetInactivePackages(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetInactivePackagesParallel syncs across multiple targets in parallel.
func SyncGetInactivePackagesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.PackagesPackage, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.PackagesPackage, error) {
        return SyncGetInactivePackages(ctx, w.Packages, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetPackagesLabSamples iterates over all pages of GetPackagesLabSamples and executes handler for each item.
func IterateGetPackagesLabSamples(ctx context.Context, s *services.PackagesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.PackagesPackage) error) error {
    page := 1
    for {
        
        response, err := s.GetPackagesLabSamples(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetPackagesLabSamples retrieves all items updated within the specified time window.
func SyncGetPackagesLabSamples(ctx context.Context, s *services.PackagesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.PackagesPackage, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.PackagesPackage) error) error {
        return IterateGetPackagesLabSamples(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetPackagesLabSamplesParallel syncs across multiple targets in parallel.
func SyncGetPackagesLabSamplesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.PackagesPackage, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.PackagesPackage, error) {
        return SyncGetPackagesLabSamples(ctx, w.Packages, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetOnHoldPackages iterates over all pages of GetOnHoldPackages and executes handler for each item.
func IterateGetOnHoldPackages(ctx context.Context, s *services.PackagesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.PackagesPackage) error) error {
    page := 1
    for {
        
        response, err := s.GetOnHoldPackages(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetOnHoldPackages retrieves all items updated within the specified time window.
func SyncGetOnHoldPackages(ctx context.Context, s *services.PackagesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.PackagesPackage, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.PackagesPackage) error) error {
        return IterateGetOnHoldPackages(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetOnHoldPackagesParallel syncs across multiple targets in parallel.
func SyncGetOnHoldPackagesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.PackagesPackage, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.PackagesPackage, error) {
        return SyncGetOnHoldPackages(ctx, w.Packages, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetPackagesTransferred iterates over all pages of GetPackagesTransferred and executes handler for each item.
func IterateGetPackagesTransferred(ctx context.Context, s *services.PackagesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.PackagesPackage) error) error {
    page := 1
    for {
        
        response, err := s.GetPackagesTransferred(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetPackagesTransferred retrieves all items updated within the specified time window.
func SyncGetPackagesTransferred(ctx context.Context, s *services.PackagesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.PackagesPackage, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.PackagesPackage) error) error {
        return IterateGetPackagesTransferred(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetPackagesTransferredParallel syncs across multiple targets in parallel.
func SyncGetPackagesTransferredParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.PackagesPackage, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.PackagesPackage, error) {
        return SyncGetPackagesTransferred(ctx, w.Packages, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}


