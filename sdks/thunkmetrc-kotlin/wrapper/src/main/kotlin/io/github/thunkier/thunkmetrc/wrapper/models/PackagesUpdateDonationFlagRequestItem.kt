package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdateDonationFlagRequestItem(
    @JsonProperty("Label")
    val label: String? = null
)
