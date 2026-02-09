package models

type TransfersUpdateIncomingExternalRequestItemDestination struct {
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
    GrossWeight float64 `json:"GrossWeight,omitempty"`
    InvoiceNumber string `json:"InvoiceNumber,omitempty"`
    Packages []*TransfersUpdateIncomingExternalRequestItemDestinationPackage `json:"Packages,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
    TransferDestinationId int `json:"TransferDestinationId,omitempty"`
    TransferTypeName string `json:"TransferTypeName,omitempty"`
    Transporters []*TransfersUpdateIncomingExternalRequestItemDestinationTransporter `json:"Transporters,omitempty"`
}
