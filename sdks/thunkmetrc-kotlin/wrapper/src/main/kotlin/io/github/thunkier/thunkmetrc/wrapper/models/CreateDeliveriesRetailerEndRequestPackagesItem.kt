package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateDeliveriesRetailerEndRequestPackagesItem(
    @JsonProperty("EndQuantity")
    val endQuantity: Int? = null,
    @JsonProperty("EndUnitOfMeasure")
    val endUnitOfMeasure: String? = null,
    @JsonProperty("Label")
    val label: String? = null
)
