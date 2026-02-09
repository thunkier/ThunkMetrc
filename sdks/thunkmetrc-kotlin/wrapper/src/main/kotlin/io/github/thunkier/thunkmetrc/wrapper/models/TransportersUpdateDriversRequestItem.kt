package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransportersUpdateDriversRequestItem(
    @JsonProperty("DriversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("EmployeeId")
    val employeeId: String? = null,
    @JsonProperty("Id")
    val id: String? = null,
    @JsonProperty("Name")
    val name: String? = null
)
