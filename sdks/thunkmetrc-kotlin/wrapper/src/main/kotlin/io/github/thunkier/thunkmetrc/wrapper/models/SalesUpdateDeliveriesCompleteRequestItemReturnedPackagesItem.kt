package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateDeliveriesCompleteRequestItemReturnedPackagesItem(
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("ReturnQuantityVerified")
    val returnQuantityVerified: Int? = null,
    @JsonProperty("ReturnReason")
    val returnReason: String? = null,
    @JsonProperty("ReturnReasonNote")
    val returnReasonNote: String? = null,
    @JsonProperty("ReturnUnitOfMeasure")
    val returnUnitOfMeasure: String? = null
)
