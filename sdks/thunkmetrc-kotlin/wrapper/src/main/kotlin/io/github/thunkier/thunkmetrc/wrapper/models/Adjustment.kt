package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Adjustment(
    @JsonProperty("AdjustedByUserName")
    val adjustedByUserName: String? = null,
    @JsonProperty("AdjustmentDate")
    val adjustmentDate: String? = null,
    @JsonProperty("AdjustmentNote")
    val adjustmentNote: String? = null,
    @JsonProperty("AdjustmentQuantity")
    val adjustmentQuantity: Double? = null,
    @JsonProperty("AdjustmentReasonName")
    val adjustmentReasonName: String? = null,
    @JsonProperty("AdjustmentUnitOfMeasureAbbreviation")
    val adjustmentUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("AdjustmentUnitOfMeasureName")
    val adjustmentUnitOfMeasureName: String? = null,
    @JsonProperty("ItemCategoryName")
    val itemCategoryName: String? = null,
    @JsonProperty("ItemName")
    val itemName: String? = null,
    @JsonProperty("PackageId")
    val packageId: Int? = null,
    @JsonProperty("PackageLabTestResultExpirationDateTime")
    val packageLabTestResultExpirationDateTime: String? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("PackagedDate")
    val packagedDate: String? = null
)
