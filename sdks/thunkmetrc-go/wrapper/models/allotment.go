package models

type Allotment struct {
    Allotment int64 `json:"Allotment,omitempty"`
    FacilityId int64 `json:"FacilityId,omitempty"`
    FacilityLicenseNumber string `json:"FacilityLicenseNumber,omitempty"`
}
