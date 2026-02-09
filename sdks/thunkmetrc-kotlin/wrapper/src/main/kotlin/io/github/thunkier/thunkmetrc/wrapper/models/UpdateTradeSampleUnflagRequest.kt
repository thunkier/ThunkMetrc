package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateTradeSampleUnflagRequest(
    @JsonProperty("Label")
    val label: String? = null
)
