package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdCreateRetailIdPackagesInfoRequest(
    @JsonProperty("packageLabels")
    val packageLabels: String? = null
)
