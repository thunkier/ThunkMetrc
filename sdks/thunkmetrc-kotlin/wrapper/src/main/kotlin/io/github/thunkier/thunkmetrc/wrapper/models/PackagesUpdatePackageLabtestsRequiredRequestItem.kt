package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageLabtestsRequiredRequestItem(
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("requiredLabTestBatches")
    val requiredLabTestBatches: String? = null
)
