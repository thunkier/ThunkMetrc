package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateAdjustPackagesRequest(
    @JsonProperty("AdjustmentDate")
    val adjustmentDate: String? = null,
    @JsonProperty("AdjustmentReason")
    val adjustmentReason: String? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("Quantity")
    val quantity: Int? = null,
    @JsonProperty("ReasonNote")
    val reasonNote: String? = null,
    @JsonProperty("UnitOfMeasure")
    val unitOfMeasure: String? = null
)
