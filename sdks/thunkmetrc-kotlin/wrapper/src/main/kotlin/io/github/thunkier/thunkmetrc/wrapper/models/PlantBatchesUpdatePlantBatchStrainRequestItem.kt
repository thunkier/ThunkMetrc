package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesUpdatePlantBatchStrainRequestItem(
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("name")
    val name: String? = null,
    @JsonProperty("strainId")
    val strainId: Int? = null,
    @JsonProperty("strainName")
    val strainName: String? = null
)
