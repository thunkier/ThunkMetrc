package models

type UnitOfMeasure struct {
    Abbreviation string `json:"Abbreviation,omitempty"`
    Name string `json:"Name,omitempty"`
    QuantityType string `json:"QuantityType,omitempty"`
}
