package models

type WasteMethod struct {
    ForPlants bool `json:"ForPlants,omitempty"`
    ForProductDestruction bool `json:"ForProductDestruction,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    Name string `json:"Name,omitempty"`
}
