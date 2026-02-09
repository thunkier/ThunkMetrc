package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageExternalidRequestItem(
    @JsonProperty("externalId")
    val externalId: String? = null,
    @JsonProperty("packageLabel")
    val packageLabel: String? = null
)
