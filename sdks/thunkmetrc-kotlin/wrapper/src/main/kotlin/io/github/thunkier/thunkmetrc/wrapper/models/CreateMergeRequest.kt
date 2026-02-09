package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateMergeRequest(
    @JsonProperty("packageLabels")
    val packageLabels: List<Any>? = null
)
