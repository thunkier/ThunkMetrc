package models

type PlantBatchesDeletePlantBatchesRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Count int `json:"Count,omitempty"`
    PlantBatch string `json:"PlantBatch,omitempty"`
    ReasonNote string `json:"ReasonNote,omitempty"`
    WasteMaterialMixed string `json:"WasteMaterialMixed,omitempty"`
    WasteMethodName string `json:"WasteMethodName,omitempty"`
    WasteReasonName string `json:"WasteReasonName,omitempty"`
    WasteUnitOfMeasure string `json:"WasteUnitOfMeasure,omitempty"`
    WasteWeight float64 `json:"WasteWeight,omitempty"`
}
