package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdatePlantsTagRequestItem(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("NewTag")
    val newTag: String? = null,
    @JsonProperty("ReplaceDate")
    val replaceDate: String? = null,
    @JsonProperty("TagId")
    val tagId: Int? = null
)
