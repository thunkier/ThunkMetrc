package extensions

import (
    "context"
    "strconv"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/services"
)



// IterateGetActivePatients iterates over all pages of GetActivePatients and executes handler for each item.
func IterateGetActivePatients(ctx context.Context, s *services.PatientsService, licenseNumber string, pageSize string, handler func(*models.Patient) error) error {
    page := 1
    for {
        
        response, err := s.GetActivePatients(ctx, licenseNumber, strconv.Itoa(page), pageSize)
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


