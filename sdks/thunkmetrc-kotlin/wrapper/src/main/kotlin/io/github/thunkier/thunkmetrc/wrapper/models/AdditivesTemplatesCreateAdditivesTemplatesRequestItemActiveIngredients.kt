package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class AdditivesTemplatesCreateAdditivesTemplatesRequestItemActiveIngredients(
    @JsonProperty("name")
    val name: String? = null,
    @JsonProperty("percentage")
    val percentage: Double? = null
)
