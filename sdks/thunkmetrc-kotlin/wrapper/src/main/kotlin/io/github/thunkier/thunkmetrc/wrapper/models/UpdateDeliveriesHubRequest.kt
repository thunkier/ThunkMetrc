package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateDeliveriesHubRequest(
    @JsonProperty("DriverEmployeeId")
    val driverEmployeeId: String? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("PhoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("PlannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("TransporterFacilityId")
    val transporterFacilityId: String? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null
)
