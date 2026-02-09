package models

type TransfersCreateOutgoingTransferTemplatesRequestItem struct {
    Destinations []*TransfersCreateOutgoingTransferTemplatesRequestItemDestination `json:"Destinations,omitempty"`
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
