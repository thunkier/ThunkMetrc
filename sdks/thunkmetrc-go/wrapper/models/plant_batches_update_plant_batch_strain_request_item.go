package models

type PlantBatchesUpdatePlantBatchStrainRequestItem struct {
    Id int `json:"Id,omitempty"`
    Name string `json:"Name,omitempty"`
    StrainId int `json:"StrainId,omitempty"`
    StrainName string `json:"StrainName,omitempty"`
}
