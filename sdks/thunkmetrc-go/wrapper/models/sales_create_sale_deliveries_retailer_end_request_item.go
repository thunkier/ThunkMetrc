package models

type SalesCreateSaleDeliveriesRetailerEndRequestItem struct {
    ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
    Packages []*SalesCreateSaleDeliveriesRetailerEndRequestItemPackage `json:"Packages,omitempty"`
    RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}
