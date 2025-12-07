package models

type WasteMethod struct {
	ForPlants bool `json:"ForPlants"`
	ForProductDestruction bool `json:"ForProductDestruction"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
}

