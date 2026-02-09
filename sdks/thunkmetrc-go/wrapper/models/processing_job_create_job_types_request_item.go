package models

type ProcessingJobCreateJobTypesRequestItem struct {
    Attributes []string `json:"Attributes,omitempty"`
    Category string `json:"Category,omitempty"`
    Description string `json:"Description,omitempty"`
    Name string `json:"Name,omitempty"`
    ProcessingSteps string `json:"ProcessingSteps,omitempty"`
}
