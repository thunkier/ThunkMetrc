package models

type SalesCreateDeliveriesRetailerRestockRequestItem struct {
    DateTime string `json:"DateTime,omitempty"`
    Destinations string `json:"Destinations,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    Packages []*SalesCreateDeliveriesRetailerRestockRequestItemPackage `json:"Packages,omitempty"`
    RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}
