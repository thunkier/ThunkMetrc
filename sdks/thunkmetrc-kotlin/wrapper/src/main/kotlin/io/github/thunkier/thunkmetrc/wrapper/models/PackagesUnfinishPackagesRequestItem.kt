package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUnfinishPackagesRequestItem(
    @JsonProperty("Label")
    val label: String? = null
)
