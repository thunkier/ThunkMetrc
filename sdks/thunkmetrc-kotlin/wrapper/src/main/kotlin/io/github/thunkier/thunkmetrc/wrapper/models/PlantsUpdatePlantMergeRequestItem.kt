package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdatePlantMergeRequestItem(
    @JsonProperty("mergeDate")
    val mergeDate: String? = null,
    @JsonProperty("sourcePlantGroupLabel")
    val sourcePlantGroupLabel: String? = null,
    @JsonProperty("targetPlantGroupLabel")
    val targetPlantGroupLabel: String? = null
)
