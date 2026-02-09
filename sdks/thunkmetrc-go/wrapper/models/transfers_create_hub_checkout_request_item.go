package models

type TransfersCreateHubCheckoutRequestItem struct {
    ShipmentDeliveryId int `json:"ShipmentDeliveryId,omitempty"`
    TransporterDirection string `json:"TransporterDirection,omitempty"`
}
