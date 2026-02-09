package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdatePackagesLocationRequest(
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("Location")
    val location: String? = null,
    @JsonProperty("MoveDate")
    val moveDate: String? = null,
    @JsonProperty("Sublocation")
    val sublocation: String? = null
)
