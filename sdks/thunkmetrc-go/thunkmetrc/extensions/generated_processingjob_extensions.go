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



// IterateGetProcessingJobActiveJobTypes iterates over all pages of GetProcessingJobActiveJobTypes and executes handler for each item.
func IterateGetProcessingJobActiveJobTypes(ctx context.Context, s *services.ProcessingJobService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.ActiveJobType) error) error {
    page := 1
    for {
        
        response, err := s.GetProcessingJobActiveJobTypes(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetProcessingJobActiveJobTypes retrieves all items updated within the specified time window.
func SyncGetProcessingJobActiveJobTypes(ctx context.Context, s *services.ProcessingJobService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.ActiveJobType, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.ActiveJobType) error) error {
        return IterateGetProcessingJobActiveJobTypes(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetProcessingJobActiveJobTypesParallel syncs across multiple targets in parallel.
func SyncGetProcessingJobActiveJobTypesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.ActiveJobType, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.ActiveJobType, error) {
        return SyncGetProcessingJobActiveJobTypes(ctx, w.ProcessingJob, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetActiveProcessingJob iterates over all pages of GetActiveProcessingJob and executes handler for each item.
func IterateGetActiveProcessingJob(ctx context.Context, s *services.ProcessingJobService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.ProcessingJob) error) error {
    page := 1
    for {
        
        response, err := s.GetActiveProcessingJob(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetActiveProcessingJob retrieves all items updated within the specified time window.
func SyncGetActiveProcessingJob(ctx context.Context, s *services.ProcessingJobService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.ProcessingJob, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.ProcessingJob) error) error {
        return IterateGetActiveProcessingJob(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetActiveProcessingJobParallel syncs across multiple targets in parallel.
func SyncGetActiveProcessingJobParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.ProcessingJob, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.ProcessingJob, error) {
        return SyncGetActiveProcessingJob(ctx, w.ProcessingJob, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetProcessingJobInactiveJobTypes iterates over all pages of GetProcessingJobInactiveJobTypes and executes handler for each item.
func IterateGetProcessingJobInactiveJobTypes(ctx context.Context, s *services.ProcessingJobService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.InactiveJobType) error) error {
    page := 1
    for {
        
        response, err := s.GetProcessingJobInactiveJobTypes(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetProcessingJobInactiveJobTypes retrieves all items updated within the specified time window.
func SyncGetProcessingJobInactiveJobTypes(ctx context.Context, s *services.ProcessingJobService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.InactiveJobType, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.InactiveJobType) error) error {
        return IterateGetProcessingJobInactiveJobTypes(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetProcessingJobInactiveJobTypesParallel syncs across multiple targets in parallel.
func SyncGetProcessingJobInactiveJobTypesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.InactiveJobType, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.InactiveJobType, error) {
        return SyncGetProcessingJobInactiveJobTypes(ctx, w.ProcessingJob, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetInactiveProcessingJob iterates over all pages of GetInactiveProcessingJob and executes handler for each item.
func IterateGetInactiveProcessingJob(ctx context.Context, s *services.ProcessingJobService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.ProcessingJob) error) error {
    page := 1
    for {
        
        response, err := s.GetInactiveProcessingJob(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetInactiveProcessingJob retrieves all items updated within the specified time window.
func SyncGetInactiveProcessingJob(ctx context.Context, s *services.ProcessingJobService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.ProcessingJob, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.ProcessingJob) error) error {
        return IterateGetInactiveProcessingJob(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetInactiveProcessingJobParallel syncs across multiple targets in parallel.
func SyncGetInactiveProcessingJobParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.ProcessingJob, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.ProcessingJob, error) {
        return SyncGetInactiveProcessingJob(ctx, w.ProcessingJob, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}


