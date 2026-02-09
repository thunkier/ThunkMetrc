package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSaleDeliveriesRetailerEndRequestItem(
    @JsonProperty("actualArrivalDateTime")
    val actualArrivalDateTime: String? = null,
    @JsonProperty("packages")
    val packages: List<SalesCreateSaleDeliveriesRetailerEndRequestItemPackages>? = null,
    @JsonProperty("retailerDeliveryId")
    val retailerDeliveryId: Int? = null
)
