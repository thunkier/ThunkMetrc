package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdateItemRequestItem(
    @JsonProperty("Item")
    val item: String? = null,
    @JsonProperty("Label")
    val label: String? = null
)
