package models

type Patient struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate"`
	LicenseNumber string `json:"LicenseNumber"`
	OtherFacilitiesCount int `json:"OtherFacilitiesCount"`
	PatientId int `json:"PatientId"`
	RecommendedPlants int `json:"RecommendedPlants"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity"`
	RegistrationDate string `json:"RegistrationDate"`
	CurrentPage int `json:"CurrentPage"`
	Data []PatientData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type PatientData struct {
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate"`
	LicenseNumber string `json:"LicenseNumber"`
	OtherFacilitiesCount int `json:"OtherFacilitiesCount"`
	PatientId int `json:"PatientId"`
	RecommendedPlants int `json:"RecommendedPlants"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity"`
	RegistrationDate string `json:"RegistrationDate"`
}

