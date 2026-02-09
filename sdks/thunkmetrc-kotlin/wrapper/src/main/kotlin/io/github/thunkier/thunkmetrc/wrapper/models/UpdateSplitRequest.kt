package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateSplitRequest(
    @JsonProperty("PlantCount")
    val plantCount: Int? = null,
    @JsonProperty("SourcePlantLabel")
    val sourcePlantLabel: String? = null,
    @JsonProperty("SplitDate")
    val splitDate: String? = null,
    @JsonProperty("StrainLabel")
    val strainLabel: String? = null,
    @JsonProperty("TagLabel")
    val tagLabel: String? = null
)
