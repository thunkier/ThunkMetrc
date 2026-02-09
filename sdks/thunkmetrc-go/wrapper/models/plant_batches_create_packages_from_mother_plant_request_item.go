package models

type PlantBatchesCreatePackagesFromMotherPlantRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Count int `json:"Count,omitempty"`
    ExpirationDate string `json:"ExpirationDate,omitempty"`
    Id int `json:"Id,omitempty"`
    IsDonation bool `json:"IsDonation,omitempty"`
    IsTradeSample bool `json:"IsTradeSample,omitempty"`
    Item string `json:"Item,omitempty"`
    Location string `json:"Location,omitempty"`
    Note string `json:"Note,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PlantBatch string `json:"PlantBatch,omitempty"`
    SellByDate string `json:"SellByDate,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
    Tag string `json:"Tag,omitempty"`
    UseByDate string `json:"UseByDate,omitempty"`
}
