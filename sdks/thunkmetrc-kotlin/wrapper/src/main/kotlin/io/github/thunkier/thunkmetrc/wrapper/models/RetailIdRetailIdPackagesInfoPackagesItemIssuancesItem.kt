package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem(
    @JsonProperty("CreatedAt")
    val createdAt: String? = null,
    @JsonProperty("LabelQuantity")
    val labelQuantity: Double? = null,
    @JsonProperty("LabelSet")
    val labelSet: Int? = null
)
