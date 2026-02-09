package models

type PlantsCreateAdditivesByLocationUsingTemplateRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
    LocationName string `json:"LocationName,omitempty"`
    Rate string `json:"Rate,omitempty"`
    SublocationName string `json:"SublocationName,omitempty"`
    TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
    TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
    Volume string `json:"Volume,omitempty"`
}
