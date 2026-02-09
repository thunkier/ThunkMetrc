package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdRetailIdPackagesInfoPackagesItem(
    @JsonProperty("EstimatedBalance")
    val estimatedBalance: Int? = null,
    @JsonProperty("HasQrs")
    val hasQrs: Boolean? = null,
    @JsonProperty("IssuanceId")
    val issuanceId: String? = null,
    @JsonProperty("Issuances")
    val issuances: List<RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem>? = null,
    @JsonProperty("QrCount")
    val qrCount: Int? = null,
    @JsonProperty("RequiresVerification")
    val requiresVerification: Boolean? = null,
    @JsonProperty("SiblingCount")
    val siblingCount: Int? = null,
    @JsonProperty("Tag")
    val tag: String? = null
)
