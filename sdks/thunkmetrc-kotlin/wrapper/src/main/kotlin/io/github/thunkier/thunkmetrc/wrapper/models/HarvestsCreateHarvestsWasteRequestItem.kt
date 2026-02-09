package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsCreateHarvestsWasteRequestItem(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("UnitOfWeight")
    val unitOfWeight: String? = null,
    @JsonProperty("WasteType")
    val wasteType: String? = null,
    @JsonProperty("WasteWeight")
    val wasteWeight: Double? = null
)
