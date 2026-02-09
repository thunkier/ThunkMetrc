package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LocationType(
    @JsonProperty("ForHarvests")
    val forHarvests: Boolean? = null,
    @JsonProperty("ForPackages")
    val forPackages: Boolean? = null,
    @JsonProperty("ForPlantBatches")
    val forPlantBatches: Boolean? = null,
    @JsonProperty("ForPlants")
    val forPlants: Boolean? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Name")
    val name: String? = null
)
