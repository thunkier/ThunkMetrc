package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesUpdateNameRequestItem(
    @JsonProperty("Group")
    val group: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("NewGroup")
    val newGroup: String? = null
)
