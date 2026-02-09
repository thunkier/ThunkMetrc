package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class FinishPackagesRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Label")
    val label: String? = null
)
