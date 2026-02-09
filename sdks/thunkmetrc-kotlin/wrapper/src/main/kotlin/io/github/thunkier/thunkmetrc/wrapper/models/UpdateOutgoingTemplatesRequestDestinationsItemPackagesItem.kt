package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateOutgoingTemplatesRequestDestinationsItemPackagesItem(
    @JsonProperty("GrossUnitOfWeightName")
    val grossUnitOfWeightName: String? = null,
    @JsonProperty("GrossWeight")
    val grossWeight: Double? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("WholesalePrice")
    val wholesalePrice: String? = null
)
