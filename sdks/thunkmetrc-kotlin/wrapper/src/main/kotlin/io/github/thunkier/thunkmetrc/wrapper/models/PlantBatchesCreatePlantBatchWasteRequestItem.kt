package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreatePlantBatchWasteRequestItem(
    @JsonProperty("mixedMaterial")
    val mixedMaterial: String? = null,
    @JsonProperty("note")
    val note: String? = null,
    @JsonProperty("plantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("reasonName")
    val reasonName: String? = null,
    @JsonProperty("unitOfMeasureName")
    val unitOfMeasureName: String? = null,
    @JsonProperty("wasteDate")
    val wasteDate: String? = null,
    @JsonProperty("wasteMethodName")
    val wasteMethodName: String? = null,
    @JsonProperty("wasteWeight")
    val wasteWeight: Double? = null
)
