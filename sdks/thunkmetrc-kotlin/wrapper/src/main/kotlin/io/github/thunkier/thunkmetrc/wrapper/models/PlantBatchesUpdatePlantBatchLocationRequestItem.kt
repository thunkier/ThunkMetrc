package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesUpdatePlantBatchLocationRequestItem(
    @JsonProperty("location")
    val location: String? = null,
    @JsonProperty("moveDate")
    val moveDate: String? = null,
    @JsonProperty("name")
    val name: String? = null,
    @JsonProperty("sublocation")
    val sublocation: String? = null
)
