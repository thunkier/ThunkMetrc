package models

type TransportersUpdateDriversRequestItem struct {
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EmployeeId string `json:"EmployeeId,omitempty"`
    Id string `json:"Id,omitempty"`
    Name string `json:"Name,omitempty"`
}
