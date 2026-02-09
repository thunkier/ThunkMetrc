package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UnfinishHarvestsRequest(
    @JsonProperty("Id")
    val id: Int? = null
)
