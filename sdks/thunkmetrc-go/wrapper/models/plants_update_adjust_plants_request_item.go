package models

type PlantsUpdateAdjustPlantsRequestItem struct {
    AdjustCount int `json:"AdjustCount,omitempty"`
    AdjustReason string `json:"AdjustReason,omitempty"`
    AdjustmentDate string `json:"AdjustmentDate,omitempty"`
    Id int `json:"Id,omitempty"`
    Label string `json:"Label,omitempty"`
    ReasonNote string `json:"ReasonNote,omitempty"`
}
