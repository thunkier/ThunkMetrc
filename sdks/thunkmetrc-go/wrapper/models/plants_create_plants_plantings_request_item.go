package models

type PlantsCreatePlantsPlantingsRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    LocationName string `json:"LocationName,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PlantBatchName string `json:"PlantBatchName,omitempty"`
    PlantBatchType string `json:"PlantBatchType,omitempty"`
    PlantCount int `json:"PlantCount,omitempty"`
    PlantLabel string `json:"PlantLabel,omitempty"`
    StrainName string `json:"StrainName,omitempty"`
    SublocationName string `json:"SublocationName,omitempty"`
}
