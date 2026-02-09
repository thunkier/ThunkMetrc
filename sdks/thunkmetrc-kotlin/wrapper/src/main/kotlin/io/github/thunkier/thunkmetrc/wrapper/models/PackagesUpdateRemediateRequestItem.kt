package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdateRemediateRequestItem(
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("RemediationDate")
    val remediationDate: String? = null,
    @JsonProperty("RemediationMethodName")
    val remediationMethodName: String? = null,
    @JsonProperty("RemediationSteps")
    val remediationSteps: String? = null
)
