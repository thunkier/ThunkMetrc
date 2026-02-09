package models

type FacilityLicense struct {
    EmployeeLicenseNumber string `json:"EmployeeLicenseNumber,omitempty"`
    EndDate string `json:"EndDate,omitempty"`
    LicenseType string `json:"LicenseType,omitempty"`
    Number string `json:"Number,omitempty"`
    StartDate string `json:"StartDate,omitempty"`
}
