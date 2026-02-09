package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateGenerateRequest(
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("Quantity")
    val quantity: Int? = null
)
