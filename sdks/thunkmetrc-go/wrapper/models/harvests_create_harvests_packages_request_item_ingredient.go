package models

type HarvestsCreateHarvestsPackagesRequestItemIngredient struct {
    HarvestId int `json:"HarvestId,omitempty"`
    HarvestName string `json:"HarvestName,omitempty"`
    UnitOfWeight string `json:"UnitOfWeight,omitempty"`
    Weight float64 `json:"Weight,omitempty"`
}
