package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty

data class PaginatedResponse<T>(
    @JsonProperty("Total")
    val Total: Int? = null,
    @JsonProperty("TotalPages")
    val TotalPages: Int? = null,
    @JsonProperty("PageSize")
    val PageSize: Int? = null,
    @JsonProperty("Current")
    val Current: Int? = null,
    @JsonProperty("Next")
    val Next: Int? = null,
    @JsonProperty("Previous")
    val Previous: Int? = null,
    @JsonProperty("Data")
    val Data: List<T>? = null
)
