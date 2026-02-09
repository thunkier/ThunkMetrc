package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateSaleDeliveriesHubVerifyIDRequestItem(
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("paymentType")
    val paymentType: String? = null
)
