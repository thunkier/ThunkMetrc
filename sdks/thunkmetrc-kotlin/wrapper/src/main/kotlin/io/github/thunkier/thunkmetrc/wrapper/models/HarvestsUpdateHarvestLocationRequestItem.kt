package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsUpdateHarvestLocationRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("dryingLocation")
    val dryingLocation: String? = null,
    @JsonProperty("dryingSublocation")
    val dryingSublocation: String? = null,
    @JsonProperty("harvestName")
    val harvestName: String? = null,
    @JsonProperty("id")
    val id: Int? = null
)
