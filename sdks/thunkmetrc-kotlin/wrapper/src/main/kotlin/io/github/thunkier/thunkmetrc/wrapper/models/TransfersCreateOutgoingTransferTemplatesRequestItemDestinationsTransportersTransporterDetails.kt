package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsTransportersTransporterDetails(
    @JsonProperty("driverLayoverLeg")
    val driverLayoverLeg: String? = null,
    @JsonProperty("driverLicenseNumber")
    val driverLicenseNumber: String? = null,
    @JsonProperty("driverName")
    val driverName: String? = null,
    @JsonProperty("driverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("vehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("vehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("vehicleModel")
    val vehicleModel: String? = null
)
