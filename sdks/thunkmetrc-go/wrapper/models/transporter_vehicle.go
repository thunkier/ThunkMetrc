package models

type TransporterVehicle struct {
    FacilityId int64 `json:"FacilityId,omitempty"`
    Id int64 `json:"Id,omitempty"`
    LicensePlateNumber string `json:"LicensePlateNumber,omitempty"`
    Make string `json:"Make,omitempty"`
    Model string `json:"Model,omitempty"`
}
