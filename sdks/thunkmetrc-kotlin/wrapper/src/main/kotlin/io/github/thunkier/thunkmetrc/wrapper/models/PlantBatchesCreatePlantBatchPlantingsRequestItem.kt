package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatchesCreatePlantBatchPlantingsRequestItem(
    @JsonProperty("actualDate")
    val actualDate: String? = null,
    @JsonProperty("count")
    val count: Int? = null,
    @JsonProperty("location")
    val location: String? = null,
    @JsonProperty("name")
    val name: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("sourcePlantBatches")
    val sourcePlantBatches: String? = null,
    @JsonProperty("strain")
    val strain: String? = null,
    @JsonProperty("sublocation")
    val sublocation: String? = null,
    @JsonProperty("type")
    val type: String? = null
)
