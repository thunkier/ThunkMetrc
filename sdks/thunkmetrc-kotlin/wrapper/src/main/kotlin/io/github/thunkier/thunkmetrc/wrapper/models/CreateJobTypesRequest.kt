package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateJobTypesRequest(
    @JsonProperty("Attributes")
    val attributes: List<Any>? = null,
    @JsonProperty("Category")
    val category: String? = null,
    @JsonProperty("Description")
    val description: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("ProcessingSteps")
    val processingSteps: String? = null
)
