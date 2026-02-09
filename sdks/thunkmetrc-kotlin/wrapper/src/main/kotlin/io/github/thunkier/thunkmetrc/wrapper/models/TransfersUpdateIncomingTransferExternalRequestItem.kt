package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersUpdateIncomingTransferExternalRequestItem(
    @JsonProperty("destinations")
    val destinations: List<TransfersUpdateIncomingTransferExternalRequestItemDestinations>? = null,
    @JsonProperty("driverLicenseNumber")
    val driverLicenseNumber: String? = null,
    @JsonProperty("driverName")
    val driverName: String? = null,
    @JsonProperty("driverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("phoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("shipperAddress1")
    val shipperAddress1: String? = null,
    @JsonProperty("shipperAddress2")
    val shipperAddress2: String? = null,
    @JsonProperty("shipperAddressCity")
    val shipperAddressCity: String? = null,
    @JsonProperty("shipperAddressPostalCode")
    val shipperAddressPostalCode: String? = null,
    @JsonProperty("shipperAddressState")
    val shipperAddressState: String? = null,
    @JsonProperty("shipperLicenseNumber")
    val shipperLicenseNumber: String? = null,
    @JsonProperty("shipperMainPhoneNumber")
    val shipperMainPhoneNumber: String? = null,
    @JsonProperty("shipperName")
    val shipperName: String? = null,
    @JsonProperty("transferId")
    val transferId: Int? = null,
    @JsonProperty("transporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("vehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("vehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("vehicleModel")
    val vehicleModel: String? = null
)
