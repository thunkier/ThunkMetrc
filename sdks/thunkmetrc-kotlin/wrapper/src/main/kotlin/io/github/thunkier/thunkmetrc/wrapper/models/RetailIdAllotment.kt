package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdAllotment(
    @JsonProperty("Allotment")
    val allotment: Int? = null,
    @JsonProperty("FacilityId")
    val facilityId: Int? = null,
    @JsonProperty("FacilityLicenseNumber")
    val facilityLicenseNumber: String? = null
)
