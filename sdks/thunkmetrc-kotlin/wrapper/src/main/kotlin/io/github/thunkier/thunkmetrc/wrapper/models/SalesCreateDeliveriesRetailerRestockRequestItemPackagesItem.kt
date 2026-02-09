package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateDeliveriesRetailerRestockRequestItemPackagesItem(
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("Quantity")
    val quantity: Double? = null,
    @JsonProperty("RemoveCurrentPackage")
    val removeCurrentPackage: Boolean? = null,
    @JsonProperty("TotalPrice")
    val totalPrice: Double? = null,
    @JsonProperty("UnitOfMeasure")
    val unitOfMeasure: String? = null
)
