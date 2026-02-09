package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSaleDeliveriesRetailerRestockRequestItem(
    @JsonProperty("dateTime")
    val dateTime: String? = null,
    @JsonProperty("destinations")
    val destinations: String? = null,
    @JsonProperty("estimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("packages")
    val packages: List<SalesCreateSaleDeliveriesRetailerRestockRequestItemPackages>? = null,
    @JsonProperty("retailerDeliveryId")
    val retailerDeliveryId: Int? = null
)
