package models

type TransportersUpdateVehiclesRequestItem struct {
    Id string `json:"Id,omitempty"`
    LicensePlateNumber string `json:"LicensePlateNumber,omitempty"`
    Make string `json:"Make,omitempty"`
    Model string `json:"Model,omitempty"`
}
