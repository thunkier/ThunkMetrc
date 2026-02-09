package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsUpdateRestoreHarvestedPlantsRequestItem(
    @JsonProperty("HarvestId")
    val harvestId: Int? = null,
    @JsonProperty("PlantTags")
    val plantTags: List<Any>? = null
)
