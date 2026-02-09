package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreateAdditivesByLocationRequestItem(
    @JsonProperty("ActiveIngredients")
    val activeIngredients: List<Any>? = null,
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("AdditiveType")
    val additiveType: String? = null,
    @JsonProperty("ApplicationDevice")
    val applicationDevice: String? = null,
    @JsonProperty("EpaRegistrationNumber")
    val epaRegistrationNumber: String? = null,
    @JsonProperty("LocationName")
    val locationName: String? = null,
    @JsonProperty("ProductSupplier")
    val productSupplier: String? = null,
    @JsonProperty("ProductTradeName")
    val productTradeName: String? = null,
    @JsonProperty("SublocationName")
    val sublocationName: String? = null,
    @JsonProperty("TotalAmountApplied")
    val totalAmountApplied: Int? = null,
    @JsonProperty("TotalAmountUnitOfMeasure")
    val totalAmountUnitOfMeasure: String? = null
)
