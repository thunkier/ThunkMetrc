package models

type TransfersCreateHubArriveRequestItem struct {
    ShipmentDeliveryId int `json:"ShipmentDeliveryId,omitempty"`
    TransporterDirection string `json:"TransporterDirection,omitempty"`
}
