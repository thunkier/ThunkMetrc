package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageUseByDateRequestItem(
    @JsonProperty("expirationDate")
    val expirationDate: String? = null,
    @JsonProperty("label")
    val label: String? = null,
    @JsonProperty("sellByDate")
    val sellByDate: String? = null,
    @JsonProperty("useByDate")
    val useByDate: String? = null
)
