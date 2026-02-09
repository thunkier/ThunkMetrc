package models

type PlantBatchesCreatePlantBatchWasteRequestItem struct {
    MixedMaterial string `json:"MixedMaterial,omitempty"`
    Note string `json:"Note,omitempty"`
    PlantBatchName string `json:"PlantBatchName,omitempty"`
    ReasonName string `json:"ReasonName,omitempty"`
    UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
    WasteDate string `json:"WasteDate,omitempty"`
    WasteMethodName string `json:"WasteMethodName,omitempty"`
    WasteWeight float64 `json:"WasteWeight,omitempty"`
}
