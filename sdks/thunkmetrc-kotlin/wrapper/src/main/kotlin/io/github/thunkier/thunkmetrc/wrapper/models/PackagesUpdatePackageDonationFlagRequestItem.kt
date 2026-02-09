package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageDonationFlagRequestItem(
    @JsonProperty("label")
    val label: String? = null
)
