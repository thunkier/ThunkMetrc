package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreateAdjustPlantBatchesRequestItem(
    @JsonProperty("AdjustmentDate")
    val adjustmentDate: String? = null,
    @JsonProperty("AdjustmentReason")
    val adjustmentReason: String? = null,
    @JsonProperty("PlantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("Quantity")
    val quantity: Double? = null,
    @JsonProperty("ReasonNote")
    val reasonNote: String? = null
)
