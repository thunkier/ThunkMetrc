package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

type ProcessingJobService struct {
	Client      *client.MetrcClient
	RateLimiter RateLimiter
}


// POST /processing/v2/adjust
// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ProcessingJobService) CreateAdjustProcessingJob(ctx context.Context, body []*models.ProcessingJobCreateAdjustProcessingJobRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateAdjustProcessingJob(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.WriteResponse: %w", err)
    }
    return typed, nil
    
}
// POST /processing/v2/jobtypes
// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ProcessingJobService) CreateProcessingJobJobTypes(ctx context.Context, body []*models.ProcessingJobCreateJobTypesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateProcessingJobJobTypes(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.WriteResponse: %w", err)
    }
    return typed, nil
    
}
// POST /processing/v2/createpackages
// Creates packages from processing jobs at a Facility, including optional location and note assignments.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ProcessingJobService) CreateProcessingJobPackages(ctx context.Context, body []*models.ProcessingJobCreateProcessingJobPackagesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.CreateProcessingJobPackages(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.WriteResponse: %w", err)
    }
    return typed, nil
    
}
// DELETE /processing/v2/jobtypes/{id}
// Archives a Processing Job Type at a Facility, making it inactive for future use.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *ProcessingJobService) DeleteProcessingJobJobTypeById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteProcessingJobJobTypeById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// DELETE /processing/v2/{id}
// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *ProcessingJobService) DeleteProcessingJobById(ctx context.Context, id string, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.DeleteProcessingJobById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /processing/v2/finish
// Completes processing jobs at a Facility by recording final notes and waste measurements.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ProcessingJobService) FinishProcessingJobProcessingJob(ctx context.Context, body []*models.ProcessingJobFinishProcessingJobRequestItem, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.FinishProcessingJobProcessingJob(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// GET /processing/v2/jobtypes/active
// Retrieves a list of all active processing job types defined within a Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *ProcessingJobService) GetProcessingJobActiveJobTypes(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.ActiveJobType], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetProcessingJobActiveJobTypes(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.ActiveJobType]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.ActiveJobType]: %w", err)
    }
    return typed, nil
    
}
// GET /processing/v2/active
// Retrieves active processing jobs at a specified Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *ProcessingJobService) GetActiveProcessingJob(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.ProcessingJob], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetActiveProcessingJob(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.ProcessingJob]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.ProcessingJob]: %w", err)
    }
    return typed, nil
    
}
// GET /processing/v2/jobtypes/inactive
// Retrieves a list of all inactive processing job types defined within a Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *ProcessingJobService) GetProcessingJobInactiveJobTypes(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.InactiveJobType], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetProcessingJobInactiveJobTypes(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.InactiveJobType]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.InactiveJobType]: %w", err)
    }
    return typed, nil
    
}
// GET /processing/v2/inactive
// Retrieves inactive processing jobs at a specified Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - lastModifiedEnd (optional): Filter by lastModifiedEnd
// - lastModifiedStart (optional): Filter by lastModifiedStart
// - licenseNumber (optional): Filter by licenseNumber
// - pageNumber (optional): Filter by pageNumber
// - pageSize (optional): Filter by pageSize
func (s *ProcessingJobService) GetInactiveProcessingJob(ctx context.Context, lastModifiedEnd string, lastModifiedStart string, licenseNumber string, pageNumber string, pageSize string) (*models.PaginatedResponse[*models.ProcessingJob], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetInactiveProcessingJob(ctx, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.ProcessingJob]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.ProcessingJob]: %w", err)
    }
    return typed, nil
    
}
// GET /processing/v2/jobtypes/attributes
// Retrieves all processing job attributes available for a Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (s *ProcessingJobService) GetProcessingJobJobTypesAttributes(ctx context.Context, licenseNumber string) (*models.PaginatedResponse[*models.JobTypesAttribute], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetProcessingJobJobTypesAttributes(ctx, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.JobTypesAttribute]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.JobTypesAttribute]: %w", err)
    }
    return typed, nil
    
}
// GET /processing/v2/jobtypes/categories
// Retrieves all processing job categories available for a specified Facility.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
func (s *ProcessingJobService) GetProcessingJobJobTypesCategories(ctx context.Context, licenseNumber string) (*models.PaginatedResponse[*models.JobTypesCategory], error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetProcessingJobJobTypesCategories(ctx, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.PaginatedResponse[*models.JobTypesCategory]
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.PaginatedResponse[*models.JobTypesCategory]: %w", err)
    }
    return typed, nil
    
}
// GET /processing/v2/{id}
// Retrieves a ProcessingJob by Id.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - id (path param)
// - licenseNumber (optional): Filter by licenseNumber
func (s *ProcessingJobService) GetProcessingJobById(ctx context.Context, id string, licenseNumber string) (*models.ProcessingJob, error) {

    res, err := s.RateLimiter.Execute(ctx, "", true, func() (interface{}, error) {
        return s.Client.GetProcessingJobById(ctx, id, licenseNumber)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.ProcessingJob
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.ProcessingJob: %w", err)
    }
    return typed, nil
    
}
// POST /processing/v2/start
// Initiates new processing jobs at a Facility, including job details and associated packages.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ProcessingJobService) StartProcessingJobProcessingJob(ctx context.Context, body []*models.ProcessingJobStartProcessingJobRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.StartProcessingJobProcessingJob(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.WriteResponse: %w", err)
    }
    return typed, nil
    
}
// PUT /processing/v2/unfinish
// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ProcessingJobService) UnfinishProcessingJobProcessingJob(ctx context.Context, body []*models.ProcessingJobUnfinishProcessingJobRequestItem, licenseNumber string) (interface{}, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UnfinishProcessingJobProcessingJob(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    return res, nil
    
}
// PUT /processing/v2/jobtypes
// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
// Permissions Required:
// - Manage Processing Job
// Parameters:
// - licenseNumber (optional): Filter by licenseNumber
// - body (request body)
        
func (s *ProcessingJobService) UpdateProcessingJobJobTypes(ctx context.Context, body []*models.ProcessingJobUpdateJobTypesRequestItem, licenseNumber string) (*models.WriteResponse, error) {

    res, err := s.RateLimiter.Execute(ctx, "", false, func() (interface{}, error) {
        return s.Client.UpdateProcessingJobJobTypes(ctx, licenseNumber, body)
    })
    if err != nil {
        return nil, err
    }

    
    var typed *models.WriteResponse
    bytes, err := json.Marshal(res)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal response for type conversion: %w", err)
    }
    if err := json.Unmarshal(bytes, &typed); err != nil {
        return nil, fmt.Errorf("failed to unmarshal to *models.WriteResponse: %w", err)
    }
    return typed, nil
    
}


