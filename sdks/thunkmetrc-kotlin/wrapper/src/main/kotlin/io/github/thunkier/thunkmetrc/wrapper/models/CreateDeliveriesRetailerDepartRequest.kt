package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateDeliveriesRetailerDepartRequest(
    @JsonProperty("RetailerDeliveryId")
    val retailerDeliveryId: Int? = null
)
