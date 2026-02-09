package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageDonationUnflagRequestItem(
    @JsonProperty("label")
    val label: String? = null
)
