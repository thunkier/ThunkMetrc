package models

type ProcessingJobUpdateJobTypesRequestItem struct {
    Attributes []string `json:"Attributes,omitempty"`
    CategoryName string `json:"CategoryName,omitempty"`
    Description string `json:"Description,omitempty"`
    Id int `json:"Id,omitempty"`
    Name string `json:"Name,omitempty"`
    ProcessingSteps string `json:"ProcessingSteps,omitempty"`
}
