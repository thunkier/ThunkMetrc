package extensions

import (
    "context"
    "strconv"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/services"
)



// IterateGetLabTestsBatches iterates over all pages of GetLabTestsBatches and executes handler for each item.
func IterateGetLabTestsBatches(ctx context.Context, s *services.LabTestsService, pageSize string, handler func(*models.Batch) error) error {
    page := 1
    for {
        
        response, err := s.GetLabTestsBatches(ctx, strconv.Itoa(page), pageSize)
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

// IterateGetLabTestsTypes iterates over all pages of GetLabTestsTypes and executes handler for each item.
func IterateGetLabTestsTypes(ctx context.Context, s *services.LabTestsService, pageSize string, handler func(*models.LabTestsType) error) error {
    page := 1
    for {
        
        response, err := s.GetLabTestsTypes(ctx, strconv.Itoa(page), pageSize)
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

// IterateGetLabTestsResults iterates over all pages of GetLabTestsResults and executes handler for each item.
func IterateGetLabTestsResults(ctx context.Context, s *services.LabTestsService, licenseNumber string, packageId string, pageSize string, handler func(*models.Result) error) error {
    page := 1
    for {
        
        response, err := s.GetLabTestsResults(ctx, licenseNumber, packageId, strconv.Itoa(page), pageSize)
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


