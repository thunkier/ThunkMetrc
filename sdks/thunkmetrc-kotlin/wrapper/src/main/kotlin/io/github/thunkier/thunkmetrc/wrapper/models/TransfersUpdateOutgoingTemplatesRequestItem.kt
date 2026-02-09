package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersUpdateOutgoingTemplatesRequestItem(
    @JsonProperty("Destinations")
    val destinations: List<Any>? = null,
    @JsonProperty("DriverLicenseNumber")
    val driverLicenseNumber: String? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("PhoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("TransferTemplateId")
    val transferTemplateId: Int? = null,
    @JsonProperty("TransporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null
)
