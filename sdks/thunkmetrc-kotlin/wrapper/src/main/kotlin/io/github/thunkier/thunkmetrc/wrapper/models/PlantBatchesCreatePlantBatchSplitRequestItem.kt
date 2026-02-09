package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreatePlantBatchSplitRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("count")
    val count: Int? = null,
    @JsonProperty("groupName")
    val groupName: String? = null,
    @JsonProperty("location")
    val location: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("plantBatch")
    val plantBatch: String? = null,
    @JsonProperty("strain")
    val strain: String? = null,
    @JsonProperty("sublocation")
    val sublocation: String? = null
)
