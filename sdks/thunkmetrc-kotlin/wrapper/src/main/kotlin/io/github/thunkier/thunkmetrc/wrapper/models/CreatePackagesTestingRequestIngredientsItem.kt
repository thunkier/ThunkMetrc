package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreatePackagesTestingRequestIngredientsItem(
    @JsonProperty("HarvestId")
    val harvestId: Int? = null,
    @JsonProperty("HarvestName")
    val harvestName: String? = null,
    @JsonProperty("UnitOfWeight")
    val unitOfWeight: String? = null,
    @JsonProperty("Weight")
    val weight: Double? = null
)
