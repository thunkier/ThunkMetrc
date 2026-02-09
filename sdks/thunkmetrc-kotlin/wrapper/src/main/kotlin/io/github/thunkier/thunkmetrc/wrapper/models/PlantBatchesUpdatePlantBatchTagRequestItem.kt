package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesUpdatePlantBatchTagRequestItem(
    @JsonProperty("group")
    val group: String? = null,
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("newTag")
    val newTag: String? = null,
    @JsonProperty("replaceDate")
    val replaceDate: String? = null,
    @JsonProperty("tagId")
    val tagId: Int? = null
)
