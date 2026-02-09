package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdateUseByDateRequestItem(
    @JsonProperty("ExpirationDate")
    val expirationDate: String? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("SellByDate")
    val sellByDate: String? = null,
    @JsonProperty("UseByDate")
    val useByDate: String? = null
)
