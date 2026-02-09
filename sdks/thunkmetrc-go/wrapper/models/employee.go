package models

type Employee struct {
    FullName string `json:"FullName,omitempty"`
    IsIndustryAdmin bool `json:"IsIndustryAdmin,omitempty"`
    IsManager bool `json:"IsManager,omitempty"`
    IsOwner bool `json:"IsOwner,omitempty"`
    License EmployeeLicense `json:"License,omitempty"`
}
