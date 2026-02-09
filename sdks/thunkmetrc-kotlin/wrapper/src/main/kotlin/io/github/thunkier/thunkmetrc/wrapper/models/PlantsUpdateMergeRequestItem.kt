package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsUpdateMergeRequestItem(
    @JsonProperty("MergeDate")
    val mergeDate: String? = null,
    @JsonProperty("SourcePlantGroupLabel")
    val sourcePlantGroupLabel: String? = null,
    @JsonProperty("TargetPlantGroupLabel")
    val targetPlantGroupLabel: String? = null
)
