package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreatePlantAdditivesByLocationUsingTemplateRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("additivesTemplateName")
    val additivesTemplateName: String? = null,
    @JsonProperty("locationName")
    val locationName: String? = null,
    @JsonProperty("rate")
    val rate: String? = null,
    @JsonProperty("sublocationName")
    val sublocationName: String? = null,
    @JsonProperty("totalAmountApplied")
    val totalAmountApplied: Int? = null,
    @JsonProperty("totalAmountUnitOfMeasure")
    val totalAmountUnitOfMeasure: String? = null,
    @JsonProperty("volume")
    val volume: String? = null
)
