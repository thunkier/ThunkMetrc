package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateHarvestRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("DryingLocation")
    val dryingLocation: String? = null,
    @JsonProperty("DryingSublocation")
    val dryingSublocation: String? = null,
    @JsonProperty("HarvestName")
    val harvestName: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("Plant")
    val plant: String? = null,
    @JsonProperty("UnitOfWeight")
    val unitOfWeight: String? = null,
    @JsonProperty("Weight")
    val weight: Double? = null
)
