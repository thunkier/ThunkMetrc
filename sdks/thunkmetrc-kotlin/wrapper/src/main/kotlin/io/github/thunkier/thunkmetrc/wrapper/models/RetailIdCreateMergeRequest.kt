package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdCreateMergeRequest(
    @JsonProperty("packageLabels")
    val packageLabels: List<Any>? = null
)
