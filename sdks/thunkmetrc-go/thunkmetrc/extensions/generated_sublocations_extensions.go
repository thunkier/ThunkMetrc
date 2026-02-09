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



// IterateGetActiveSublocations iterates over all pages of GetActiveSublocations and executes handler for each item.
func IterateGetActiveSublocations(ctx context.Context, s *services.SublocationsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Sublocation) error) error {
    page := 1
    for {
        
        response, err := s.GetActiveSublocations(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetActiveSublocations retrieves all items updated within the specified time window.
func SyncGetActiveSublocations(ctx context.Context, s *services.SublocationsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Sublocation, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Sublocation) error) error {
        return IterateGetActiveSublocations(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetActiveSublocationsParallel syncs across multiple targets in parallel.
func SyncGetActiveSublocationsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Sublocation, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Sublocation, error) {
        return SyncGetActiveSublocations(ctx, w.Sublocations, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetInactiveSublocations iterates over all pages of GetInactiveSublocations and executes handler for each item.
func IterateGetInactiveSublocations(ctx context.Context, s *services.SublocationsService, licenseNumber string, pageSize string, handler func(*models.Sublocation) error) error {
    page := 1
    for {
        
        response, err := s.GetInactiveSublocations(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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


