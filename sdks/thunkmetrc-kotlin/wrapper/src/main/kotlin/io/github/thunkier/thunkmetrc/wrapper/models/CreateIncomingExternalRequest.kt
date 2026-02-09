package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateIncomingExternalRequest(
    @JsonProperty("Destinations")
    val destinations: List<Any>? = null,
    @JsonProperty("DriverLicenseNumber")
    val driverLicenseNumber: String? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("PhoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("ShipperAddress1")
    val shipperAddress1: String? = null,
    @JsonProperty("ShipperAddress2")
    val shipperAddress2: String? = null,
    @JsonProperty("ShipperAddressCity")
    val shipperAddressCity: String? = null,
    @JsonProperty("ShipperAddressPostalCode")
    val shipperAddressPostalCode: String? = null,
    @JsonProperty("ShipperAddressState")
    val shipperAddressState: String? = null,
    @JsonProperty("ShipperLicenseNumber")
    val shipperLicenseNumber: String? = null,
    @JsonProperty("ShipperMainPhoneNumber")
    val shipperMainPhoneNumber: String? = null,
    @JsonProperty("ShipperName")
    val shipperName: String? = null,
    @JsonProperty("TransporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null
)
