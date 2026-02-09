package models

type PlantsUpdatePlantMergeRequestItem struct {
    MergeDate string `json:"MergeDate,omitempty"`
    SourcePlantGroupLabel string `json:"SourcePlantGroupLabel,omitempty"`
    TargetPlantGroupLabel string `json:"TargetPlantGroupLabel,omitempty"`
}
