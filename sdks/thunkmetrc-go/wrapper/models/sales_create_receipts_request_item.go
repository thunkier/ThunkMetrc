package models

type SalesCreateReceiptsRequestItem struct {
    CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
    ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
    IdentificationMethod string `json:"IdentificationMethod,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
    SalesCustomerType string `json:"SalesCustomerType,omitempty"`
    SalesDateTime string `json:"SalesDateTime,omitempty"`
    Transactions []*SalesCreateReceiptsRequestItemTransaction `json:"Transactions,omitempty"`
}
