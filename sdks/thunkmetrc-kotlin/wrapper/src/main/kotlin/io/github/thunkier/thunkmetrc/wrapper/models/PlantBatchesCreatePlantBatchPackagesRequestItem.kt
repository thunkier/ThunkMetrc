package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreatePlantBatchPackagesRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("count")
    val count: Int? = null,
    @JsonProperty("expirationDate")
    val expirationDate: String? = null,
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("isDonation")
    val isDonation: Boolean? = null,
    @JsonProperty("isTradeSample")
    val isTradeSample: Boolean? = null,
    @JsonProperty("item")
    val item: String? = null,
    @JsonProperty("location")
    val location: String? = null,
    @JsonProperty("note")
    val note: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("plantBatch")
    val plantBatch: String? = null,
    @JsonProperty("sellByDate")
    val sellByDate: String? = null,
    @JsonProperty("sublocation")
    val sublocation: String? = null,
    @JsonProperty("tag")
    val tag: String? = null,
    @JsonProperty("useByDate")
    val useByDate: String? = null
)
