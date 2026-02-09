package models

type PlantBatchesCreatePlantBatchesAdditivesRequestItem struct {
    ActiveIngredients []*PlantBatchesCreatePlantBatchesAdditivesRequestItemActiveIngredient `json:"ActiveIngredients,omitempty"`
    ActualDate string `json:"ActualDate,omitempty"`
    AdditiveType string `json:"AdditiveType,omitempty"`
    ApplicationDevice string `json:"ApplicationDevice,omitempty"`
    EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
    PlantBatchName string `json:"PlantBatchName,omitempty"`
    ProductSupplier string `json:"ProductSupplier,omitempty"`
    ProductTradeName string `json:"ProductTradeName,omitempty"`
    TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
    TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
}
