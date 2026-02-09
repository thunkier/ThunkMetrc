package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsCreateLabTestRecordRequestItemResults(
    @JsonProperty("labTestTypeName")
    val labTestTypeName: String? = null,
    @JsonProperty("notes")
    val notes: String? = null,
    @JsonProperty("passed")
    val passed: Boolean? = null,
    @JsonProperty("quantity")
    val quantity: Double? = null
)
