package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateSaleDeliveriesRetailerRequestItemPackages(
    @JsonProperty("dateTime")
    val dateTime: String? = null,
    @JsonProperty("packageLabel")
    val packageLabel: String? = null,
    @JsonProperty("quantity")
    val quantity: Int? = null,
    @JsonProperty("totalPrice")
    val totalPrice: Double? = null,
    @JsonProperty("unitOfMeasure")
    val unitOfMeasure: String? = null
)
