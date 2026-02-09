package models

type PlantBatchesUpdatePlantBatchesTagRequestItem struct {
    Group string `json:"Group,omitempty"`
    Id int `json:"Id,omitempty"`
    NewTag string `json:"NewTag,omitempty"`
    ReplaceDate string `json:"ReplaceDate,omitempty"`
    TagId int `json:"TagId,omitempty"`
}
