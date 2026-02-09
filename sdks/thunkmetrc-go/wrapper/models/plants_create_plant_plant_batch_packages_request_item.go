package models

type PlantsCreatePlantPlantBatchPackagesRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Count int `json:"Count,omitempty"`
    IsDonation bool `json:"IsDonation,omitempty"`
    IsTradeSample bool `json:"IsTradeSample,omitempty"`
    Item string `json:"Item,omitempty"`
    Location string `json:"Location,omitempty"`
    Note string `json:"Note,omitempty"`
    PackageTag string `json:"PackageTag,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PlantBatchType string `json:"PlantBatchType,omitempty"`
    PlantLabel string `json:"PlantLabel,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
}
