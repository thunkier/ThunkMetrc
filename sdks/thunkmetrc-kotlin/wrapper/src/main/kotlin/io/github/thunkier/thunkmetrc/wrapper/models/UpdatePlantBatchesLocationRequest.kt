package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdatePlantBatchesLocationRequest(
    @JsonProperty("Location")
    val location: String? = null,
    @JsonProperty("MoveDate")
    val moveDate: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("Sublocation")
    val sublocation: String? = null
)
