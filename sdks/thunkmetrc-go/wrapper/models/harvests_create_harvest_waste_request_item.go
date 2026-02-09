package models

type HarvestsCreateHarvestWasteRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Id int `json:"Id,omitempty"`
    UnitOfWeight string `json:"UnitOfWeight,omitempty"`
    WasteType string `json:"WasteType,omitempty"`
    WasteWeight float64 `json:"WasteWeight,omitempty"`
}
