package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateSaleDeliveriesCompleteRequestItem(
    @JsonProperty("acceptedPackages")
    val acceptedPackages: String? = null,
    @JsonProperty("actualArrivalDateTime")
    val actualArrivalDateTime: String? = null,
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("paymentType")
    val paymentType: String? = null,
    @JsonProperty("returnedPackages")
    val returnedPackages: List<SalesUpdateSaleDeliveriesCompleteRequestItemReturnedPackages>? = null
)
