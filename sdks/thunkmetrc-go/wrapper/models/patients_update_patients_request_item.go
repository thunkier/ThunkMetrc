package models

type PatientsUpdatePatientsRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
    FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
    HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
    InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
    LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
    LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
    LicenseNumber string `json:"LicenseNumber,omitempty"`
    MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
    MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
    NewLicenseNumber string `json:"NewLicenseNumber,omitempty"`
    RecommendedPlants int `json:"RecommendedPlants,omitempty"`
    RecommendedSmokableQuantity float64 `json:"RecommendedSmokableQuantity,omitempty"`
    ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
}
