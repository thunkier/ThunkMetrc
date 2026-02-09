package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateRecordRequestResultsItem(
    @JsonProperty("LabTestTypeName")
    val labTestTypeName: String? = null,
    @JsonProperty("Notes")
    val notes: String? = null,
    @JsonProperty("Passed")
    val passed: Boolean? = null,
    @JsonProperty("Quantity")
    val quantity: Double? = null
)
