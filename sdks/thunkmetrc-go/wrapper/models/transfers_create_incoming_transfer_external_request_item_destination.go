package models

type TransfersCreateIncomingTransferExternalRequestItemDestination struct {
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
    GrossWeight float64 `json:"GrossWeight,omitempty"`
    InvoiceNumber string `json:"InvoiceNumber,omitempty"`
    Packages []*TransfersCreateIncomingTransferExternalRequestItemDestinationPackage `json:"Packages,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
    TransferTypeName string `json:"TransferTypeName,omitempty"`
    Transporters []*TransfersCreateIncomingTransferExternalRequestItemDestinationTransporter `json:"Transporters,omitempty"`
}
