package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateIncomingExternalRequestDestinationsItem(
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("GrossUnitOfWeightId")
    val grossUnitOfWeightId: Int? = null,
    @JsonProperty("GrossWeight")
    val grossWeight: Double? = null,
    @JsonProperty("InvoiceNumber")
    val invoiceNumber: String? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("PlannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("RecipientLicenseNumber")
    val recipientLicenseNumber: String? = null,
    @JsonProperty("TransferDestinationId")
    val transferDestinationId: Int? = null,
    @JsonProperty("TransferTypeName")
    val transferTypeName: String? = null,
    @JsonProperty("Transporters")
    val transporters: List<Any>? = null
)
