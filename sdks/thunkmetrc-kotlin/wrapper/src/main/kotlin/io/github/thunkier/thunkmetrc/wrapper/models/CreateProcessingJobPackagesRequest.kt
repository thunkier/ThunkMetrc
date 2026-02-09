package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateProcessingJobPackagesRequest(
    @JsonProperty("ExpirationDate")
    val expirationDate: String? = null,
    @JsonProperty("FinishDate")
    val finishDate: String? = null,
    @JsonProperty("FinishNote")
    val finishNote: String? = null,
    @JsonProperty("FinishProcessingJob")
    val finishProcessingJob: Boolean? = null,
    @JsonProperty("Item")
    val item: String? = null,
    @JsonProperty("JobName")
    val jobName: String? = null,
    @JsonProperty("Location")
    val location: String? = null,
    @JsonProperty("Note")
    val note: String? = null,
    @JsonProperty("PackageDate")
    val packageDate: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("ProductionBatchNumber")
    val productionBatchNumber: String? = null,
    @JsonProperty("Quantity")
    val quantity: Int? = null,
    @JsonProperty("SellByDate")
    val sellByDate: String? = null,
    @JsonProperty("Sublocation")
    val sublocation: String? = null,
    @JsonProperty("Tag")
    val tag: String? = null,
    @JsonProperty("UnitOfMeasure")
    val unitOfMeasure: String? = null,
    @JsonProperty("UseByDate")
    val useByDate: String? = null,
    @JsonProperty("WasteCountQuantity")
    val wasteCountQuantity: String? = null,
    @JsonProperty("WasteCountUnitOfMeasureName")
    val wasteCountUnitOfMeasureName: String? = null,
    @JsonProperty("WasteVolumeQuantity")
    val wasteVolumeQuantity: String? = null,
    @JsonProperty("WasteVolumeUnitOfMeasureName")
    val wasteVolumeUnitOfMeasureName: String? = null,
    @JsonProperty("WasteWeightQuantity")
    val wasteWeightQuantity: String? = null,
    @JsonProperty("WasteWeightUnitOfMeasureName")
    val wasteWeightUnitOfMeasureName: String? = null
)
