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



// IterateGetPlantsAdditives iterates over all pages of GetPlantsAdditives and executes handler for each item.
func IterateGetPlantsAdditives(ctx context.Context, s *services.PlantsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Additive) error) error {
    page := 1
    for {
        
        response, err := s.GetPlantsAdditives(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetPlantsAdditives retrieves all items updated within the specified time window.
func SyncGetPlantsAdditives(ctx context.Context, s *services.PlantsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Additive, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Additive) error) error {
        return IterateGetPlantsAdditives(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetPlantsAdditivesParallel syncs across multiple targets in parallel.
func SyncGetPlantsAdditivesParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Additive, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Additive, error) {
        return SyncGetPlantsAdditives(ctx, w.Plants, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetFloweringPlants iterates over all pages of GetFloweringPlants and executes handler for each item.
func IterateGetFloweringPlants(ctx context.Context, s *services.PlantsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Plant) error) error {
    page := 1
    for {
        
        response, err := s.GetFloweringPlants(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetFloweringPlants retrieves all items updated within the specified time window.
func SyncGetFloweringPlants(ctx context.Context, s *services.PlantsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Plant, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Plant) error) error {
        return IterateGetFloweringPlants(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetFloweringPlantsParallel syncs across multiple targets in parallel.
func SyncGetFloweringPlantsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Plant, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Plant, error) {
        return SyncGetFloweringPlants(ctx, w.Plants, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetInactivePlants iterates over all pages of GetInactivePlants and executes handler for each item.
func IterateGetInactivePlants(ctx context.Context, s *services.PlantsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Plant) error) error {
    page := 1
    for {
        
        response, err := s.GetInactivePlants(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetInactivePlants retrieves all items updated within the specified time window.
func SyncGetInactivePlants(ctx context.Context, s *services.PlantsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Plant, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Plant) error) error {
        return IterateGetInactivePlants(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetInactivePlantsParallel syncs across multiple targets in parallel.
func SyncGetInactivePlantsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Plant, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Plant, error) {
        return SyncGetInactivePlants(ctx, w.Plants, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetMotherInactivePlants iterates over all pages of GetMotherInactivePlants and executes handler for each item.
func IterateGetMotherInactivePlants(ctx context.Context, s *services.PlantsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Mother) error) error {
    page := 1
    for {
        
        response, err := s.GetMotherInactivePlants(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetMotherInactivePlants retrieves all items updated within the specified time window.
func SyncGetMotherInactivePlants(ctx context.Context, s *services.PlantsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Mother, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Mother) error) error {
        return IterateGetMotherInactivePlants(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetMotherInactivePlantsParallel syncs across multiple targets in parallel.
func SyncGetMotherInactivePlantsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Mother, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Mother, error) {
        return SyncGetMotherInactivePlants(ctx, w.Plants, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetMotherOnHoldPlants iterates over all pages of GetMotherOnHoldPlants and executes handler for each item.
func IterateGetMotherOnHoldPlants(ctx context.Context, s *services.PlantsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Mother) error) error {
    page := 1
    for {
        
        response, err := s.GetMotherOnHoldPlants(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetMotherOnHoldPlants retrieves all items updated within the specified time window.
func SyncGetMotherOnHoldPlants(ctx context.Context, s *services.PlantsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Mother, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Mother) error) error {
        return IterateGetMotherOnHoldPlants(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetMotherOnHoldPlantsParallel syncs across multiple targets in parallel.
func SyncGetMotherOnHoldPlantsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Mother, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Mother, error) {
        return SyncGetMotherOnHoldPlants(ctx, w.Plants, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetMotherPlants iterates over all pages of GetMotherPlants and executes handler for each item.
func IterateGetMotherPlants(ctx context.Context, s *services.PlantsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Mother) error) error {
    page := 1
    for {
        
        response, err := s.GetMotherPlants(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetMotherPlants retrieves all items updated within the specified time window.
func SyncGetMotherPlants(ctx context.Context, s *services.PlantsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Mother, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Mother) error) error {
        return IterateGetMotherPlants(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetMotherPlantsParallel syncs across multiple targets in parallel.
func SyncGetMotherPlantsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Mother, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Mother, error) {
        return SyncGetMotherPlants(ctx, w.Plants, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetOnHoldPlants iterates over all pages of GetOnHoldPlants and executes handler for each item.
func IterateGetOnHoldPlants(ctx context.Context, s *services.PlantsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Plant) error) error {
    page := 1
    for {
        
        response, err := s.GetOnHoldPlants(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetOnHoldPlants retrieves all items updated within the specified time window.
func SyncGetOnHoldPlants(ctx context.Context, s *services.PlantsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Plant, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Plant) error) error {
        return IterateGetOnHoldPlants(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetOnHoldPlantsParallel syncs across multiple targets in parallel.
func SyncGetOnHoldPlantsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Plant, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Plant, error) {
        return SyncGetOnHoldPlants(ctx, w.Plants, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetPlantsWaste iterates over all pages of GetPlantsWaste and executes handler for each item.
func IterateGetPlantsWaste(ctx context.Context, s *services.PlantsService, licenseNumber string, pageSize string, handler func(*models.PlantsWaste) error) error {
    page := 1
    for {
        
        response, err := s.GetPlantsWaste(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetPlantsWasteMethods iterates over all pages of GetPlantsWasteMethods and executes handler for each item.
func IterateGetPlantsWasteMethods(ctx context.Context, s *services.PlantsService, pageSize string, handler func(*models.WasteMethod) error) error {
    page := 1
    for {
        
        response, err := s.GetPlantsWasteMethods(ctx, strconv.Itoa(page), pageSize)
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

// IterateGetPlantsWasteReasons iterates over all pages of GetPlantsWasteReasons and executes handler for each item.
func IterateGetPlantsWasteReasons(ctx context.Context, s *services.PlantsService, licenseNumber string, pageSize string, handler func(*models.WasteReason) error) error {
    page := 1
    for {
        
        response, err := s.GetPlantsWasteReasons(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetVegetativePlants iterates over all pages of GetVegetativePlants and executes handler for each item.
func IterateGetVegetativePlants(ctx context.Context, s *services.PlantsService, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageSize string, handler func(*models.Plant) error) error {
    page := 1
    for {
        
        response, err := s.GetVegetativePlants(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, strconv.Itoa(page), pageSize)
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

// SyncGetVegetativePlants retrieves all items updated within the specified time window.
func SyncGetVegetativePlants(ctx context.Context, s *services.PlantsService, lastKnownSync *time.Time, bufferMinutes int,
    licenseNumber string,
    pageSize string,
) ([]*models.Plant, error) {
    window := utils.GetTimeWindow(lastKnownSync, bufferMinutes)
    
    return utils.Collect(ctx, func(handler func(*models.Plant) error) error {
        return IterateGetVegetativePlants(ctx, s,window.EndString,window.StartString,
            licenseNumber,
            pageSize,
            handler,
        )
    })
}

// SyncGetVegetativePlantsParallel syncs across multiple targets in parallel.
func SyncGetVegetativePlantsParallel(ctx context.Context, targets []utils.SynchronizationTarget, lastKnownSync *time.Time, bufferMinutes int, concurrencyLimit int,
    pageSize string,
) (map[string][]*models.Plant, error) {
    return utils.ParallelSync(ctx, targets, concurrencyLimit, func(ctx context.Context, w *wrapper.MetrcWrapper, licenseNumber string) ([]*models.Plant, error) {
        return SyncGetVegetativePlants(ctx, w.Plants, lastKnownSync, bufferMinutes,licenseNumber,
            pageSize,
        )
    })
}

// IterateGetPlantsWasteById iterates over all pages of GetPlantsWasteById and executes handler for each item.
func IterateGetPlantsWasteById(ctx context.Context, s *services.PlantsService, id string, licenseNumber string, pageSize string, handler func(*models.PlantsWaste) error) error {
    page := 1
    for {
        
        response, err := s.GetPlantsWasteById(ctx, id, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetPlantsWastePackageById iterates over all pages of GetPlantsWastePackageById and executes handler for each item.
func IterateGetPlantsWastePackageById(ctx context.Context, s *services.PlantsService, id string, licenseNumber string, pageSize string, handler func(*models.WastePackage) error) error {
    page := 1
    for {
        
        response, err := s.GetPlantsWastePackageById(ctx, id, licenseNumber, strconv.Itoa(page), pageSize)
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


