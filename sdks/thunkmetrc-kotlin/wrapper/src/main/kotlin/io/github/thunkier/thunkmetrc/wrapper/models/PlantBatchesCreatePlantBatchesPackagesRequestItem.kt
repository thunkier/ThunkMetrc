package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreatePlantBatchesPackagesRequestItem(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Count")
    val count: Int? = null,
    @JsonProperty("ExpirationDate")
    val expirationDate: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsDonation")
    val isDonation: Boolean? = null,
    @JsonProperty("IsTradeSample")
    val isTradeSample: Boolean? = null,
    @JsonProperty("Item")
    val item: String? = null,
    @JsonProperty("Location")
    val location: String? = null,
    @JsonProperty("Note")
    val note: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PlantBatch")
    val plantBatch: String? = null,
    @JsonProperty("SellByDate")
    val sellByDate: String? = null,
    @JsonProperty("Sublocation")
    val sublocation: String? = null,
    @JsonProperty("Tag")
    val tag: String? = null,
    @JsonProperty("UseByDate")
    val useByDate: String? = null
)
