package models

type TransfersUpdateIncomingTransferExternalRequestItemDestination struct {
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
    GrossWeight float64 `json:"GrossWeight,omitempty"`
    InvoiceNumber string `json:"InvoiceNumber,omitempty"`
    Packages []*TransfersUpdateIncomingTransferExternalRequestItemDestinationPackage `json:"Packages,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
    TransferDestinationId int `json:"TransferDestinationId,omitempty"`
    TransferTypeName string `json:"TransferTypeName,omitempty"`
    Transporters []*TransfersUpdateIncomingTransferExternalRequestItemDestinationTransporter `json:"Transporters,omitempty"`
}
