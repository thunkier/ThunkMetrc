package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesCreatePackagePlantingsRequestItem(
    @JsonProperty("locationName")
    val locationName: String? = null,
    @JsonProperty("packageAdjustmentAmount")
    val packageAdjustmentAmount: Int? = null,
    @JsonProperty("packageAdjustmentUnitOfMeasureName")
    val packageAdjustmentUnitOfMeasureName: String? = null,
    @JsonProperty("packageLabel")
    val packageLabel: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("plantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("plantBatchType")
    val plantBatchType: String? = null,
    @JsonProperty("plantCount")
    val plantCount: Int? = null,
    @JsonProperty("plantedDate")
    val plantedDate: String? = null,
    @JsonProperty("strainName")
    val strainName: String? = null,
    @JsonProperty("sublocationName")
    val sublocationName: String? = null,
    @JsonProperty("unpackagedDate")
    val unpackagedDate: String? = null
)
