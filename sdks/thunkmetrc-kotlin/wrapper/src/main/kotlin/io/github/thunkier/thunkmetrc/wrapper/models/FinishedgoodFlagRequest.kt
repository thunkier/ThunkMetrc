package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class FinishedgoodFlagRequest(
    @JsonProperty("Label")
    val label: String? = null
)
