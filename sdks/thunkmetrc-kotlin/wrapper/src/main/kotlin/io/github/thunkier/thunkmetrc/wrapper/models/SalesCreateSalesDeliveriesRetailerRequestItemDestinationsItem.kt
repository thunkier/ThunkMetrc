package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSalesDeliveriesRetailerRequestItemDestinationsItem(
    @JsonProperty("ConsumerId")
    val consumerId: String? = null,
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("RecipientAddressCity")
    val recipientAddressCity: String? = null,
    @JsonProperty("RecipientAddressCounty")
    val recipientAddressCounty: String? = null,
    @JsonProperty("RecipientAddressPostalCode")
    val recipientAddressPostalCode: String? = null,
    @JsonProperty("RecipientAddressState")
    val recipientAddressState: String? = null,
    @JsonProperty("RecipientAddressStreet1")
    val recipientAddressStreet1: String? = null,
    @JsonProperty("RecipientAddressStreet2")
    val recipientAddressStreet2: String? = null,
    @JsonProperty("RecipientName")
    val recipientName: String? = null,
    @JsonProperty("RecipientZoneId")
    val recipientZoneId: String? = null,
    @JsonProperty("SalesCustomerType")
    val salesCustomerType: String? = null,
    @JsonProperty("Transactions")
    val transactions: List<Any>? = null
)
