package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreatePlantBatchAdditivesRequestItem(
    @JsonProperty("activeIngredients")
    val activeIngredients: List<PlantBatchesCreatePlantBatchAdditivesRequestItemActiveIngredients>? = null,
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("additiveType")
    val additiveType: String? = null,
    @JsonProperty("applicationDevice")
    val applicationDevice: String? = null,
    @JsonProperty("epaRegistrationNumber")
    val epaRegistrationNumber: String? = null,
    @JsonProperty("plantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("productSupplier")
    val productSupplier: String? = null,
    @JsonProperty("productTradeName")
    val productTradeName: String? = null,
    @JsonProperty("totalAmountApplied")
    val totalAmountApplied: Int? = null,
    @JsonProperty("totalAmountUnitOfMeasure")
    val totalAmountUnitOfMeasure: String? = null
)
