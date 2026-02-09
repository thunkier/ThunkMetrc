package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateLocationsRequest(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("LocationTypeName")
    val locationTypeName: String? = null,
    @JsonProperty("Name")
    val name: String? = null
)
