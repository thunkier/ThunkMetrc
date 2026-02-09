package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class WriteResponse(
    @JsonProperty("Ids")
    val ids: List<Int>? = null,
    @JsonProperty("Warnings")
    val warnings: String? = null
)
