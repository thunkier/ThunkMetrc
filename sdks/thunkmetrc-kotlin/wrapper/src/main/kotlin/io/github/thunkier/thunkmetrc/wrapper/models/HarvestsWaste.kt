package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsWaste(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("UnitOfWeightName")
    val unitOfWeightName: String? = null,
    @JsonProperty("WasteTypeName")
    val wasteTypeName: String? = null,
    @JsonProperty("WasteWeight")
    val wasteWeight: Double? = null
)
