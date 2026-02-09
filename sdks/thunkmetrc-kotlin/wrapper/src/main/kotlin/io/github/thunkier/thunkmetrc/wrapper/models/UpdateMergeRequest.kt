package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateMergeRequest(
    @JsonProperty("MergeDate")
    val mergeDate: String? = null,
    @JsonProperty("SourcePlantGroupLabel")
    val sourcePlantGroupLabel: String? = null,
    @JsonProperty("TargetPlantGroupLabel")
    val targetPlantGroupLabel: String? = null
)
