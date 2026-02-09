
export interface TransfersType {
    BypassApproval?: boolean;
    ExternalIncomingCanRecordExternalIdentifier?: boolean;
    ExternalIncomingExternalIdentifierRequired?: boolean;
    ExternalOutgoingCanRecordExternalIdentifier?: boolean;
    ExternalOutgoingExternalIdentifierRequired?: boolean;
    ForExternalIncomingShipments?: boolean;
    ForExternalOutgoingShipments?: boolean;
    ForLicensedShipments?: boolean;
    Name?: string;
    RequiresDestinationGrossWeight?: boolean;
    RequiresInvoiceNumber?: boolean;
    RequiresPDFDocument?: boolean;
    RequiresPackagesGrossWeight?: boolean;
    TransactionType?: string;
}
