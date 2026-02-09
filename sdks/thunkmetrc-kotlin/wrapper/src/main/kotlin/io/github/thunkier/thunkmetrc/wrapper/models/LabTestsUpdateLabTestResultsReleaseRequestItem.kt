package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsUpdateLabTestResultsReleaseRequestItem(
    @JsonProperty("packageLabel")
    val packageLabel: String? = null
)
