package models

type Vehicle struct {
    FacilityId int64 `json:"FacilityId,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsArchived bool `json:"IsArchived,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    LicensePlateNumber string `json:"LicensePlateNumber,omitempty"`
    Make string `json:"Make,omitempty"`
    Model string `json:"Model,omitempty"`
}
