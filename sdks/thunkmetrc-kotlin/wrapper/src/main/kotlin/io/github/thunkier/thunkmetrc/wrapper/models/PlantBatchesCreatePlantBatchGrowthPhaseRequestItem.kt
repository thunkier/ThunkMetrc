package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreatePlantBatchGrowthPhaseRequestItem(
    @JsonProperty("count")
    val count: Int? = null,
    @JsonProperty("growthDate")
    val growthDate: String? = null,
    @JsonProperty("growthPhase")
    val growthPhase: String? = null,
    @JsonProperty("name")
    val name: String? = null,
    @JsonProperty("newLocation")
    val newLocation: String? = null,
    @JsonProperty("newSublocation")
    val newSublocation: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("startingTag")
    val startingTag: String? = null
)
