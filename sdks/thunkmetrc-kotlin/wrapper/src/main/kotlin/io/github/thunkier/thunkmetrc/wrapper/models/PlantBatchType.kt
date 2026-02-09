package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchType(
    @JsonProperty("CanBeCloned")
    val canBeCloned: Boolean? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("Name")
    val name: String? = null
)
