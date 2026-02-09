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



// IterateGetActiveAdditivesTemplates iterates over all pages of GetActiveAdditivesTemplates and executes handler for each item.
func IterateGetActiveAdditivesTemplates(ctx context.Context, s *services.AdditivesTemplatesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.AdditivesTemplate) error) error {
    page := 1
    for {
        
        response, err := s.GetActiveAdditivesTemplates(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetActiveAdditivesTemplates retrieves all items updated within the specified time window.
func SyncGetActiveAdditivesTemplates(ctx context.Context, s *services.AdditivesTemplatesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.AdditivesTemplate, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.AdditivesTemplate) error) error {
        return IterateGetActiveAdditivesTemplates(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetActiveAdditivesTemplatesParallel syncs across multiple targets in parallel.
func SyncGetActiveAdditivesTemplatesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.AdditivesTemplate, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.AdditivesTemplate, error) {
        return SyncGetActiveAdditivesTemplates(ctx, w.AdditivesTemplates, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetInactiveAdditivesTemplates iterates over all pages of GetInactiveAdditivesTemplates and executes handler for each item.
func IterateGetInactiveAdditivesTemplates(ctx context.Context, s *services.AdditivesTemplatesService, licenseNumber string, pageSize string, handler func(*models.AdditivesTemplate) error) error {
    page := 1
    for {
        
        response, err := s.GetInactiveAdditivesTemplates(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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


