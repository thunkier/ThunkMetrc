package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreatePlantManicureRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("dryingLocation")
    val dryingLocation: String? = null,
    @JsonProperty("dryingSublocation")
    val dryingSublocation: String? = null,
    @JsonProperty("harvestName")
    val harvestName: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("plant")
    val plant: String? = null,
    @JsonProperty("plantCount")
    val plantCount: Int? = null,
    @JsonProperty("unitOfWeight")
    val unitOfWeight: String? = null,
    @JsonProperty("weight")
    val weight: Double? = null
)
