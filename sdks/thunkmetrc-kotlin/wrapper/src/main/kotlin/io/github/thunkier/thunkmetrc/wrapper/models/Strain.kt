package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Strain(
    @JsonProperty("CbdLevel")
    val cbdLevel: String? = null,
    @JsonProperty("Genetics")
    val genetics: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IndicaPercentage")
    val indicaPercentage: Int? = null,
    @JsonProperty("IsUsed")
    val isUsed: Boolean? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("SativaPercentage")
    val sativaPercentage: Int? = null,
    @JsonProperty("TestingStatus")
    val testingStatus: String? = null,
    @JsonProperty("ThcLevel")
    val thcLevel: String? = null
)
