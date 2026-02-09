package models

type TransferType struct {
    BypassApproval bool `json:"BypassApproval,omitempty"`
    ExternalIncomingCanRecordExternalIdentifier bool `json:"ExternalIncomingCanRecordExternalIdentifier,omitempty"`
    ExternalIncomingExternalIdentifierRequired bool `json:"ExternalIncomingExternalIdentifierRequired,omitempty"`
    ExternalOutgoingCanRecordExternalIdentifier bool `json:"ExternalOutgoingCanRecordExternalIdentifier,omitempty"`
    ExternalOutgoingExternalIdentifierRequired bool `json:"ExternalOutgoingExternalIdentifierRequired,omitempty"`
    ForExternalIncomingShipments bool `json:"ForExternalIncomingShipments,omitempty"`
    ForExternalOutgoingShipments bool `json:"ForExternalOutgoingShipments,omitempty"`
    ForLicensedShipments bool `json:"ForLicensedShipments,omitempty"`
    Name string `json:"Name,omitempty"`
    RequiresDestinationGrossWeight bool `json:"RequiresDestinationGrossWeight,omitempty"`
    RequiresInvoiceNumber bool `json:"RequiresInvoiceNumber,omitempty"`
    RequiresPDFDocument bool `json:"RequiresPDFDocument,omitempty"`
    RequiresPackagesGrossWeight bool `json:"RequiresPackagesGrossWeight,omitempty"`
    TransactionType string `json:"TransactionType,omitempty"`
}
