package models

type PlantsUpdatePlantSplitRequestItem struct {
    PlantCount int `json:"PlantCount,omitempty"`
    SourcePlantLabel string `json:"SourcePlantLabel,omitempty"`
    SplitDate string `json:"SplitDate,omitempty"`
    StrainLabel string `json:"StrainLabel,omitempty"`
    TagLabel string `json:"TagLabel,omitempty"`
}
