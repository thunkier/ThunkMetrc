package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ProcessingJobStartProcessingJobRequestItemPackages(
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("quantity")
    val quantity: Int? = null,
    @JsonProperty("unitOfMeasure")
    val unitOfMeasure: String? = null
)
