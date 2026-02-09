package models

type JobTypesCategory struct {
    ForItems bool `json:"ForItems,omitempty"`
    ForProcessingJobs bool `json:"ForProcessingJobs,omitempty"`
    Id int64 `json:"Id,omitempty"`
    Name string `json:"Name,omitempty"`
    RequiresProcessingJobAttributes bool `json:"RequiresProcessingJobAttributes,omitempty"`
}
