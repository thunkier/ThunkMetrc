package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateStrainsRequest(
    @JsonProperty("CbdLevel")
    val cbdLevel: Double? = null,
    @JsonProperty("IndicaPercentage")
    val indicaPercentage: Double? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("SativaPercentage")
    val sativaPercentage: Double? = null,
    @JsonProperty("TestingStatus")
    val testingStatus: String? = null,
    @JsonProperty("ThcLevel")
    val thcLevel: Double? = null
)
