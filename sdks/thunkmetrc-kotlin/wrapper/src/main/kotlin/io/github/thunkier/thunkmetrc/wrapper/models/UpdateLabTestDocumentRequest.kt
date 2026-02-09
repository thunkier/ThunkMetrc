package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateLabTestDocumentRequest(
    @JsonProperty("DocumentFileBase64")
    val documentFileBase64: String? = null,
    @JsonProperty("DocumentFileName")
    val documentFileName: String? = null,
    @JsonProperty("LabTestResultId")
    val labTestResultId: Int? = null
)
