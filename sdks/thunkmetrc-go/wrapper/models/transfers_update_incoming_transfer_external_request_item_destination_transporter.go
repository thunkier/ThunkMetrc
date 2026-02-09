package models

type TransfersUpdateIncomingTransferExternalRequestItemDestinationTransporter struct {
    DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
    DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    IsLayover bool `json:"IsLayover,omitempty"`
    PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
    TransporterDetails []*TransfersUpdateIncomingTransferExternalRequestItemDestinationTransporterTransporterDetail `json:"TransporterDetails,omitempty"`
    TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
