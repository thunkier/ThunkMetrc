package models

type SalesUpdateSaleDeliveriesRetailerRequestItemPackage struct {
    DateTime string `json:"DateTime,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    Quantity int `json:"Quantity,omitempty"`
    TotalPrice float64 `json:"TotalPrice,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}
