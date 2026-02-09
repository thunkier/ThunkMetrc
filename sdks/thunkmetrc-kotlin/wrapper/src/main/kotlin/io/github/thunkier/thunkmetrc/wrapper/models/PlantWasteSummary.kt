package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantWasteSummary(
    @JsonProperty("PlantBatchId")
    val plantBatchId: Int? = null,
    @JsonProperty("PlantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("PlantCount")
    val plantCount: Int? = null,
    @JsonProperty("PlantWasteNumber")
    val plantWasteNumber: String? = null,
    @JsonProperty("WasteDate")
    val wasteDate: String? = null,
    @JsonProperty("WasteMethodName")
    val wasteMethodName: String? = null,
    @JsonProperty("WasteReasonName")
    val wasteReasonName: String? = null,
    @JsonProperty("WasteUnitOfMeasureName")
    val wasteUnitOfMeasureName: String? = null,
    @JsonProperty("WasteWeight")
    val wasteWeight: Double? = null
)
