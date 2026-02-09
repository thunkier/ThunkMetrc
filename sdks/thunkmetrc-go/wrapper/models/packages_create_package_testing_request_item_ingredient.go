package models

type PackagesCreatePackageTestingRequestItemIngredient struct {
    Package string `json:"Package,omitempty"`
    Quantity int `json:"Quantity,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}
