package models

type PlantBatchesWaste struct {
    PlantBatchId int64 `json:"PlantBatchId,omitempty"`
    PlantBatchName string `json:"PlantBatchName,omitempty"`
    PlantCount int64 `json:"PlantCount,omitempty"`
    PlantWasteNumber string `json:"PlantWasteNumber,omitempty"`
    WasteDate string `json:"WasteDate,omitempty"`
    WasteMethodName string `json:"WasteMethodName,omitempty"`
    WasteReasonName string `json:"WasteReasonName,omitempty"`
    WasteUnitOfMeasureName string `json:"WasteUnitOfMeasureName,omitempty"`
    WasteWeight float64 `json:"WasteWeight,omitempty"`
}
