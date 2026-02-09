package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdateAdjustPlantsRequestItem(
    @JsonProperty("AdjustCount")
    val adjustCount: Int? = null,
    @JsonProperty("AdjustReason")
    val adjustReason: String? = null,
    @JsonProperty("AdjustmentDate")
    val adjustmentDate: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("ReasonNote")
    val reasonNote: String? = null
)
