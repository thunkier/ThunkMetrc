package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class WasteType(
    @JsonProperty("Name")
    val name: String? = null
)
