package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateDeliveriesHubVerifyIDRequestItem(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("PaymentType")
    val paymentType: String? = null
)
