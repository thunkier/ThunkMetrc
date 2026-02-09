package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Patient(
    @JsonProperty("HasSalesLimitExemption")
    val hasSalesLimitExemption: Boolean? = null,
    @JsonProperty("LicenseEffectiveEndDate")
    val licenseEffectiveEndDate: String? = null,
    @JsonProperty("LicenseEffectiveStartDate")
    val licenseEffectiveStartDate: String? = null,
    @JsonProperty("LicenseNumber")
    val licenseNumber: String? = null,
    @JsonProperty("OtherFacilitiesCount")
    val otherFacilitiesCount: Int? = null,
    @JsonProperty("PatientId")
    val patientId: Int? = null,
    @JsonProperty("RecommendedPlants")
    val recommendedPlants: Int? = null,
    @JsonProperty("RecommendedSmokableQuantity")
    val recommendedSmokableQuantity: Double? = null,
    @JsonProperty("RegistrationDate")
    val registrationDate: String? = null
)
