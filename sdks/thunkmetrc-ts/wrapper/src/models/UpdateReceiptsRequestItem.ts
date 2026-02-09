
export interface UpdateReceiptsRequestItem {
    CaregiverLicenseNumber?: string;
    ExternalReceiptNumber?: string;
    Id?: number;
    IdentificationMethod?: string;
    PatientLicenseNumber?: string;
    PatientRegistrationLocationId?: number;
    SalesCustomerType?: string;
    SalesDateTime?: string;
    Transactions?: any[];
}
