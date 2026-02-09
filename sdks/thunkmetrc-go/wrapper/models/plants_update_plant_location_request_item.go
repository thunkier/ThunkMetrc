package models

type PlantsUpdatePlantLocationRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Id int `json:"Id,omitempty"`
    Label string `json:"Label,omitempty"`
    Location string `json:"Location,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
}
