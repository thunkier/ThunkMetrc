package models

type TransfersCreateHubDepartRequestItem struct {
    ShipmentDeliveryId int `json:"ShipmentDeliveryId,omitempty"`
    TransporterDirection string `json:"TransporterDirection,omitempty"`
}
