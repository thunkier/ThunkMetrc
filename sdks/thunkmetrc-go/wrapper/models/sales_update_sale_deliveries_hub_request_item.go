package models

type SalesUpdateSaleDeliveriesHubRequestItem struct {
    DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    Id int `json:"Id,omitempty"`
    PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
    PlannedRoute string `json:"PlannedRoute,omitempty"`
    TransporterFacilityId string `json:"TransporterFacilityId,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
