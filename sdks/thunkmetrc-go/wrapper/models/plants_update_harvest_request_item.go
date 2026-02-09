package models

type PlantsUpdateHarvestRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    DryingLocation string `json:"DryingLocation,omitempty"`
    DryingSublocation string `json:"DryingSublocation,omitempty"`
    HarvestName string `json:"HarvestName,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    Plant string `json:"Plant,omitempty"`
    UnitOfWeight string `json:"UnitOfWeight,omitempty"`
    Weight float64 `json:"Weight,omitempty"`
}
