package models

type HarvestWaste struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Id int64 `json:"Id,omitempty"`
    UnitOfWeightName string `json:"UnitOfWeightName,omitempty"`
    WasteTypeName string `json:"WasteTypeName,omitempty"`
    WasteWeight float64 `json:"WasteWeight,omitempty"`
}
