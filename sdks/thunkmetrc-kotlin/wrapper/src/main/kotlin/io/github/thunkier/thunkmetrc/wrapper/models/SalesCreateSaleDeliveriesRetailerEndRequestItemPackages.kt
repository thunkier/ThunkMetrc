package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSaleDeliveriesRetailerEndRequestItemPackages(
    @JsonProperty("endQuantity")
    val endQuantity: Int? = null,
    @JsonProperty("endUnitOfMeasure")
    val endUnitOfMeasure: String? = null,
    @JsonProperty("label")
    val label: String? = null
)
