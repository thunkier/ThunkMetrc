package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdateLabtestsRequiredRequestItem(
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("RequiredLabTestBatches")
    val requiredLabTestBatches: List<Any>? = null
)
