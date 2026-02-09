package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ProcessingJobTypesCategory(
    @JsonProperty("ForItems")
    val forItems: Boolean? = null,
    @JsonProperty("ForProcessingJobs")
    val forProcessingJobs: Boolean? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("RequiresProcessingJobAttributes")
    val requiresProcessingJobAttributes: Boolean? = null
)
