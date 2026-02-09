package models

type PlantsDeletePlantsRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Count int `json:"Count,omitempty"`
    Id int `json:"Id,omitempty"`
    Label string `json:"Label,omitempty"`
    ReasonNote string `json:"ReasonNote,omitempty"`
    WasteMaterialMixed string `json:"WasteMaterialMixed,omitempty"`
    WasteMethodName string `json:"WasteMethodName,omitempty"`
    WasteReasonName string `json:"WasteReasonName,omitempty"`
    WasteUnitOfMeasureName string `json:"WasteUnitOfMeasureName,omitempty"`
    WasteWeight float64 `json:"WasteWeight,omitempty"`
}
