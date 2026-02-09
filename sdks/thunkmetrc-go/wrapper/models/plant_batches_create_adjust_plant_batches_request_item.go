package models

type PlantBatchesCreateAdjustPlantBatchesRequestItem struct {
    AdjustmentDate string `json:"AdjustmentDate,omitempty"`
    AdjustmentReason string `json:"AdjustmentReason,omitempty"`
    PlantBatchName string `json:"PlantBatchName,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
    ReasonNote string `json:"ReasonNote,omitempty"`
}
