package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdatePlantsStrainRequestItem(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("StrainId")
    val strainId: Int? = null,
    @JsonProperty("StrainName")
    val strainName: String? = null
)
