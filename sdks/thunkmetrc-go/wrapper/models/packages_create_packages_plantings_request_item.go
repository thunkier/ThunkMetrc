package models

type PackagesCreatePackagesPlantingsRequestItem struct {
    LocationName string `json:"LocationName,omitempty"`
    PackageAdjustmentAmount int `json:"PackageAdjustmentAmount,omitempty"`
    PackageAdjustmentUnitOfMeasureName string `json:"PackageAdjustmentUnitOfMeasureName,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PlantBatchName string `json:"PlantBatchName,omitempty"`
    PlantBatchType string `json:"PlantBatchType,omitempty"`
    PlantCount int `json:"PlantCount,omitempty"`
    PlantedDate string `json:"PlantedDate,omitempty"`
    StrainName string `json:"StrainName,omitempty"`
    SublocationName string `json:"SublocationName,omitempty"`
    UnpackagedDate string `json:"UnpackagedDate,omitempty"`
}
