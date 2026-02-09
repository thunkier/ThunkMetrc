package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateResultsReleaseRequest(
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null
)
