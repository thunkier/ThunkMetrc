package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SaleDeliveriesReturnReason(
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("RequiresImmatureWasteWeight")
    val requiresImmatureWasteWeight: Boolean? = null,
    @JsonProperty("RequiresMatureWasteWeight")
    val requiresMatureWasteWeight: Boolean? = null,
    @JsonProperty("RequiresNote")
    val requiresNote: Boolean? = null,
    @JsonProperty("RequiresWasteWeight")
    val requiresWasteWeight: Boolean? = null
)
