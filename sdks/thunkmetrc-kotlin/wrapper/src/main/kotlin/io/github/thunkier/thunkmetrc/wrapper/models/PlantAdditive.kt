package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantAdditive(
    @JsonProperty("AdditiveTypeName")
    val additiveTypeName: String? = null,
    @JsonProperty("AmountUnitOfMeasure")
    val amountUnitOfMeasure: String? = null,
    @JsonProperty("ApplicationDevice")
    val applicationDevice: String? = null,
    @JsonProperty("EpaRegistrationNumber")
    val epaRegistrationNumber: String? = null,
    @JsonProperty("Note")
    val note: String? = null,
    @JsonProperty("PlantBatchId")
    val plantBatchId: Int? = null,
    @JsonProperty("PlantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("PlantCount")
    val plantCount: Int? = null,
    @JsonProperty("ProductSupplier")
    val productSupplier: String? = null,
    @JsonProperty("ProductTradeName")
    val productTradeName: String? = null,
    @JsonProperty("Rate")
    val rate: String? = null,
    @JsonProperty("RestrictiveEntryIntervalQuantityDescription")
    val restrictiveEntryIntervalQuantityDescription: String? = null,
    @JsonProperty("RestrictiveEntryIntervalTimeDescription")
    val restrictiveEntryIntervalTimeDescription: String? = null,
    @JsonProperty("TotalAmountApplied")
    val totalAmountApplied: Int? = null,
    @JsonProperty("Volume")
    val volume: Double? = null
)
