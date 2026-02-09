package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateSaleDeliveriesCompleteRequestItemReturnedPackages(
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("returnQuantityVerified")
    val returnQuantityVerified: Int? = null,
    @JsonProperty("returnReason")
    val returnReason: String? = null,
    @JsonProperty("returnReasonNote")
    val returnReasonNote: String? = null,
    @JsonProperty("returnUnitOfMeasure")
    val returnUnitOfMeasure: String? = null
)
