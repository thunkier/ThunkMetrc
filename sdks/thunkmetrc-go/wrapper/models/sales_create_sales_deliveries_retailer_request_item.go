package models

type SalesCreateSalesDeliveriesRetailerRequestItem struct {
    DateTime string `json:"DateTime,omitempty"`
    Destinations []*SalesCreateSalesDeliveriesRetailerRequestItemDestination `json:"Destinations,omitempty"`
    DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    Packages []*SalesCreateSalesDeliveriesRetailerRequestItemPackage `json:"Packages,omitempty"`
    PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
