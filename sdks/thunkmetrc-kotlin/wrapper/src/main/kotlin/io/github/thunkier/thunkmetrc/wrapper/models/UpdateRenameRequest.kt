package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateRenameRequest(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("NewName")
    val newName: String? = null,
    @JsonProperty("OldName")
    val oldName: String? = null
)
