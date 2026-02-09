package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalePatientRegistrationLocation(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Name")
    val name: String? = null
)
