package models

type SalesUpdateSaleReceiptsRequestItem struct {
    CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
    ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
    Id int `json:"Id,omitempty"`
    IdentificationMethod string `json:"IdentificationMethod,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
    SalesCustomerType string `json:"SalesCustomerType,omitempty"`
    SalesDateTime string `json:"SalesDateTime,omitempty"`
    Transactions []*SalesUpdateSaleReceiptsRequestItemTransaction `json:"Transactions,omitempty"`
}
