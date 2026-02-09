package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSaleDeliveriesRetailerDepartRequestItem(
    @JsonProperty("retailerDeliveryId")
    val retailerDeliveryId: Int? = null
)
