package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreatePlantBatchesPlantingsRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Count")
    val count: Int? = null,
    @JsonProperty("Location")
    val location: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("SourcePlantBatches")
    val sourcePlantBatches: String? = null,
    @JsonProperty("Strain")
    val strain: String? = null,
    @JsonProperty("Sublocation")
    val sublocation: String? = null,
    @JsonProperty("Type")
    val type: String? = null
)
