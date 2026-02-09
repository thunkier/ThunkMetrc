package models

type TransfersUpdateOutgoingTransferTemplatesRequestItemDestination struct {
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    InvoiceNumber string `json:"InvoiceNumber,omitempty"`
    Packages []*TransfersUpdateOutgoingTransferTemplatesRequestItemDestinationPackage `json:"Packages,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
    TransferDestinationId int `json:"TransferDestinationId,omitempty"`
    TransferTypeName string `json:"TransferTypeName,omitempty"`
    Transporters []*TransfersUpdateOutgoingTransferTemplatesRequestItemDestinationTransporter `json:"Transporters,omitempty"`
}
