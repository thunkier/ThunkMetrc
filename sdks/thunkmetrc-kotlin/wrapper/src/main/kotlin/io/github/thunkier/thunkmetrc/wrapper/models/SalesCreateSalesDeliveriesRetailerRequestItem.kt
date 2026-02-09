package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSalesDeliveriesRetailerRequestItem(
    @JsonProperty("DateTime")
    val dateTime: String? = null,
    @JsonProperty("Destinations")
    val destinations: List<Any>? = null,
    @JsonProperty("DriverEmployeeId")
    val driverEmployeeId: String? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("PhoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null
)
