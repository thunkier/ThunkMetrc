package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsPackages(
    @JsonProperty("grossUnitOfWeightName")
    val grossUnitOfWeightName: String? = null,
    @JsonProperty("grossWeight")
    val grossWeight: Double? = null,
    @JsonProperty("packageLabel")
    val packageLabel: String? = null,
    @JsonProperty("wholesalePrice")
    val wholesalePrice: String? = null
)
