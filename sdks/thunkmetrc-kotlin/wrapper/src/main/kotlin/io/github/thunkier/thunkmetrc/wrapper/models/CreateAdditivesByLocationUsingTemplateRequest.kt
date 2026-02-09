package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateAdditivesByLocationUsingTemplateRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("AdditivesTemplateName")
    val additivesTemplateName: String? = null,
    @JsonProperty("LocationName")
    val locationName: String? = null,
    @JsonProperty("Rate")
    val rate: String? = null,
    @JsonProperty("SublocationName")
    val sublocationName: String? = null,
    @JsonProperty("TotalAmountApplied")
    val totalAmountApplied: Int? = null,
    @JsonProperty("TotalAmountUnitOfMeasure")
    val totalAmountUnitOfMeasure: String? = null,
    @JsonProperty("Volume")
    val volume: String? = null
)
