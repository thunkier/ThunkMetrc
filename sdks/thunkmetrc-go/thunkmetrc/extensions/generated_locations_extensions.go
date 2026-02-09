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



// IterateGetActiveLocations iterates over all pages of GetActiveLocations and executes handler for each item.
func IterateGetActiveLocations(ctx context.Context, s *services.LocationsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.LocationsLocation) error) error {
    page := 1
    for {
        
        response, err := s.GetActiveLocations(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetActiveLocations retrieves all items updated within the specified time window.
func SyncGetActiveLocations(ctx context.Context, s *services.LocationsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.LocationsLocation, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.LocationsLocation) error) error {
        return IterateGetActiveLocations(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetActiveLocationsParallel syncs across multiple targets in parallel.
func SyncGetActiveLocationsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.LocationsLocation, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.LocationsLocation, error) {
        return SyncGetActiveLocations(ctx, w.Locations, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetInactiveLocations iterates over all pages of GetInactiveLocations and executes handler for each item.
func IterateGetInactiveLocations(ctx context.Context, s *services.LocationsService, licenseNumber string, pageSize string, handler func(*models.LocationsLocation) error) error {
    page := 1
    for {
        
        response, err := s.GetInactiveLocations(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetLocationsTypes iterates over all pages of GetLocationsTypes and executes handler for each item.
func IterateGetLocationsTypes(ctx context.Context, s *services.LocationsService, licenseNumber string, pageSize string, handler func(*models.LocationsType) error) error {
    page := 1
    for {
        
        response, err := s.GetLocationsTypes(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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


