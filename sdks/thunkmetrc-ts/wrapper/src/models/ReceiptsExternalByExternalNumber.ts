
export interface ReceiptsExternalByExternalNumber {
    ArchivedDate?: string;
    CaregiverLicenseNumber?: string;
    ExternalReceiptNumber?: string;
    Id?: number;
    IdentificationMethod?: string;
    IsFinal?: boolean;
    LastModified?: string;
    PatientLicenseNumber?: string;
    PatientRegistrationLocationId?: number;
    ReceiptNumber?: string;
    RecordedByUserName?: string;
    RecordedDateTime?: string;
    SalesCustomerType?: string;
    SalesDateTime?: string;
    TotalPackages?: number;
    TotalPrice?: number;
    Transactions?: any[];
}
