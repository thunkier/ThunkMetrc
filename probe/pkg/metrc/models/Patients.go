package models

type PatientsRequest struct {
	ActualDate string `json:"ActualDate"`
	ConcentrateOuncesAllowed *int `json:"ConcentrateOuncesAllowed"`
	FlowerOuncesAllowed *int `json:"FlowerOuncesAllowed"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption"`
	InfusedOuncesAllowed *int `json:"InfusedOuncesAllowed"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate"`
	LicenseNumber string `json:"LicenseNumber"`
	MaxConcentrateThcPercentAllowed *int `json:"MaxConcentrateThcPercentAllowed"`
	MaxFlowerThcPercentAllowed *int `json:"MaxFlowerThcPercentAllowed"`
	NewLicenseNumber *string `json:"NewLicenseNumber"`
	RecommendedPlants int `json:"RecommendedPlants"`
	RecommendedSmokableQuantity float64 `json:"RecommendedSmokableQuantity"`
	ThcOuncesAllowed *int `json:"ThcOuncesAllowed"`
}
