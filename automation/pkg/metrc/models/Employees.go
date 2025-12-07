package models

type Employee struct {
	FullName string `json:"FullName"`
	IsIndustryAdmin bool `json:"IsIndustryAdmin"`
	IsManager bool `json:"IsManager"`
	IsOwner bool `json:"IsOwner"`
	License EmployeeLicense `json:"License"`
	CurrentPage int `json:"CurrentPage"`
	Data []EmployeeData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type EmployeeData struct {
	FullName string `json:"FullName"`
	IsIndustryAdmin bool `json:"IsIndustryAdmin"`
	IsManager bool `json:"IsManager"`
	IsOwner bool `json:"IsOwner"`
	License EmployeeDataLicense `json:"License"`
}

type EmployeeDataLicense struct {
	EmployeeLicenseNumber string `json:"EmployeeLicenseNumber"`
	EndDate string `json:"EndDate"`
	LicenseType string `json:"LicenseType"`
	Number string `json:"Number"`
	StartDate string `json:"StartDate"`
}

type EmployeeLicense struct {
	EmployeeLicenseNumber string `json:"EmployeeLicenseNumber"`
	EndDate string `json:"EndDate"`
	LicenseType string `json:"LicenseType"`
	Number string `json:"Number"`
	StartDate string `json:"StartDate"`
}

