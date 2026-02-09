package models

type TransfersCreateIncomingExternalRequestItemDestination struct {
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
    GrossWeight float64 `json:"GrossWeight,omitempty"`
    InvoiceNumber string `json:"InvoiceNumber,omitempty"`
    Packages []*TransfersCreateIncomingExternalRequestItemDestinationPackage `json:"Packages,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
    TransferTypeName string `json:"TransferTypeName,omitempty"`
    Transporters []*TransfersCreateIncomingExternalRequestItemDestinationTransporter `json:"Transporters,omitempty"`
}
