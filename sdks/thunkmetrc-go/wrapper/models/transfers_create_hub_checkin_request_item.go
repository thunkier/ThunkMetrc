package models

type TransfersCreateHubCheckinRequestItem struct {
    ShipmentDeliveryId int `json:"ShipmentDeliveryId,omitempty"`
    TransporterDirection string `json:"TransporterDirection,omitempty"`
}
