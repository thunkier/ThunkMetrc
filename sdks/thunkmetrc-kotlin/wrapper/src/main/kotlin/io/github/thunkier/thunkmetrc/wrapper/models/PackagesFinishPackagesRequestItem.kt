package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesFinishPackagesRequestItem(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Label")
    val label: String? = null
)
