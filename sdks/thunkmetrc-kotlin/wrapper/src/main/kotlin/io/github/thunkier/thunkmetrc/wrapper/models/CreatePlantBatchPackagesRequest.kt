package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreatePlantBatchPackagesRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("Count")
    val count: Int? = null,
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
    @JsonProperty("PackageTag")
    val packageTag: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PlantBatchType")
    val plantBatchType: String? = null,
    @JsonProperty("PlantLabel")
    val plantLabel: String? = null,
    @JsonProperty("Sublocation")
    val sublocation: String? = null
)
