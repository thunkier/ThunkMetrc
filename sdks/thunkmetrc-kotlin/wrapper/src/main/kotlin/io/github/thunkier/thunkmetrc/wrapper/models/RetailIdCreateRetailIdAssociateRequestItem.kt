package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdCreateRetailIdAssociateRequestItem(
    @JsonProperty("packageLabel")
    val packageLabel: String? = null,
    @JsonProperty("qrUrls")
    val qrUrls: String? = null
)
