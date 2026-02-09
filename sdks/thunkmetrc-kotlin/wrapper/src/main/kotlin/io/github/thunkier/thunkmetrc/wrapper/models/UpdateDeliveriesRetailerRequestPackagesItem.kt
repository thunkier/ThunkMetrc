package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateDeliveriesRetailerRequestPackagesItem(
    @JsonProperty("DateTime")
    val dateTime: String? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("Quantity")
    val quantity: Int? = null,
    @JsonProperty("TotalPrice")
    val totalPrice: Double? = null,
    @JsonProperty("UnitOfMeasure")
    val unitOfMeasure: String? = null
)
