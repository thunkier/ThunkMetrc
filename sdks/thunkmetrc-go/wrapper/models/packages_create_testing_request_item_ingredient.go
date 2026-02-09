package models

type PackagesCreateTestingRequestItemIngredient struct {
    Package string `json:"Package,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}
