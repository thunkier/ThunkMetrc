package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesCreatePackagesRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("expirationDate")
    val expirationDate: String? = null,
    @JsonProperty("ingredients")
    val ingredients: List<PackagesCreatePackagesRequestItemIngredients>? = null,
    @JsonProperty("isDonation")
    val isDonation: Boolean? = null,
    @JsonProperty("isProductionBatch")
    val isProductionBatch: Boolean? = null,
    @JsonProperty("isTradeSample")
    val isTradeSample: Boolean? = null,
    @JsonProperty("item")
    val item: String? = null,
    @JsonProperty("labTestStageId")
    val labTestStageId: Int? = null,
    @JsonProperty("location")
    val location: String? = null,
    @JsonProperty("note")
    val note: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("processingJobTypeId")
    val processingJobTypeId: Int? = null,
    @JsonProperty("productRequiresRemediation")
    val productRequiresRemediation: Boolean? = null,
    @JsonProperty("productionBatchNumber")
    val productionBatchNumber: String? = null,
    @JsonProperty("quantity")
    val quantity: Int? = null,
    @JsonProperty("requiredLabTestBatches")
    val requiredLabTestBatches: Boolean? = null,
    @JsonProperty("sellByDate")
    val sellByDate: String? = null,
    @JsonProperty("sublocation")
    val sublocation: String? = null,
    @JsonProperty("tag")
    val tag: String? = null,
    @JsonProperty("unitOfMeasure")
    val unitOfMeasure: String? = null,
    @JsonProperty("useByDate")
    val useByDate: String? = null,
    @JsonProperty("useSameItem")
    val useSameItem: Boolean? = null
)
