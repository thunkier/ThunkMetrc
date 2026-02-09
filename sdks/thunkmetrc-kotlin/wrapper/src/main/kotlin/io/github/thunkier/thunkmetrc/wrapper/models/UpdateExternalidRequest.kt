package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateExternalidRequest(
    @JsonProperty("ExternalId")
    val externalId: String? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null
)
