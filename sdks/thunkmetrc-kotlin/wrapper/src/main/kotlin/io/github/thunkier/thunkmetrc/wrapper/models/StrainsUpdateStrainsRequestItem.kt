package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class StrainsUpdateStrainsRequestItem(
    @JsonProperty("CbdLevel")
    val cbdLevel: Double? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IndicaPercentage")
    val indicaPercentage: Int? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("SativaPercentage")
    val sativaPercentage: Int? = null,
    @JsonProperty("TestingStatus")
    val testingStatus: String? = null,
    @JsonProperty("ThcLevel")
    val thcLevel: Double? = null
)
