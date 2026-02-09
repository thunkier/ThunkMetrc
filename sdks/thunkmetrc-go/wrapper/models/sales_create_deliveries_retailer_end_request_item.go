package models

type SalesCreateDeliveriesRetailerEndRequestItem struct {
    ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
    Packages []*SalesCreateDeliveriesRetailerEndRequestItemPackage `json:"Packages,omitempty"`
    RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}
