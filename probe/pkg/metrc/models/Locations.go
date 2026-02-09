package models

type LocationsLocation struct {
	ForHarvests bool `json:"ForHarvests"`
	ForPackages bool `json:"ForPackages"`
	ForPlantBatches bool `json:"ForPlantBatches"`
	ForPlants bool `json:"ForPlants"`
	Id int `json:"Id"`
	LocationTypeId int `json:"LocationTypeId"`
	LocationTypeName string `json:"LocationTypeName"`
	Name string `json:"Name"`
}

type LocationsRequest struct {
	LocationTypeName string `json:"LocationTypeName"`
	Name string `json:"Name"`
}

type LocationsType struct {
	ForHarvests bool `json:"ForHarvests"`
	ForPackages bool `json:"ForPackages"`
	ForPlantBatches bool `json:"ForPlantBatches"`
	ForPlants bool `json:"ForPlants"`
	Id int `json:"Id"`
	Name string `json:"Name"`
}
