
export interface CreateReceiptsRequestItem {
    CaregiverLicenseNumber?: string;
    ExternalReceiptNumber?: string;
    IdentificationMethod?: string;
    PatientLicenseNumber?: string;
    PatientRegistrationLocationId?: number;
    SalesCustomerType?: string;
    SalesDateTime?: string;
    Transactions?: any[];
}
