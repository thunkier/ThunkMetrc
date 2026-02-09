package models

type PlantBatchesCreateGrowthPhaseRequestItem struct {
    Count int `json:"Count,omitempty"`
    GrowthDate string `json:"GrowthDate,omitempty"`
    GrowthPhase string `json:"GrowthPhase,omitempty"`
    Name string `json:"Name,omitempty"`
    NewLocation string `json:"NewLocation,omitempty"`
    NewSublocation string `json:"NewSublocation,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    StartingTag string `json:"StartingTag,omitempty"`
}
