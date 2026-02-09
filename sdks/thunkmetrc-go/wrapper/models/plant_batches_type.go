package models

type PlantBatchesType struct {
    CanBeCloned bool `json:"CanBeCloned,omitempty"`
    Id int64 `json:"Id,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    Name string `json:"Name,omitempty"`
}
