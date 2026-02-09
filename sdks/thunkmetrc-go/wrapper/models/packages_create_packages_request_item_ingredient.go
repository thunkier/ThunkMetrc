package models

type PackagesCreatePackagesRequestItemIngredient struct {
    Package string `json:"Package,omitempty"`
    Quantity int `json:"Quantity,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}
