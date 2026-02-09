package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ItemFile(
    @JsonProperty("ContentType")
    val contentType: String? = null,
    @JsonProperty("FileContents")
    val fileContents: String? = null,
    @JsonProperty("FileDownloadName")
    val fileDownloadName: String? = null
)
