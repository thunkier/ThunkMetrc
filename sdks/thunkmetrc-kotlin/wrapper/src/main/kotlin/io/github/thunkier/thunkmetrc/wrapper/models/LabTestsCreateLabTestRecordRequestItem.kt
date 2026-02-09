package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsCreateLabTestRecordRequestItem(
    @JsonProperty("documentFileBase64")
    val documentFileBase64: String? = null,
    @JsonProperty("documentFileName")
    val documentFileName: String? = null,
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("resultDate")
    val resultDate: String? = null,
    @JsonProperty("results")
    val results: List<LabTestsCreateLabTestRecordRequestItemResults>? = null
)
