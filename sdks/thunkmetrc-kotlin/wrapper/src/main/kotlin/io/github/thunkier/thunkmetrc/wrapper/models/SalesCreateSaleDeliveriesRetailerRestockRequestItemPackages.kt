package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSaleDeliveriesRetailerRestockRequestItemPackages(
    @JsonProperty("packageLabel")
    val packageLabel: String? = null,
    @JsonProperty("quantity")
    val quantity: Int? = null,
    @JsonProperty("removeCurrentPackage")
    val removeCurrentPackage: Boolean? = null,
    @JsonProperty("totalPrice")
    val totalPrice: Double? = null,
    @JsonProperty("unitOfMeasure")
    val unitOfMeasure: String? = null
)
