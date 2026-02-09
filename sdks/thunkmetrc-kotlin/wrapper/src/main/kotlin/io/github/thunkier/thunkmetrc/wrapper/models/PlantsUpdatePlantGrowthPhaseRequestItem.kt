package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdatePlantGrowthPhaseRequestItem(
    @JsonProperty("growthDate")
    val growthDate: String? = null,
    @JsonProperty("growthPhase")
    val growthPhase: String? = null,
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("newLocation")
    val newLocation: String? = null,
    @JsonProperty("newSublocation")
    val newSublocation: String? = null,
    @JsonProperty("newTag")
    val newTag: String? = null
)
