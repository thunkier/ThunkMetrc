package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageLocationRequestItem(
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("location")
    val location: String? = null,
    @JsonProperty("moveDate")
    val moveDate: String? = null,
    @JsonProperty("sublocation")
    val sublocation: String? = null
)
