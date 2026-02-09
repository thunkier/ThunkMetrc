package models

type PlantsUpdateSplitRequestItem struct {
    PlantCount int `json:"PlantCount,omitempty"`
    SourcePlantLabel string `json:"SourcePlantLabel,omitempty"`
    SplitDate string `json:"SplitDate,omitempty"`
    StrainLabel string `json:"StrainLabel,omitempty"`
    TagLabel string `json:"TagLabel,omitempty"`
}
