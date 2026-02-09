package models

type TransfersCreateIncomingTransferExternalRequestItem struct {
    Destinations []*TransfersCreateIncomingTransferExternalRequestItemDestination `json:"Destinations,omitempty"`
    DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
    PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
    ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
    ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
    ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
    ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
    ShipperAddressState string `json:"ShipperAddressState,omitempty"`
    ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
    ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
    ShipperName string `json:"ShipperName,omitempty"`
    TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
