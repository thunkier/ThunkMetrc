package models

type TransfersCreateOutgoingTemplatesRequestItem struct {
    Destinations []*TransfersCreateOutgoingTemplatesRequestItemDestination `json:"Destinations,omitempty"`
    DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
    Name string `json:"Name,omitempty"`
    PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
    TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
