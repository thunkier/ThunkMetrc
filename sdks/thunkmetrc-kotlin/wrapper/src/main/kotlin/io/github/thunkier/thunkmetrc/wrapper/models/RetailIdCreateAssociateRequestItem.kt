package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdCreateAssociateRequestItem(
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("QrUrls")
    val qrUrls: List<Any>? = null
)
