package models

type SalesCreateSaleDeliveriesRetailerRestockRequestItem struct {
    DateTime string `json:"DateTime,omitempty"`
    Destinations string `json:"Destinations,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    Packages []*SalesCreateSaleDeliveriesRetailerRestockRequestItemPackage `json:"Packages,omitempty"`
    RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}
