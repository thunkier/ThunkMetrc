package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UnfinishPackagesRequest(
    @JsonProperty("Label")
    val label: String? = null
)
