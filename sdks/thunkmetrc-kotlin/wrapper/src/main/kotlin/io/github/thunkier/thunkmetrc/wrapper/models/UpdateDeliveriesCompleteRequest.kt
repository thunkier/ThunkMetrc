package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateDeliveriesCompleteRequest(
    @JsonProperty("AcceptedPackages")
    val acceptedPackages: List<Any>? = null,
    @JsonProperty("ActualArrivalDateTime")
    val actualArrivalDateTime: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("PaymentType")
    val paymentType: String? = null,
    @JsonProperty("ReturnedPackages")
    val returnedPackages: List<Any>? = null
)
