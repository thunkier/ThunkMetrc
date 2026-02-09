package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PatientCheckInsUpdatePatientCheckInsRequestItem(
    @JsonProperty("CheckInDate")
    val checkInDate: String? = null,
    @JsonProperty("CheckInLocationId")
    val checkInLocationId: Int? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("RegistrationExpiryDate")
    val registrationExpiryDate: String? = null,
    @JsonProperty("RegistrationStartDate")
    val registrationStartDate: String? = null
)
