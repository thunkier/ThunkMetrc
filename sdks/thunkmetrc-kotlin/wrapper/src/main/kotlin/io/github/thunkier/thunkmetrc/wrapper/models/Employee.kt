package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Employee(
    @JsonProperty("FullName")
    val fullName: String? = null,
    @JsonProperty("IsIndustryAdmin")
    val isIndustryAdmin: Boolean? = null,
    @JsonProperty("IsManager")
    val isManager: Boolean? = null,
    @JsonProperty("IsOwner")
    val isOwner: Boolean? = null,
    @JsonProperty("License")
    val license: EmployeeLicense? = null
)
