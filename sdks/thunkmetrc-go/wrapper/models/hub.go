package models

type Hub struct {
    ActualArrivalDateTime *string `json:"ActualArrivalDateTime,omitempty"`
    ActualDepartureDateTime *string `json:"ActualDepartureDateTime,omitempty"`
    ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime,omitempty"`
    ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime,omitempty"`
    CreatedByUserName *string `json:"CreatedByUserName,omitempty"`
    CreatedDateTime string `json:"CreatedDateTime,omitempty"`
    DeliveryCount int64 `json:"DeliveryCount,omitempty"`
    DeliveryId int64 `json:"DeliveryId,omitempty"`
    DeliveryPackageCount int64 `json:"DeliveryPackageCount,omitempty"`
    DeliveryReceivedPackageCount int64 `json:"DeliveryReceivedPackageCount,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
    DriverVehicleLicenseNumber string `json:"DriverVehicleLicenseNumber,omitempty"`
    EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime,omitempty"`
    EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsLayover bool `json:"IsLayover,omitempty"`
    IsVoided bool `json:"IsVoided,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    ManifestNumber string `json:"ManifestNumber,omitempty"`
    PackageCount int64 `json:"PackageCount,omitempty"`
    ReceivedDateTime *string `json:"ReceivedDateTime,omitempty"`
    ReceivedDeliveryCount int64 `json:"ReceivedDeliveryCount,omitempty"`
    ReceivedPackageCount int64 `json:"ReceivedPackageCount,omitempty"`
    RecipientFacilityLicenseNumber *string `json:"RecipientFacilityLicenseNumber,omitempty"`
    RecipientFacilityName *string `json:"RecipientFacilityName,omitempty"`
    RejectedPackagesReturned bool `json:"RejectedPackagesReturned,omitempty"`
    ShipmentTransactionType int64 `json:"ShipmentTransactionType,omitempty"`
    ShipmentTransporterDetails []TransfersHubShipmentTransporterDetailsItem `json:"ShipmentTransporterDetails,omitempty"`
    ShipmentTypeName *string `json:"ShipmentTypeName,omitempty"`
    ShipperFacilityLicenseNumber string `json:"ShipperFacilityLicenseNumber,omitempty"`
    ShipperFacilityName string `json:"ShipperFacilityName,omitempty"`
    TransporterAcceptedDateTime *string `json:"TransporterAcceptedDateTime,omitempty"`
    TransporterActualArrivalDateTime *string `json:"TransporterActualArrivalDateTime,omitempty"`
    TransporterActualDepartureDateTime *string `json:"TransporterActualDepartureDateTime,omitempty"`
    TransporterEstimatedArrivalDateTime *string `json:"TransporterEstimatedArrivalDateTime,omitempty"`
    TransporterEstimatedDepartureDateTime *string `json:"TransporterEstimatedDepartureDateTime,omitempty"`
    TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
    TransporterFacilityName string `json:"TransporterFacilityName,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
}
