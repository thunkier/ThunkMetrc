package models

type Employee struct {
	FullName string `json:"FullName"`
	IsIndustryAdmin bool `json:"IsIndustryAdmin"`
	IsManager bool `json:"IsManager"`
	IsOwner bool `json:"IsOwner"`
	License EmployeeLicense `json:"License"`
}

type EmployeeLicense struct {
	EmployeeLicenseNumber string `json:"EmployeeLicenseNumber"`
	EndDate string `json:"EndDate"`
	LicenseType string `json:"LicenseType"`
	Number string `json:"Number"`
	StartDate string `json:"StartDate"`
}
