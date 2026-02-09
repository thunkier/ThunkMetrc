package models

type PlantsCreatePlantsWasteRequestItem struct {
    LocationName string `json:"LocationName,omitempty"`
    MixedMaterial string `json:"MixedMaterial,omitempty"`
    Note string `json:"Note,omitempty"`
    PlantLabels []interface{} `json:"PlantLabels,omitempty"`
    ReasonName string `json:"ReasonName,omitempty"`
    SublocationName string `json:"SublocationName,omitempty"`
    UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
    WasteDate string `json:"WasteDate,omitempty"`
    WasteMethodName string `json:"WasteMethodName,omitempty"`
    WasteWeight float64 `json:"WasteWeight,omitempty"`
}
