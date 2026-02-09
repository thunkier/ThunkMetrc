package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ItemsCreatePhotoRequestItem(
    @JsonProperty("EncodedImageBase64")
    val encodedImageBase64: String? = null,
    @JsonProperty("FileName")
    val fileName: String? = null
)
