package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersHubShipmentTransporterDetailsItem(
    @JsonProperty("ActualDriverStartDateTime")
    val actualDriverStartDateTime: String? = null,
    @JsonProperty("DriverLayoverLeg")
    val driverLayoverLeg: String? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("DriverVehicleLicenseNumber")
    val driverVehicleLicenseNumber: String? = null,
    @JsonProperty("ShipmentDeliveryId")
    val shipmentDeliveryId: Int? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null
)
