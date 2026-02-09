package models

type DeliveryTransporter struct {
    ShipmentDeliveryId int64 `json:"ShipmentDeliveryId,omitempty"`
    TransporterDirection string `json:"TransporterDirection,omitempty"`
    TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
    TransporterFacilityName string `json:"TransporterFacilityName,omitempty"`
}
