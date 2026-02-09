package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsCreateHarvestPackagesRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("decontaminateProduct")
    val decontaminateProduct: Boolean? = null,
    @JsonProperty("decontaminationDate")
    val decontaminationDate: String? = null,
    @JsonProperty("decontaminationSteps")
    val decontaminationSteps: String? = null,
    @JsonProperty("expirationDate")
    val expirationDate: String? = null,
    @JsonProperty("ingredients")
    val ingredients: List<HarvestsCreateHarvestPackagesRequestItemIngredients>? = null,
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
    @JsonProperty("productRequiresDecontamination")
    val productRequiresDecontamination: Boolean? = null,
    @JsonProperty("productRequiresRemediation")
    val productRequiresRemediation: Boolean? = null,
    @JsonProperty("productionBatchNumber")
    val productionBatchNumber: String? = null,
    @JsonProperty("remediateProduct")
    val remediateProduct: Boolean? = null,
    @JsonProperty("remediationDate")
    val remediationDate: String? = null,
    @JsonProperty("remediationMethodId")
    val remediationMethodId: Int? = null,
    @JsonProperty("remediationSteps")
    val remediationSteps: String? = null,
    @JsonProperty("requiredLabTestBatches")
    val requiredLabTestBatches: String? = null,
    @JsonProperty("sellByDate")
    val sellByDate: String? = null,
    @JsonProperty("sublocation")
    val sublocation: String? = null,
    @JsonProperty("tag")
    val tag: String? = null,
    @JsonProperty("unitOfWeight")
    val unitOfWeight: String? = null,
    @JsonProperty("useByDate")
    val useByDate: String? = null
)
