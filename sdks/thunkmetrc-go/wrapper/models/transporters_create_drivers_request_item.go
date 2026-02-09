package models

type TransportersCreateDriversRequestItem struct {
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EmployeeId string `json:"EmployeeId,omitempty"`
    Name string `json:"Name,omitempty"`
}
