package models

type TransferTemplateDeliveryTransporterDetail struct {
    ActualDriverStartDateTime *string `json:"ActualDriverStartDateTime,omitempty"`
    DriverLayoverLeg *string `json:"DriverLayoverLeg,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
    DriverVehicleLicenseNumber string `json:"DriverVehicleLicenseNumber,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
