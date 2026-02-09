package models

type TransferDeliveryTransporter struct {
    TransporterDirection string `json:"TransporterDirection,omitempty"`
    TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
    TransporterFacilityName string `json:"TransporterFacilityName,omitempty"`
}
