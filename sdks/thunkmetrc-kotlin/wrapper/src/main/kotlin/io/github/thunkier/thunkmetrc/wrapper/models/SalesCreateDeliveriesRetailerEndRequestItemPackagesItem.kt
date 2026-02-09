package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateDeliveriesRetailerEndRequestItemPackagesItem(
    @JsonProperty("EndQuantity")
    val endQuantity: Double? = null,
    @JsonProperty("EndUnitOfMeasure")
    val endUnitOfMeasure: String? = null,
    @JsonProperty("Label")
    val label: String? = null
)
