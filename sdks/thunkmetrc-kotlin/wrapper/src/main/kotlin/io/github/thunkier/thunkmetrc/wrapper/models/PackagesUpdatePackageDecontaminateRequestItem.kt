package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageDecontaminateRequestItem(
    @JsonProperty("decontaminationDate")
    val decontaminationDate: String? = null,
    @JsonProperty("decontaminationMethodName")
    val decontaminationMethodName: String? = null,
    @JsonProperty("decontaminationSteps")
    val decontaminationSteps: String? = null,
    @JsonProperty("packageLabel")
    val packageLabel: String? = null
)
