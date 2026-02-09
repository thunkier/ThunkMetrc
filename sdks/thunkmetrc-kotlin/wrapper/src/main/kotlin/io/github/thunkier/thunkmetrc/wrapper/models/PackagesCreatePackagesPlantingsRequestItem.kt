package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesCreatePackagesPlantingsRequestItem(
    @JsonProperty("LocationName")
    val locationName: String? = null,
    @JsonProperty("PackageAdjustmentAmount")
    val packageAdjustmentAmount: Int? = null,
    @JsonProperty("PackageAdjustmentUnitOfMeasureName")
    val packageAdjustmentUnitOfMeasureName: String? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PlantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("PlantBatchType")
    val plantBatchType: String? = null,
    @JsonProperty("PlantCount")
    val plantCount: Int? = null,
    @JsonProperty("PlantedDate")
    val plantedDate: String? = null,
    @JsonProperty("StrainName")
    val strainName: String? = null,
    @JsonProperty("SublocationName")
    val sublocationName: String? = null,
    @JsonProperty("UnpackagedDate")
    val unpackagedDate: String? = null
)
