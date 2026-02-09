package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsUpdateHarvestsLocationRequestItem(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("DryingLocation")
    val dryingLocation: String? = null,
    @JsonProperty("DryingSublocation")
    val dryingSublocation: String? = null,
    @JsonProperty("HarvestName")
    val harvestName: String? = null,
    @JsonProperty("Id")
    val id: Int? = null
)
