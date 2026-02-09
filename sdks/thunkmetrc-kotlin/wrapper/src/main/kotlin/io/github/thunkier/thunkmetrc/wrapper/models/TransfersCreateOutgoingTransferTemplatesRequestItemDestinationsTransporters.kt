package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsTransporters(
    @JsonProperty("driverLayoverLeg")
    val driverLayoverLeg: String? = null,
    @JsonProperty("driverLicenseNumber")
    val driverLicenseNumber: String? = null,
    @JsonProperty("driverName")
    val driverName: String? = null,
    @JsonProperty("driverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("estimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("estimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("isLayover")
    val isLayover: Boolean? = null,
    @JsonProperty("phoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("transporterDetails")
    val transporterDetails: List<TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsTransportersTransporterDetails>? = null,
    @JsonProperty("transporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("vehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("vehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("vehicleModel")
    val vehicleModel: String? = null
)
