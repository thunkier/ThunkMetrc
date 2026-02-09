package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsUpdateResultsReleaseRequestItem(
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null
)
