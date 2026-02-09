package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesUpdatePlantBatchNameRequestItem(
    @JsonProperty("group")
    val group: String? = null,
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("newGroup")
    val newGroup: String? = null
)
