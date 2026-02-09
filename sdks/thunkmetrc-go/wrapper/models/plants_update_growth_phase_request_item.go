package models

type PlantsUpdateGrowthPhaseRequestItem struct {
    GrowthDate string `json:"GrowthDate,omitempty"`
    GrowthPhase string `json:"GrowthPhase,omitempty"`
    Id int `json:"Id,omitempty"`
    Label string `json:"Label,omitempty"`
    NewLocation string `json:"NewLocation,omitempty"`
    NewSublocation string `json:"NewSublocation,omitempty"`
    NewTag string `json:"NewTag,omitempty"`
}
