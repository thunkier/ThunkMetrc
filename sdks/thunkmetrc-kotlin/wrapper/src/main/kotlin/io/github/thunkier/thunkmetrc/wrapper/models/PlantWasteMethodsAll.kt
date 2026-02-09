package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantWasteMethodsAll(
    @JsonProperty("ForPlants")
    val forPlants: Boolean? = null,
    @JsonProperty("ForProductDestruction")
    val forProductDestruction: Boolean? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("Name")
    val name: String? = null
)
