package models

type TransfersCreateOutgoingTransferTemplatesRequestItemDestination struct {
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    InvoiceNumber string `json:"InvoiceNumber,omitempty"`
    Packages []*TransfersCreateOutgoingTransferTemplatesRequestItemDestinationPackage `json:"Packages,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
    TransferTypeName string `json:"TransferTypeName,omitempty"`
    Transporters []*TransfersCreateOutgoingTransferTemplatesRequestItemDestinationTransporter `json:"Transporters,omitempty"`
}
