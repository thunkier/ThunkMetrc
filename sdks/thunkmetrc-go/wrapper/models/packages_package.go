package models

type PackagesPackage struct {
    ActualDepartureDateTime string `json:"ActualDepartureDateTime,omitempty"`
    ExternalId *int64 `json:"ExternalId,omitempty"`
    GrossUnitOfWeightAbbreviation string `json:"GrossUnitOfWeightAbbreviation,omitempty"`
    GrossWeight float64 `json:"GrossWeight,omitempty"`
    Id int64 `json:"Id,omitempty"`
    ItemStrainName string `json:"ItemStrainName,omitempty"`
    LabTestingStateName string `json:"LabTestingStateName,omitempty"`
    ManifestNumber string `json:"ManifestNumber,omitempty"`
    PackageId int64 `json:"PackageId,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
    ProductCategoryName string `json:"ProductCategoryName,omitempty"`
    ProductName string `json:"ProductName,omitempty"`
    ReceivedDateTime string `json:"ReceivedDateTime,omitempty"`
    ReceivedQuantity float64 `json:"ReceivedQuantity,omitempty"`
    ReceivedUnitOfMeasureAbbreviation string `json:"ReceivedUnitOfMeasureAbbreviation,omitempty"`
    ReceiverWholesalePrice float64 `json:"ReceiverWholesalePrice,omitempty"`
    RecipientFacilityLicenseNumber string `json:"RecipientFacilityLicenseNumber,omitempty"`
    RecipientFacilityName string `json:"RecipientFacilityName,omitempty"`
    ShipmentPackageStateName string `json:"ShipmentPackageStateName,omitempty"`
    ShippedQuantity float64 `json:"ShippedQuantity,omitempty"`
    ShippedUnitOfMeasureAbbreviation string `json:"ShippedUnitOfMeasureAbbreviation,omitempty"`
    ShipperWholesalePrice float64 `json:"ShipperWholesalePrice,omitempty"`
    SourceHarvestNames string `json:"SourceHarvestNames,omitempty"`
    SourcePackageLabels string `json:"SourcePackageLabels,omitempty"`
}
