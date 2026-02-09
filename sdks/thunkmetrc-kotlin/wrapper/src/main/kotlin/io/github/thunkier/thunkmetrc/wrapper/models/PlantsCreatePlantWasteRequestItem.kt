package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreatePlantWasteRequestItem(
    @JsonProperty("locationName")
    val locationName: String? = null,
    @JsonProperty("mixedMaterial")
    val mixedMaterial: String? = null,
    @JsonProperty("note")
    val note: String? = null,
    @JsonProperty("plantLabels")
    val plantLabels: String? = null,
    @JsonProperty("reasonName")
    val reasonName: String? = null,
    @JsonProperty("sublocationName")
    val sublocationName: String? = null,
    @JsonProperty("unitOfMeasureName")
    val unitOfMeasureName: String? = null,
    @JsonProperty("wasteDate")
    val wasteDate: String? = null,
    @JsonProperty("wasteMethodName")
    val wasteMethodName: String? = null,
    @JsonProperty("wasteWeight")
    val wasteWeight: Double? = null
)
