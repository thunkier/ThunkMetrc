package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsCreateRecordRequestItem(
    @JsonProperty("DocumentFileBase64")
    val documentFileBase64: String? = null,
    @JsonProperty("DocumentFileName")
    val documentFileName: String? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("ResultDate")
    val resultDate: String? = null,
    @JsonProperty("Results")
    val results: List<Any>? = null
)
