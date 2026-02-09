package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateOutgoingTransferTemplatesRequestItemDestinations(
    @JsonProperty("estimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("estimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("invoiceNumber")
    val invoiceNumber: String? = null,
    @JsonProperty("packages")
    val packages: List<TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsPackages>? = null,
    @JsonProperty("plannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("recipientLicenseNumber")
    val recipientLicenseNumber: String? = null,
    @JsonProperty("transferTypeName")
    val transferTypeName: String? = null,
    @JsonProperty("transporters")
    val transporters: List<TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsTransporters>? = null
)
