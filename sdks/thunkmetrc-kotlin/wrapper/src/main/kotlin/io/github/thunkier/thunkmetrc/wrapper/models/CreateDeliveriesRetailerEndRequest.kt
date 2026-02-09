package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateDeliveriesRetailerEndRequest(
    @JsonProperty("ActualArrivalDateTime")
    val actualArrivalDateTime: String? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("RetailerDeliveryId")
    val retailerDeliveryId: Int? = null
)
