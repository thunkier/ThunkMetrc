package models

type SaleReceiptsExternalByExternalNumber struct {
    ArchivedDate *string `json:"ArchivedDate,omitempty"`
    CaregiverLicenseNumber *string `json:"CaregiverLicenseNumber,omitempty"`
    ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IdentificationMethod *string `json:"IdentificationMethod,omitempty"`
    IsFinal bool `json:"IsFinal,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    PatientLicenseNumber *string `json:"PatientLicenseNumber,omitempty"`
    PatientRegistrationLocationId *int64 `json:"PatientRegistrationLocationId,omitempty"`
    ReceiptNumber *string `json:"ReceiptNumber,omitempty"`
    RecordedByUserName *string `json:"RecordedByUserName,omitempty"`
    RecordedDateTime string `json:"RecordedDateTime,omitempty"`
    SalesCustomerType string `json:"SalesCustomerType,omitempty"`
    SalesDateTime string `json:"SalesDateTime,omitempty"`
    TotalPackages int64 `json:"TotalPackages,omitempty"`
    TotalPrice int64 `json:"TotalPrice,omitempty"`
    Transactions []interface{} `json:"Transactions,omitempty"`
}
