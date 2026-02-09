package models

type Location struct {
    ForHarvests bool `json:"ForHarvests,omitempty"`
    ForPackages bool `json:"ForPackages,omitempty"`
    ForPlantBatches bool `json:"ForPlantBatches,omitempty"`
    ForPlants bool `json:"ForPlants,omitempty"`
    Id int64 `json:"Id,omitempty"`
    LocationTypeId int64 `json:"LocationTypeId,omitempty"`
    LocationTypeName string `json:"LocationTypeName,omitempty"`
    Name string `json:"Name,omitempty"`
}
