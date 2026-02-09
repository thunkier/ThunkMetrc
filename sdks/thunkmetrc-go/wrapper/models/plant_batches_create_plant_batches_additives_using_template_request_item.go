package models

type PlantBatchesCreatePlantBatchesAdditivesUsingTemplateRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
    PlantBatchName string `json:"PlantBatchName,omitempty"`
    Rate string `json:"Rate,omitempty"`
    TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
    TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
    Volume string `json:"Volume,omitempty"`
}
