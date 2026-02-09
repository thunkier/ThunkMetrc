package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransportersCreateTransporterDriversRequestItem(
    @JsonProperty("driversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("employeeId")
    val employeeId: String? = null,
    @JsonProperty("name")
    val name: String? = null
)
