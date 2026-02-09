package models

type PlantsUpdatePlantTagRequestItem struct {
    Id int `json:"Id,omitempty"`
    Label string `json:"Label,omitempty"`
    NewTag string `json:"NewTag,omitempty"`
    ReplaceDate string `json:"ReplaceDate,omitempty"`
    TagId int `json:"TagId,omitempty"`
}
