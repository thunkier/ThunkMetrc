package extensions

import (
    "context"
    "strconv"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/services"
)



// IterateGetTransportersDrivers iterates over all pages of GetTransportersDrivers and executes handler for each item.
func IterateGetTransportersDrivers(ctx context.Context, s *services.TransportersService, licenseNumber string, pageSize string, handler func(*models.TransportersDriver) error) error {
    page := 1
    for {
        
        response, err := s.GetTransportersDrivers(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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

// IterateGetTransportersVehicles iterates over all pages of GetTransportersVehicles and executes handler for each item.
func IterateGetTransportersVehicles(ctx context.Context, s *services.TransportersService, licenseNumber string, pageSize string, handler func(*models.TransportersVehicle) error) error {
    page := 1
    for {
        
        response, err := s.GetTransportersVehicles(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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


