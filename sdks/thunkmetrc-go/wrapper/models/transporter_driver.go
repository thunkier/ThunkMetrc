package models

type TransporterDriver struct {
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EmployeeId string `json:"EmployeeId,omitempty"`
    FacilityId int64 `json:"FacilityId,omitempty"`
    Id int64 `json:"Id,omitempty"`
    Name string `json:"Name,omitempty"`
}
