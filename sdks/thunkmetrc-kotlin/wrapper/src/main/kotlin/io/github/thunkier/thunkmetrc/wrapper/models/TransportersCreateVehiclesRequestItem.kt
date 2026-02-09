package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransportersCreateVehiclesRequestItem(
    @JsonProperty("LicensePlateNumber")
    val licensePlateNumber: String? = null,
    @JsonProperty("Make")
    val make: String? = null,
    @JsonProperty("Model")
    val model: String? = null
)
