package models

type Location struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
	ForHarvests bool `json:"ForHarvests"`
	ForPackages bool `json:"ForPackages"`
	ForPlantBatches bool `json:"ForPlantBatches"`
	ForPlants bool `json:"ForPlants"`
	Id int `json:"Id"`
	LocationTypeId int `json:"LocationTypeId"`
	LocationTypeName string `json:"LocationTypeName"`
	Name string `json:"Name"`
	CurrentPage int `json:"CurrentPage"`
	Data []LocationData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type LocationData struct {
	ForHarvests bool `json:"ForHarvests"`
	ForPackages bool `json:"ForPackages"`
	ForPlantBatches bool `json:"ForPlantBatches"`
	ForPlants bool `json:"ForPlants"`
	Id int `json:"Id"`
	LocationTypeId int `json:"LocationTypeId"`
	LocationTypeName string `json:"LocationTypeName"`
	Name string `json:"Name"`
}

