package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersType(
    @JsonProperty("BypassApproval")
    val bypassApproval: Boolean? = null,
    @JsonProperty("ExternalIncomingCanRecordExternalIdentifier")
    val externalIncomingCanRecordExternalIdentifier: Boolean? = null,
    @JsonProperty("ExternalIncomingExternalIdentifierRequired")
    val externalIncomingExternalIdentifierRequired: Boolean? = null,
    @JsonProperty("ExternalOutgoingCanRecordExternalIdentifier")
    val externalOutgoingCanRecordExternalIdentifier: Boolean? = null,
    @JsonProperty("ExternalOutgoingExternalIdentifierRequired")
    val externalOutgoingExternalIdentifierRequired: Boolean? = null,
    @JsonProperty("ForExternalIncomingShipments")
    val forExternalIncomingShipments: Boolean? = null,
    @JsonProperty("ForExternalOutgoingShipments")
    val forExternalOutgoingShipments: Boolean? = null,
    @JsonProperty("ForLicensedShipments")
    val forLicensedShipments: Boolean? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("RequiresDestinationGrossWeight")
    val requiresDestinationGrossWeight: Boolean? = null,
    @JsonProperty("RequiresInvoiceNumber")
    val requiresInvoiceNumber: Boolean? = null,
    @JsonProperty("RequiresPDFDocument")
    val requiresPDFDocument: Boolean? = null,
    @JsonProperty("RequiresPackagesGrossWeight")
    val requiresPackagesGrossWeight: Boolean? = null,
    @JsonProperty("TransactionType")
    val transactionType: String? = null
)
