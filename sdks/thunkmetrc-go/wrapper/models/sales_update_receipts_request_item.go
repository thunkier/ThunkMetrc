package models

type SalesUpdateReceiptsRequestItem struct {
    CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
    ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
    Id int `json:"Id,omitempty"`
    IdentificationMethod string `json:"IdentificationMethod,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
    SalesCustomerType string `json:"SalesCustomerType,omitempty"`
    SalesDateTime string `json:"SalesDateTime,omitempty"`
    Transactions []*SalesUpdateReceiptsRequestItemTransaction `json:"Transactions,omitempty"`
}
