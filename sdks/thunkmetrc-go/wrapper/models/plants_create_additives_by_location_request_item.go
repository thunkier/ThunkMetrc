package models

type PlantsCreateAdditivesByLocationRequestItem struct {
    ActiveIngredients []*PlantsCreateAdditivesByLocationRequestItemActiveIngredient `json:"ActiveIngredients,omitempty"`
    ActualDate string `json:"ActualDate,omitempty"`
    AdditiveType string `json:"AdditiveType,omitempty"`
    ApplicationDevice string `json:"ApplicationDevice,omitempty"`
    EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
    LocationName string `json:"LocationName,omitempty"`
    ProductSupplier string `json:"ProductSupplier,omitempty"`
    ProductTradeName string `json:"ProductTradeName,omitempty"`
    SublocationName string `json:"SublocationName,omitempty"`
    TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
    TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
}
