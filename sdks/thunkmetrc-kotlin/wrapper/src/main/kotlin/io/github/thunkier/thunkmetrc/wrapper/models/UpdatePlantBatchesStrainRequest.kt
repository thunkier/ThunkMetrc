package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdatePlantBatchesStrainRequest(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("StrainId")
    val strainId: Int? = null,
    @JsonProperty("StrainName")
    val strainName: String? = null
)
