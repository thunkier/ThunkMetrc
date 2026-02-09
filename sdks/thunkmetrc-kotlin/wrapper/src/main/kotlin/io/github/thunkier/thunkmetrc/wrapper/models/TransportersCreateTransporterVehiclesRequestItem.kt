package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransportersCreateTransporterVehiclesRequestItem(
    @JsonProperty("licensePlateNumber")
    val licensePlateNumber: String? = null,
    @JsonProperty("make")
    val make: String? = null,
    @JsonProperty("model")
    val model: String? = null
)
