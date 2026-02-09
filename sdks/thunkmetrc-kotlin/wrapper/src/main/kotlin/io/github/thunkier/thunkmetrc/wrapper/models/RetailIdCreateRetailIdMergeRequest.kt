package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdCreateRetailIdMergeRequest(
    @JsonProperty("packageLabels")
    val packageLabels: String? = null
)
