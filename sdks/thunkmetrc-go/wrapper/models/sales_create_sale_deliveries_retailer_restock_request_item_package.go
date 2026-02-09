package models

type SalesCreateSaleDeliveriesRetailerRestockRequestItemPackage struct {
    PackageLabel string `json:"PackageLabel,omitempty"`
    Quantity int `json:"Quantity,omitempty"`
    RemoveCurrentPackage bool `json:"RemoveCurrentPackage,omitempty"`
    TotalPrice float64 `json:"TotalPrice,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}
