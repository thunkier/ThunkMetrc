package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantWaste(
    @JsonProperty("GrowthPhase")
    val growthPhase: Int? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("LocationId")
    val locationId: Int? = null,
    @JsonProperty("LocationName")
    val locationName: String? = null,
    @JsonProperty("PlantId")
    val plantId: Int? = null,
    @JsonProperty("PlantWasteId")
    val plantWasteId: Int? = null,
    @JsonProperty("StateName")
    val stateName: String? = null,
    @JsonProperty("SublocationId")
    val sublocationId: Int? = null,
    @JsonProperty("SublocationName")
    val sublocationName: String? = null,
    @JsonProperty("WasteAmount")
    val wasteAmount: Int? = null,
    @JsonProperty("WasteUnitOfMeasureAbbreviation")
    val wasteUnitOfMeasureAbbreviation: String? = null
)
