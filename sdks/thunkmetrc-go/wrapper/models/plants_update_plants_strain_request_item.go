package models

type PlantsUpdatePlantsStrainRequestItem struct {
    Id int `json:"Id,omitempty"`
    Label string `json:"Label,omitempty"`
    StrainId int `json:"StrainId,omitempty"`
    StrainName string `json:"StrainName,omitempty"`
}
