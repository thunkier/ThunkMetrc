package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class RetailIdPackagesInfo(
    @JsonProperty("Packages")
    val packages: List<RetailIdRetailIdPackagesInfoPackagesItem>? = null
)
