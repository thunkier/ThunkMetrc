package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageItemRequestItem(
    @JsonProperty("item")
    val item: String? = null,
    @JsonProperty("label")
    val label: String? = null
)
