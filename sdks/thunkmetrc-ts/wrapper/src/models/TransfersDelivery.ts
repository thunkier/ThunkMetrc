
export interface TransfersDelivery {
    ActualArrivalDateTime?: string;
    ActualDepartureDateTime?: string;
    ActualReturnArrivalDateTime?: string;
    ActualReturnDepartureDateTime?: string;
    DeliveryNumber?: number;
    DeliveryPackageCount?: number;
    DeliveryReceivedPackageCount?: number;
    EstimatedArrivalDateTime?: string;
    EstimatedDepartureDateTime?: string;
    EstimatedReturnArrivalDateTime?: string;
    EstimatedReturnDepartureDateTime?: string;
    GrossUnitOfWeightId?: number;
    GrossUnitOfWeightName?: string;
    GrossWeight?: number;
    Id?: number;
    InvoiceNumber?: string;
    ManifestNumber?: string;
    PDFDocumentFileSystemId?: number;
    PlannedRoute?: string;
    ReceivedDateTime?: string;
    RecipientFacilityLicenseNumber?: string;
    RecipientFacilityName?: string;
    RejectedPackagesReturned?: boolean;
    ShipmentTransactionType?: string;
    ShipmentTypeName?: string;
}
