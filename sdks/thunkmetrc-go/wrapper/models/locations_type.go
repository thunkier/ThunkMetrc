package models

type LocationsType struct {
    ForHarvests bool `json:"ForHarvests,omitempty"`
    ForPackages bool `json:"ForPackages,omitempty"`
    ForPlantBatches bool `json:"ForPlantBatches,omitempty"`
    ForPlants bool `json:"ForPlants,omitempty"`
    Id int64 `json:"Id,omitempty"`
    Name string `json:"Name,omitempty"`
}
