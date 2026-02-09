package models

type SalesUpdateDeliveriesRetailerRequestItemDestination struct {
    ConsumerId string `json:"ConsumerId,omitempty"`
    DriverEmployeeId int `json:"DriverEmployeeId,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    Id int `json:"Id,omitempty"`
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
    RecipientZoneId string `json:"RecipientZoneId,omitempty"`
    SalesCustomerType string `json:"SalesCustomerType,omitempty"`
    SalesDateTime string `json:"SalesDateTime,omitempty"`
    Transactions []*SalesUpdateDeliveriesRetailerRequestItemDestinationTransaction `json:"Transactions,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
