package models

type SalesCreateDeliveriesRetailerRestockRequestItemPackage struct {
    PackageLabel string `json:"PackageLabel,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
    RemoveCurrentPackage bool `json:"RemoveCurrentPackage,omitempty"`
    TotalPrice float64 `json:"TotalPrice,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}
