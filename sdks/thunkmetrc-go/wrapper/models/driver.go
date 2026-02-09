package models

type Driver struct {
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EmployeeId string `json:"EmployeeId,omitempty"`
    FacilityId int64 `json:"FacilityId,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsArchived bool `json:"IsArchived,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    Name string `json:"Name,omitempty"`
}
