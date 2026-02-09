package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateVehiclesRequest(
    @JsonProperty("Id")
    val id: String? = null,
    @JsonProperty("LicensePlateNumber")
    val licensePlateNumber: String? = null,
    @JsonProperty("Make")
    val make: String? = null,
    @JsonProperty("Model")
    val model: String? = null
)
