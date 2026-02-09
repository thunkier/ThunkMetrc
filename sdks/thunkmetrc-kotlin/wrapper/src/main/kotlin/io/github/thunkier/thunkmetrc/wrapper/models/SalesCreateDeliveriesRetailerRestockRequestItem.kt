package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateDeliveriesRetailerRestockRequestItem(
    @JsonProperty("DateTime")
    val dateTime: String? = null,
    @JsonProperty("Destinations")
    val destinations: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("RetailerDeliveryId")
    val retailerDeliveryId: Int? = null
)
