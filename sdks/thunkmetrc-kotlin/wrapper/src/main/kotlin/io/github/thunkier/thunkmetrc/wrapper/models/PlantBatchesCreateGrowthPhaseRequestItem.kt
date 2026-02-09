package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreateGrowthPhaseRequestItem(
    @JsonProperty("Count")
    val count: Int? = null,
    @JsonProperty("GrowthDate")
    val growthDate: String? = null,
    @JsonProperty("GrowthPhase")
    val growthPhase: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("NewLocation")
    val newLocation: String? = null,
    @JsonProperty("NewSublocation")
    val newSublocation: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("StartingTag")
    val startingTag: String? = null
)
