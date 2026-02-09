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



// IterateGetActiveHarvests iterates over all pages of GetActiveHarvests and executes handler for each item.
func IterateGetActiveHarvests(ctx context.Context, s *services.HarvestsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Harvest) error) error {
    page := 1
    for {
        
        response, err := s.GetActiveHarvests(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetActiveHarvests retrieves all items updated within the specified time window.
func SyncGetActiveHarvests(ctx context.Context, s *services.HarvestsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Harvest, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Harvest) error) error {
        return IterateGetActiveHarvests(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetActiveHarvestsParallel syncs across multiple targets in parallel.
func SyncGetActiveHarvestsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Harvest, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Harvest, error) {
        return SyncGetActiveHarvests(ctx, w.Harvests, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetHarvestsWaste iterates over all pages of GetHarvestsWaste and executes handler for each item.
func IterateGetHarvestsWaste(ctx context.Context, s *services.HarvestsService, harvestId string, licenseNumber string, pageSize string, handler func(*models.HarvestsWaste) error) error {
    page := 1
    for {
        
        response, err := s.GetHarvestsWaste(ctx, harvestId, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetInactiveHarvests iterates over all pages of GetInactiveHarvests and executes handler for each item.
func IterateGetInactiveHarvests(ctx context.Context, s *services.HarvestsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Harvest) error) error {
    page := 1
    for {
        
        response, err := s.GetInactiveHarvests(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetInactiveHarvests retrieves all items updated within the specified time window.
func SyncGetInactiveHarvests(ctx context.Context, s *services.HarvestsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Harvest, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Harvest) error) error {
        return IterateGetInactiveHarvests(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetInactiveHarvestsParallel syncs across multiple targets in parallel.
func SyncGetInactiveHarvestsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Harvest, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Harvest, error) {
        return SyncGetInactiveHarvests(ctx, w.Harvests, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetOnHoldHarvests iterates over all pages of GetOnHoldHarvests and executes handler for each item.
func IterateGetOnHoldHarvests(ctx context.Context, s *services.HarvestsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Harvest) error) error {
    page := 1
    for {
        
        response, err := s.GetOnHoldHarvests(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetOnHoldHarvests retrieves all items updated within the specified time window.
func SyncGetOnHoldHarvests(ctx context.Context, s *services.HarvestsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Harvest, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Harvest) error) error {
        return IterateGetOnHoldHarvests(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetOnHoldHarvestsParallel syncs across multiple targets in parallel.
func SyncGetOnHoldHarvestsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Harvest, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Harvest, error) {
        return SyncGetOnHoldHarvests(ctx, w.Harvests, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetHarvestsWasteTypes iterates over all pages of GetHarvestsWasteTypes and executes handler for each item.
func IterateGetHarvestsWasteTypes(ctx context.Context, s *services.HarvestsService, pageSize string, handler func(*models.WasteType) error) error {
    page := 1
    for {
        
        response, err := s.GetHarvestsWasteTypes(ctx, strconv.Itoa(page), pageSize)
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


