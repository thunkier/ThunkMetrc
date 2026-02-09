package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreatePlantingsRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("locationName")
    val locationName: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("plantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("plantBatchType")
    val plantBatchType: String? = null,
    @JsonProperty("plantCount")
    val plantCount: Int? = null,
    @JsonProperty("plantLabel")
    val plantLabel: String? = null,
    @JsonProperty("strainName")
    val strainName: String? = null,
    @JsonProperty("sublocationName")
    val sublocationName: String? = null
)
