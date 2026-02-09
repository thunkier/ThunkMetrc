package models

type PlantsCreatePlantAdditivesUsingTemplateRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
    PlantLabels []string `json:"PlantLabels,omitempty"`
    Rate string `json:"Rate,omitempty"`
    TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
    TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
    Volume string `json:"Volume,omitempty"`
}
