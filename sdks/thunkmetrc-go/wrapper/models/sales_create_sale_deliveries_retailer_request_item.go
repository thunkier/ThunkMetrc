package models

type SalesCreateSaleDeliveriesRetailerRequestItem struct {
    ConsumerId int `json:"ConsumerId,omitempty"`
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
    RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
    RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
    RecipientAddressState string `json:"RecipientAddressState,omitempty"`
    RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
    RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
    RecipientName string `json:"RecipientName,omitempty"`
    RecipientZoneId int `json:"RecipientZoneId,omitempty"`
    RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
    SalesCustomerType string `json:"SalesCustomerType,omitempty"`
    SalesDateTime string `json:"SalesDateTime,omitempty"`
    Transactions []*SalesCreateSaleDeliveriesRetailerRequestItemTransaction `json:"Transactions,omitempty"`
}
