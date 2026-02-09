package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateOutgoingTemplatesRequestDestinationsItem(
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("InvoiceNumber")
    val invoiceNumber: String? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("PlannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("RecipientLicenseNumber")
    val recipientLicenseNumber: String? = null,
    @JsonProperty("TransferTypeName")
    val transferTypeName: String? = null,
    @JsonProperty("Transporters")
    val transporters: List<Any>? = null
)
