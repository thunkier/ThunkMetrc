package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateJobTypesRequest(
    @JsonProperty("Attributes")
    val attributes: List<Any>? = null,
    @JsonProperty("CategoryName")
    val categoryName: String? = null,
    @JsonProperty("Description")
    val description: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("ProcessingSteps")
    val processingSteps: String? = null
)
