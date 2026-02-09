package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesUpdatePackageNoteRequestItem(
    @JsonProperty("note")
    val note: String? = null,
    @JsonProperty("packageLabel")
    val packageLabel: String? = null
)
