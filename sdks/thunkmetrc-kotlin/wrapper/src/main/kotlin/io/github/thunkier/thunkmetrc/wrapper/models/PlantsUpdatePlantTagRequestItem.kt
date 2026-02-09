package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdatePlantTagRequestItem(
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("newTag")
    val newTag: String? = null,
    @JsonProperty("replaceDate")
    val replaceDate: String? = null,
    @JsonProperty("tagId")
    val tagId: Int? = null
)
