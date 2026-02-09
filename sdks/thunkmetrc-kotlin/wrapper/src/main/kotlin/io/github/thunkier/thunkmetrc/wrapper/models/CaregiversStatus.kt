package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CaregiversStatus(
    @JsonProperty("Active")
    val active: Boolean? = null,
    @JsonProperty("CaregiverLicenseNumber")
    val caregiverLicenseNumber: String? = null,
    @JsonProperty("Patients")
    val patients: List<String>? = null
)
