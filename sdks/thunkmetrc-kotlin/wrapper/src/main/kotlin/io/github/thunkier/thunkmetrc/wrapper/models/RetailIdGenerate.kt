package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdGenerate(
    @JsonProperty("IssuanceId")
    val issuanceId: String? = null
)
