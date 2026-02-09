package models

type ActiveJobType struct {
    Attributes []ProcessingJobActiveJobTypeAttributesItem `json:"Attributes,omitempty"`
    CategoryName string `json:"CategoryName,omitempty"`
    Description string `json:"Description,omitempty"`
    ForItems bool `json:"ForItems,omitempty"`
    ForProcessingJobs bool `json:"ForProcessingJobs,omitempty"`
    Id int64 `json:"Id,omitempty"`
    Name string `json:"Name,omitempty"`
    ProcessingSteps string `json:"ProcessingSteps,omitempty"`
    RequiresProcessingJobAttributes bool `json:"RequiresProcessingJobAttributes,omitempty"`
}
