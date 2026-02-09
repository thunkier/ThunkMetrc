package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ItemsUpdateItemBrandRequestItem(
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("name")
    val name: String? = null
)
