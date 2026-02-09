package models

type PlantBatchesCreatePlantBatchPlantingsRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    Count int `json:"Count,omitempty"`
    Location string `json:"Location,omitempty"`
    Name string `json:"Name,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    SourcePlantBatches string `json:"SourcePlantBatches,omitempty"`
    Strain string `json:"Strain,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
    Type string `json:"Type,omitempty"`
}
