package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdateDecontaminateRequestItem(
    @JsonProperty("DecontaminationDate")
    val decontaminationDate: String? = null,
    @JsonProperty("DecontaminationMethodName")
    val decontaminationMethodName: String? = null,
    @JsonProperty("DecontaminationSteps")
    val decontaminationSteps: String? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null
)
