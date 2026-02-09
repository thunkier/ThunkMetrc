package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class DeletePlantBatchesRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Count")
    val count: Int? = null,
    @JsonProperty("PlantBatch")
    val plantBatch: String? = null,
    @JsonProperty("ReasonNote")
    val reasonNote: String? = null,
    @JsonProperty("WasteMaterialMixed")
    val wasteMaterialMixed: String? = null,
    @JsonProperty("WasteMethodName")
    val wasteMethodName: String? = null,
    @JsonProperty("WasteReasonName")
    val wasteReasonName: String? = null,
    @JsonProperty("WasteUnitOfMeasure")
    val wasteUnitOfMeasure: String? = null,
    @JsonProperty("WasteWeight")
    val wasteWeight: Double? = null
)
