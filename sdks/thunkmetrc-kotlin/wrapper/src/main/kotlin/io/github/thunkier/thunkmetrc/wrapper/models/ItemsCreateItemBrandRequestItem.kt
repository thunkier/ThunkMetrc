package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ItemsCreateItemBrandRequestItem(
    @JsonProperty("name")
    val name: String? = null
)
