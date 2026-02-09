package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreatePlantPlantBatchPackagesRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("count")
    val count: Int? = null,
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
    @JsonProperty("packageTag")
    val packageTag: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("plantBatchType")
    val plantBatchType: String? = null,
    @JsonProperty("plantLabel")
    val plantLabel: String? = null,
    @JsonProperty("sublocation")
    val sublocation: String? = null
)
