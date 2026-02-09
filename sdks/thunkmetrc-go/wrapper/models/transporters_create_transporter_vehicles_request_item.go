package models

type TransportersCreateTransporterVehiclesRequestItem struct {
    LicensePlateNumber string `json:"LicensePlateNumber,omitempty"`
    Make string `json:"Make,omitempty"`
    Model string `json:"Model,omitempty"`
}
