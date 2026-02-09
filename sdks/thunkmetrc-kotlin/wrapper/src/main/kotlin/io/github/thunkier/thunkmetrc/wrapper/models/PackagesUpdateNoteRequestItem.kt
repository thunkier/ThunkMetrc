package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdateNoteRequestItem(
    @JsonProperty("Note")
    val note: String? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null
)
