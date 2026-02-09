package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantsCreatePlantsPlantingsRequestItem(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("LocationName")
    val locationName: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PlantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("PlantBatchType")
    val plantBatchType: String? = null,
    @JsonProperty("PlantCount")
    val plantCount: Int? = null,
    @JsonProperty("PlantLabel")
    val plantLabel: String? = null,
    @JsonProperty("StrainName")
    val strainName: String? = null,
    @JsonProperty("SublocationName")
    val sublocationName: String? = null
)
