package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransportersDriver(
    @JsonProperty("DriversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("EmployeeId")
    val employeeId: String? = null,
    @JsonProperty("FacilityId")
    val facilityId: Int? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsArchived")
    val isArchived: Boolean? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("Name")
    val name: String? = null
)
