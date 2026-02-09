package models

type Facility struct {
    Alias *string `json:"Alias,omitempty"`
    AllowTransferOfOnHoldPackages bool `json:"AllowTransferOfOnHoldPackages,omitempty"`
    AllowTransferOfOnRecallPackages bool `json:"AllowTransferOfOnRecallPackages,omitempty"`
    CredentialedDate string `json:"CredentialedDate,omitempty"`
    DisplayName string `json:"DisplayName,omitempty"`
    Email *string `json:"Email,omitempty"`
    FacilityId int64 `json:"FacilityId,omitempty"`
    FacilityType FacilityFacilityType `json:"FacilityType,omitempty"`
    HireDate string `json:"HireDate,omitempty"`
    IsFinancialContact bool `json:"IsFinancialContact,omitempty"`
    IsManager bool `json:"IsManager,omitempty"`
    IsOwner bool `json:"IsOwner,omitempty"`
    License FacilityLicense `json:"License,omitempty"`
    Name string `json:"Name,omitempty"`
    Occupations []interface{} `json:"Occupations,omitempty"`
    SupportActivationDate *string `json:"SupportActivationDate,omitempty"`
    SupportExpirationDate *string `json:"SupportExpirationDate,omitempty"`
}
