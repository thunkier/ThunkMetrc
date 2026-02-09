package models

type SalesCreateDeliveriesRequestItem struct {
    ConsumerId int `json:"ConsumerId,omitempty"`
    DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
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
    SalesCustomerType string `json:"SalesCustomerType,omitempty"`
    SalesDateTime string `json:"SalesDateTime,omitempty"`
    Transactions []*SalesCreateDeliveriesRequestItemTransaction `json:"Transactions,omitempty"`
    TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
