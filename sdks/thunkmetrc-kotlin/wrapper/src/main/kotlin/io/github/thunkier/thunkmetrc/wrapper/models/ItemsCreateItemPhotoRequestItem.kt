package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ItemsCreateItemPhotoRequestItem(
    @JsonProperty("encodedImageBase64")
    val encodedImageBase64: String? = null,
    @JsonProperty("fileName")
    val fileName: String? = null
)
