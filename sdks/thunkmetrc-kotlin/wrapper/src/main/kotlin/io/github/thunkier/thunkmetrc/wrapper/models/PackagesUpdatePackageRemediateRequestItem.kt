package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageRemediateRequestItem(
    @JsonProperty("packageLabel")
    val packageLabel: String? = null,
    @JsonProperty("remediationDate")
    val remediationDate: String? = null,
    @JsonProperty("remediationMethodName")
    val remediationMethodName: String? = null,
    @JsonProperty("remediationSteps")
    val remediationSteps: String? = null
)
