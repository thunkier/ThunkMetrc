package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdCreatePackagesInfoRequest(
    @JsonProperty("packageLabels")
    val packageLabels: List<Any>? = null
)
