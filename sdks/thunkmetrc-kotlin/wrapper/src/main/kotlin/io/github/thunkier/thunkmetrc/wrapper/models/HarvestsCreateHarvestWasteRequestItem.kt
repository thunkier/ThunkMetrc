package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsCreateHarvestWasteRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("unitOfWeight")
    val unitOfWeight: String? = null,
    @JsonProperty("wasteType")
    val wasteType: String? = null,
    @JsonProperty("wasteWeight")
    val wasteWeight: Double? = null
)
