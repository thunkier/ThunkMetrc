package models

type SalesUpdateDeliveriesRetailerRequestItem struct {
    DateTime string `json:"DateTime,omitempty"`
    Destinations []*SalesUpdateDeliveriesRetailerRequestItemDestination `json:"Destinations,omitempty"`
    DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    Id int `json:"Id,omitempty"`
    Packages []*SalesUpdateDeliveriesRetailerRequestItemPackage `json:"Packages,omitempty"`
    PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
