package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateLocationsRequest(
    @JsonProperty("LocationTypeName")
    val locationTypeName: String? = null,
    @JsonProperty("Name")
    val name: String? = null
)
