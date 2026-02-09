package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdCreateRetailIdGenerateRequest(
    @JsonProperty("packageLabel")
    val packageLabel: String? = null,
    @JsonProperty("quantity")
    val quantity: Int? = null
)
