package models

type SalesCreateSalesDeliveriesRetailerRequestItemDestination struct {
    ConsumerId string `json:"ConsumerId,omitempty"`
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
    RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
    RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
    RecipientAddressState string `json:"RecipientAddressState,omitempty"`
    RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
    RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
    RecipientName string `json:"RecipientName,omitempty"`
    RecipientZoneId string `json:"RecipientZoneId,omitempty"`
    SalesCustomerType string `json:"SalesCustomerType,omitempty"`
    Transactions []*SalesCreateSalesDeliveriesRetailerRequestItemDestinationTransaction `json:"Transactions,omitempty"`
}
