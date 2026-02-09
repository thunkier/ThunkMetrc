package models

type ActiveSaleDeliveriesRetailer struct {
    AcceptedDateTime *string `json:"AcceptedDateTime,omitempty"`
    ActualDepartureDateTime *string `json:"ActualDepartureDateTime,omitempty"`
    AllowFullEdit bool `json:"AllowFullEdit,omitempty"`
    CompletedDateTime *string `json:"CompletedDateTime,omitempty"`
    DateTime string `json:"DateTime,omitempty"`
    Destinations []interface{} `json:"Destinations,omitempty"`
    Direction string `json:"Direction,omitempty"`
    DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
    DriverName string `json:"DriverName,omitempty"`
    DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
    EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
    FacilityLicenseNumber string `json:"FacilityLicenseNumber,omitempty"`
    FacilityName string `json:"FacilityName,omitempty"`
    Id int64 `json:"Id,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    Leg int64 `json:"Leg,omitempty"`
    Packages []interface{} `json:"Packages,omitempty"`
    RecordedByUserName *string `json:"RecordedByUserName,omitempty"`
    RecordedDateTime string `json:"RecordedDateTime,omitempty"`
    RestockDateTime *string `json:"RestockDateTime,omitempty"`
    RetailerDeliveryNumber string `json:"RetailerDeliveryNumber,omitempty"`
    RetailerDeliveryState string `json:"RetailerDeliveryState,omitempty"`
    TotalPackages int64 `json:"TotalPackages,omitempty"`
    TotalPrice int64 `json:"TotalPrice,omitempty"`
    TotalPriceSold int64 `json:"TotalPriceSold,omitempty"`
    VehicleInfo string `json:"VehicleInfo,omitempty"`
    VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
    VehicleMake string `json:"VehicleMake,omitempty"`
    VehicleModel string `json:"VehicleModel,omitempty"`
    VoidedDate *string `json:"VoidedDate,omitempty"`
}
