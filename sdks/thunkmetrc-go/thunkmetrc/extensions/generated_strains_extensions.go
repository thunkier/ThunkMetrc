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



// IterateGetActiveStrains iterates over all pages of GetActiveStrains and executes handler for each item.
func IterateGetActiveStrains(ctx context.Context, s *services.StrainsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Strain) error) error {
    page := 1
    for {
        
        response, err := s.GetActiveStrains(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetActiveStrains retrieves all items updated within the specified time window.
func SyncGetActiveStrains(ctx context.Context, s *services.StrainsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Strain, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Strain) error) error {
        return IterateGetActiveStrains(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetActiveStrainsParallel syncs across multiple targets in parallel.
func SyncGetActiveStrainsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Strain, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Strain, error) {
        return SyncGetActiveStrains(ctx, w.Strains, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetInactiveStrains iterates over all pages of GetInactiveStrains and executes handler for each item.
func IterateGetInactiveStrains(ctx context.Context, s *services.StrainsService, licenseNumber string, pageSize string, handler func(*models.Strain) error) error {
    page := 1
    for {
        
        response, err := s.GetInactiveStrains(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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


