package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Vehicle(
    @JsonProperty("FacilityId")
    val facilityId: Int? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsArchived")
    val isArchived: Boolean? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("LicensePlateNumber")
    val licensePlateNumber: String? = null,
    @JsonProperty("Make")
    val make: String? = null,
    @JsonProperty("Model")
    val model: String? = null
)
