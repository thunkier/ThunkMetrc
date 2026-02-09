package models

type TransfersUpdateOutgoingTransferTemplatesRequestItemDestinationPackage struct {
    GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
    GrossWeight float64 `json:"GrossWeight,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    WholesalePrice string `json:"WholesalePrice,omitempty"`
}
