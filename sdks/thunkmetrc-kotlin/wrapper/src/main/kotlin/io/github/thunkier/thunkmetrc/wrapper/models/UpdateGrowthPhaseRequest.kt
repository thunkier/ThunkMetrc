package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateGrowthPhaseRequest(
    @JsonProperty("GrowthDate")
    val growthDate: String? = null,
    @JsonProperty("GrowthPhase")
    val growthPhase: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("NewLocation")
    val newLocation: String? = null,
    @JsonProperty("NewSublocation")
    val newSublocation: String? = null,
    @JsonProperty("NewTag")
    val newTag: String? = null
)
