package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdatePlantStrainRequestItem(
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("strainId")
    val strainId: Int? = null,
    @JsonProperty("strainName")
    val strainName: String? = null
)
