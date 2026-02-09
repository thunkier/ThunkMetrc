package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdatePlantsLocationRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("Location")
    val location: String? = null,
    @JsonProperty("Sublocation")
    val sublocation: String? = null
)
