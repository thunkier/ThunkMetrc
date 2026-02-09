package models

type TransfersCreateOutgoingTemplatesRequestItemDestination struct {
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    InvoiceNumber string `json:"InvoiceNumber,omitempty"`
    Packages []*TransfersCreateOutgoingTemplatesRequestItemDestinationPackage `json:"Packages,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
    TransferTypeName string `json:"TransferTypeName,omitempty"`
    Transporters []*TransfersCreateOutgoingTemplatesRequestItemDestinationTransporter `json:"Transporters,omitempty"`
}
