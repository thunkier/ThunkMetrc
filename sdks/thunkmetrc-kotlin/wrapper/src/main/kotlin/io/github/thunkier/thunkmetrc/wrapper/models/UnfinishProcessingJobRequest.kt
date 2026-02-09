package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UnfinishProcessingJobRequest(
    @JsonProperty("Id")
    val id: Int? = null
)
