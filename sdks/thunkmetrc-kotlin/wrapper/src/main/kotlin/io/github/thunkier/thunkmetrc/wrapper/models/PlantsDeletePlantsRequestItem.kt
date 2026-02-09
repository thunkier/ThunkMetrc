package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsDeletePlantsRequestItem(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Count")
    val count: Int? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("ReasonNote")
    val reasonNote: String? = null,
    @JsonProperty("WasteMaterialMixed")
    val wasteMaterialMixed: String? = null,
    @JsonProperty("WasteMethodName")
    val wasteMethodName: String? = null,
    @JsonProperty("WasteReasonName")
    val wasteReasonName: String? = null,
    @JsonProperty("WasteUnitOfMeasureName")
    val wasteUnitOfMeasureName: String? = null,
    @JsonProperty("WasteWeight")
    val wasteWeight: Double? = null
)
