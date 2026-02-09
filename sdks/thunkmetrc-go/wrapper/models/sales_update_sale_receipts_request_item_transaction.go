package models

type SalesUpdateSaleReceiptsRequestItemTransaction struct {
    CityTax string `json:"CityTax,omitempty"`
    CountyTax string `json:"CountyTax,omitempty"`
    DiscountAmount string `json:"DiscountAmount,omitempty"`
    ExciseTax string `json:"ExciseTax,omitempty"`
    InvoiceNumber string `json:"InvoiceNumber,omitempty"`
    MunicipalTax string `json:"MunicipalTax,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    Price string `json:"Price,omitempty"`
    QrCodes string `json:"QrCodes,omitempty"`
    Quantity int `json:"Quantity,omitempty"`
    SalesTax string `json:"SalesTax,omitempty"`
    SubTotal string `json:"SubTotal,omitempty"`
    TotalAmount float64 `json:"TotalAmount,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
    UnitThcContent float64 `json:"UnitThcContent,omitempty"`
    UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
    UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
    UnitWeight float64 `json:"UnitWeight,omitempty"`
    UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
}
