package models

type TransfersDelivery struct {
    ActualArrivalDateTime *string `json:"ActualArrivalDateTime,omitempty"`
    ActualDepartureDateTime *string `json:"ActualDepartureDateTime,omitempty"`
    ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime,omitempty"`
    ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime,omitempty"`
    DeliveryNumber int64 `json:"DeliveryNumber,omitempty"`
    DeliveryPackageCount int64 `json:"DeliveryPackageCount,omitempty"`
    DeliveryReceivedPackageCount int64 `json:"DeliveryReceivedPackageCount,omitempty"`
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime,omitempty"`
    EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime,omitempty"`
    GrossUnitOfWeightId *int64 `json:"GrossUnitOfWeightId,omitempty"`
    GrossUnitOfWeightName *string `json:"GrossUnitOfWeightName,omitempty"`
    GrossWeight *float64 `json:"GrossWeight,omitempty"`
    Id int64 `json:"Id,omitempty"`
    InvoiceNumber *string `json:"InvoiceNumber,omitempty"`
    ManifestNumber *string `json:"ManifestNumber,omitempty"`
    PDFDocumentFileSystemId *int64 `json:"PDFDocumentFileSystemId,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    ReceivedDateTime string `json:"ReceivedDateTime,omitempty"`
    RecipientFacilityLicenseNumber string `json:"RecipientFacilityLicenseNumber,omitempty"`
    RecipientFacilityName string `json:"RecipientFacilityName,omitempty"`
    RejectedPackagesReturned bool `json:"RejectedPackagesReturned,omitempty"`
    ShipmentTransactionType string `json:"ShipmentTransactionType,omitempty"`
    ShipmentTypeName string `json:"ShipmentTypeName,omitempty"`
}
