package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UnitsOfMeasure(
    @JsonProperty("Abbreviation")
    val abbreviation: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("QuantityType")
    val quantityType: String? = null
)
