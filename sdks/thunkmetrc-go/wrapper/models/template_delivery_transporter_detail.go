package models

type TemplateDeliveryTransporterDetail struct {
    ActualDriverStartDateTime *string `json:"ActualDriverStartDateTime,omitempty"`
    DriverLayoverLeg *string `json:"DriverLayoverLeg,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
    DriverVehicleLicenseNumber string `json:"DriverVehicleLicenseNumber,omitempty"`
    ShipmentDeliveryId int64 `json:"ShipmentDeliveryId,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
