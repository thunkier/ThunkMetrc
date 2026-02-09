package models

type PlantBatchesUpdateNameRequestItem struct {
    Group string `json:"Group,omitempty"`
    Id int `json:"Id,omitempty"`
    NewGroup string `json:"NewGroup,omitempty"`
}
