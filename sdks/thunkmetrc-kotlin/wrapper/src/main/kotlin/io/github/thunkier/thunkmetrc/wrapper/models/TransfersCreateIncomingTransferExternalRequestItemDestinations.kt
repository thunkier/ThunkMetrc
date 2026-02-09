package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateIncomingTransferExternalRequestItemDestinations(
    @JsonProperty("estimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("estimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("grossUnitOfWeightId")
    val grossUnitOfWeightId: Int? = null,
    @JsonProperty("grossWeight")
    val grossWeight: Double? = null,
    @JsonProperty("invoiceNumber")
    val invoiceNumber: String? = null,
    @JsonProperty("packages")
    val packages: List<TransfersCreateIncomingTransferExternalRequestItemDestinationsPackages>? = null,
    @JsonProperty("plannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("recipientLicenseNumber")
    val recipientLicenseNumber: String? = null,
    @JsonProperty("transferTypeName")
    val transferTypeName: String? = null,
    @JsonProperty("transporters")
    val transporters: List<TransfersCreateIncomingTransferExternalRequestItemDestinationsTransporters>? = null
)
