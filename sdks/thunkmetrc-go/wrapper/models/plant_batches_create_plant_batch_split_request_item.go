package models

type PlantBatchesCreatePlantBatchSplitRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Count int `json:"Count,omitempty"`
    GroupName string `json:"GroupName,omitempty"`
    Location string `json:"Location,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PlantBatch string `json:"PlantBatch,omitempty"`
    Strain string `json:"Strain,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
}
