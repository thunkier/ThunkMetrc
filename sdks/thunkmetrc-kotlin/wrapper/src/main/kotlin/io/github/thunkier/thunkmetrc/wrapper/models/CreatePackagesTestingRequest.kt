package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreatePackagesTestingRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("DecontaminateProduct")
    val decontaminateProduct: Boolean? = null,
    @JsonProperty("DecontaminationDate")
    val decontaminationDate: String? = null,
    @JsonProperty("DecontaminationSteps")
    val decontaminationSteps: String? = null,
    @JsonProperty("ExpirationDate")
    val expirationDate: String? = null,
    @JsonProperty("Ingredients")
    val ingredients: List<Any>? = null,
    @JsonProperty("IsDonation")
    val isDonation: Boolean? = null,
    @JsonProperty("IsProductionBatch")
    val isProductionBatch: Boolean? = null,
    @JsonProperty("IsTradeSample")
    val isTradeSample: Boolean? = null,
    @JsonProperty("Item")
    val item: String? = null,
    @JsonProperty("LabTestStageId")
    val labTestStageId: Int? = null,
    @JsonProperty("Location")
    val location: String? = null,
    @JsonProperty("Note")
    val note: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("ProcessingJobTypeId")
    val processingJobTypeId: Int? = null,
    @JsonProperty("ProductRequiresDecontamination")
    val productRequiresDecontamination: Boolean? = null,
    @JsonProperty("ProductRequiresRemediation")
    val productRequiresRemediation: Boolean? = null,
    @JsonProperty("ProductionBatchNumber")
    val productionBatchNumber: String? = null,
    @JsonProperty("RemediateProduct")
    val remediateProduct: Boolean? = null,
    @JsonProperty("RemediationDate")
    val remediationDate: String? = null,
    @JsonProperty("RemediationMethodId")
    val remediationMethodId: Int? = null,
    @JsonProperty("RemediationSteps")
    val remediationSteps: String? = null,
    @JsonProperty("RequiredLabTestBatches")
    val requiredLabTestBatches: List<Any>? = null,
    @JsonProperty("SellByDate")
    val sellByDate: String? = null,
    @JsonProperty("Sublocation")
    val sublocation: String? = null,
    @JsonProperty("Tag")
    val tag: String? = null,
    @JsonProperty("UnitOfWeight")
    val unitOfWeight: String? = null,
    @JsonProperty("UseByDate")
    val useByDate: String? = null
)
