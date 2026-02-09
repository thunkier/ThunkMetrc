package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreatePlantAdditivesUsingTemplateRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("additivesTemplateName")
    val additivesTemplateName: String? = null,
    @JsonProperty("plantLabels")
    val plantLabels: String? = null,
    @JsonProperty("rate")
    val rate: String? = null,
    @JsonProperty("totalAmountApplied")
    val totalAmountApplied: Int? = null,
    @JsonProperty("totalAmountUnitOfMeasure")
    val totalAmountUnitOfMeasure: String? = null,
    @JsonProperty("volume")
    val volume: String? = null
)
