package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreateAdditivesByLocationRequestItemActiveIngredientsItem(
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("Percentage")
    val percentage: Int? = null
)
