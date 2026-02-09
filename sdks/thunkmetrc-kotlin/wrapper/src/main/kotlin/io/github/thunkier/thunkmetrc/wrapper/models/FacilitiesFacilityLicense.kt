package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class FacilitiesFacilityLicense(
    @JsonProperty("EmployeeLicenseNumber")
    val employeeLicenseNumber: String? = null,
    @JsonProperty("EndDate")
    val endDate: String? = null,
    @JsonProperty("LicenseType")
    val licenseType: String? = null,
    @JsonProperty("Number")
    val number: String? = null,
    @JsonProperty("StartDate")
    val startDate: String? = null
)
