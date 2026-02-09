package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class AdditivesTemplateActiveIngredientsItem(
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("Percentage")
    val percentage: Double? = null
)
