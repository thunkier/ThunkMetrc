package models

type Patient struct {
    HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
    LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
    LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
    LicenseNumber string `json:"LicenseNumber,omitempty"`
    OtherFacilitiesCount int64 `json:"OtherFacilitiesCount,omitempty"`
    PatientId int64 `json:"PatientId,omitempty"`
    RecommendedPlants int64 `json:"RecommendedPlants,omitempty"`
    RecommendedSmokableQuantity float64 `json:"RecommendedSmokableQuantity,omitempty"`
    RegistrationDate string `json:"RegistrationDate,omitempty"`
}
