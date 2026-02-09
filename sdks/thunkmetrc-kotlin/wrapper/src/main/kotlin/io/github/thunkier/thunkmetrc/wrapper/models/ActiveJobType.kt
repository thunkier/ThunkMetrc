package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ActiveJobType(
    @JsonProperty("Attributes")
    val attributes: List<ProcessingJobActiveJobTypeAttributesItem>? = null,
    @JsonProperty("CategoryName")
    val categoryName: String? = null,
    @JsonProperty("Description")
    val description: String? = null,
    @JsonProperty("ForItems")
    val forItems: Boolean? = null,
    @JsonProperty("ForProcessingJobs")
    val forProcessingJobs: Boolean? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("ProcessingSteps")
    val processingSteps: String? = null,
    @JsonProperty("RequiresProcessingJobAttributes")
    val requiresProcessingJobAttributes: Boolean? = null
)
