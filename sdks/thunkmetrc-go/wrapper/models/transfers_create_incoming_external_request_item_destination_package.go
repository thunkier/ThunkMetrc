package models

type TransfersCreateIncomingExternalRequestItemDestinationPackage struct {
    ExpirationDate string `json:"ExpirationDate,omitempty"`
    ExternalId string `json:"ExternalId,omitempty"`
    GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
    GrossWeight float64 `json:"GrossWeight,omitempty"`
    IsFinishedGood bool `json:"IsFinishedGood,omitempty"`
    ItemName string `json:"ItemName,omitempty"`
    PackagedDate string `json:"PackagedDate,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
    SellByDate string `json:"SellByDate,omitempty"`
    UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
    UseByDate string `json:"UseByDate,omitempty"`
    WholesalePrice string `json:"WholesalePrice,omitempty"`
}
