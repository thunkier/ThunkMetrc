package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsCreateHarvestPackagesTestingRequestItemIngredients(
    @JsonProperty("harvestId")
    val harvestId: Int? = null,
    @JsonProperty("harvestName")
    val harvestName: String? = null,
    @JsonProperty("unitOfWeight")
    val unitOfWeight: String? = null,
    @JsonProperty("weight")
    val weight: Double? = null
)
