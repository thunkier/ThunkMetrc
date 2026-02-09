package models

type PackagesCreateAdjustPackagesRequestItem struct {
    AdjustmentDate string `json:"AdjustmentDate,omitempty"`
    AdjustmentReason string `json:"AdjustmentReason,omitempty"`
    Label string `json:"Label,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
    ReasonNote string `json:"ReasonNote,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}
