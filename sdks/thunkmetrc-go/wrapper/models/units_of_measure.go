package models

type UnitsOfMeasure struct {
    Abbreviation string `json:"Abbreviation,omitempty"`
    Name string `json:"Name,omitempty"`
    QuantityType string `json:"QuantityType,omitempty"`
}
