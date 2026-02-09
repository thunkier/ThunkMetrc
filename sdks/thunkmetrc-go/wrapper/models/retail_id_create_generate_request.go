package models

type RetailIdCreateGenerateRequest struct {
    PackageLabel string `json:"PackageLabel,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
}
