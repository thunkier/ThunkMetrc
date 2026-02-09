package models

type SalesUpdateDeliveriesRetailerRequestItemPackage struct {
    DateTime string `json:"DateTime,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
    TotalPrice float64 `json:"TotalPrice,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}
