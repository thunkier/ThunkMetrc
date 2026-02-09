package models

type TransfersCreateTransferHubCheckinRequestItem struct {
    ShipmentDeliveryId int `json:"ShipmentDeliveryId,omitempty"`
    TransporterDirection string `json:"TransporterDirection,omitempty"`
}
