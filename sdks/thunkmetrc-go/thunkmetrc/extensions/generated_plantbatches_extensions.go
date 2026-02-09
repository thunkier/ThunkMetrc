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



// IterateGetActivePlantBatches iterates over all pages of GetActivePlantBatches and executes handler for each item.
func IterateGetActivePlantBatches(ctx context.Context, s *services.PlantBatchesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.PlantBatch) error) error {
    page := 1
    for {
        
        response, err := s.GetActivePlantBatches(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetActivePlantBatches retrieves all items updated within the specified time window.
func SyncGetActivePlantBatches(ctx context.Context, s *services.PlantBatchesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.PlantBatch, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.PlantBatch) error) error {
        return IterateGetActivePlantBatches(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetActivePlantBatchesParallel syncs across multiple targets in parallel.
func SyncGetActivePlantBatchesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.PlantBatch, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.PlantBatch, error) {
        return SyncGetActivePlantBatches(ctx, w.PlantBatches, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetInactivePlantBatches iterates over all pages of GetInactivePlantBatches and executes handler for each item.
func IterateGetInactivePlantBatches(ctx context.Context, s *services.PlantBatchesService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.PlantBatch) error) error {
    page := 1
    for {
        
        response, err := s.GetInactivePlantBatches(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetInactivePlantBatches retrieves all items updated within the specified time window.
func SyncGetInactivePlantBatches(ctx context.Context, s *services.PlantBatchesService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.PlantBatch, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.PlantBatch) error) error {
        return IterateGetInactivePlantBatches(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetInactivePlantBatchesParallel syncs across multiple targets in parallel.
func SyncGetInactivePlantBatchesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.PlantBatch, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.PlantBatch, error) {
        return SyncGetInactivePlantBatches(ctx, w.PlantBatches, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetPlantBatchesTypes iterates over all pages of GetPlantBatchesTypes and executes handler for each item.
func IterateGetPlantBatchesTypes(ctx context.Context, s *services.PlantBatchesService, pageSize string, handler func(*models.PlantBatchesType) error) error {
    page := 1
    for {
        
        response, err := s.GetPlantBatchesTypes(ctx, strconv.Itoa(page), pageSize)
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

// IterateGetPlantBatchesWaste iterates over all pages of GetPlantBatchesWaste and executes handler for each item.
func IterateGetPlantBatchesWaste(ctx context.Context, s *services.PlantBatchesService, licenseNumber string, pageSize string, handler func(*models.PlantBatchesWaste) error) error {
    page := 1
    for {
        
        response, err := s.GetPlantBatchesWaste(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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


