package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreatePlantBatchesWasteRequestItem(
    @JsonProperty("MixedMaterial")
    val mixedMaterial: String? = null,
    @JsonProperty("Note")
    val note: String? = null,
    @JsonProperty("PlantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("ReasonName")
    val reasonName: String? = null,
    @JsonProperty("UnitOfMeasureName")
    val unitOfMeasureName: String? = null,
    @JsonProperty("WasteDate")
    val wasteDate: String? = null,
    @JsonProperty("WasteMethodName")
    val wasteMethodName: String? = null,
    @JsonProperty("WasteWeight")
    val wasteWeight: Double? = null
)
