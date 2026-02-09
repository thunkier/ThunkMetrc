package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdatePlantSplitRequestItem(
    @JsonProperty("plantCount")
    val plantCount: Int? = null,
    @JsonProperty("sourcePlantLabel")
    val sourcePlantLabel: String? = null,
    @JsonProperty("splitDate")
    val splitDate: String? = null,
    @JsonProperty("strainLabel")
    val strainLabel: String? = null,
    @JsonProperty("tagLabel")
    val tagLabel: String? = null
)
