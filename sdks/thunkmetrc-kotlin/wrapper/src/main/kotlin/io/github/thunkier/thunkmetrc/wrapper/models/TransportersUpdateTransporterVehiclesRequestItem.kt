package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransportersUpdateTransporterVehiclesRequestItem(
    @JsonProperty("id")
    val id: String? = null,
    @JsonProperty("licensePlateNumber")
    val licensePlateNumber: String? = null,
    @JsonProperty("make")
    val make: String? = null,
    @JsonProperty("model")
    val model: String? = null
)
