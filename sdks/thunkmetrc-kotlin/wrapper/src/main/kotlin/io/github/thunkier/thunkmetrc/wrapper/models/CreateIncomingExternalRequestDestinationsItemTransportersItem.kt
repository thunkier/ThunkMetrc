package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateIncomingExternalRequestDestinationsItemTransportersItem(
    @JsonProperty("DriverLayoverLeg")
    val driverLayoverLeg: String? = null,
    @JsonProperty("DriverLicenseNumber")
    val driverLicenseNumber: String? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("IsLayover")
    val isLayover: Boolean? = null,
    @JsonProperty("PhoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("TransporterDetails")
    val transporterDetails: List<Any>? = null,
    @JsonProperty("TransporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null
)
