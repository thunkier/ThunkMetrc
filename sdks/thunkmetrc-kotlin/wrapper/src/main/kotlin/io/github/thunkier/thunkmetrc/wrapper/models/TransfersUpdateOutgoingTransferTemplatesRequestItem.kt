package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersUpdateOutgoingTransferTemplatesRequestItem(
    @JsonProperty("destinations")
    val destinations: List<TransfersUpdateOutgoingTransferTemplatesRequestItemDestinations>? = null,
    @JsonProperty("driverLicenseNumber")
    val driverLicenseNumber: String? = null,
    @JsonProperty("driverName")
    val driverName: String? = null,
    @JsonProperty("driverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("name")
    val name: String? = null,
    @JsonProperty("phoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("transferTemplateId")
    val transferTemplateId: Int? = null,
    @JsonProperty("transporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("vehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("vehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("vehicleModel")
    val vehicleModel: String? = null
)
