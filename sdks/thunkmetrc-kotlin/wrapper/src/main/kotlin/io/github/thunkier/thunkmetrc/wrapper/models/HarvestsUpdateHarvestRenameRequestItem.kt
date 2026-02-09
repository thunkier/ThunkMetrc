package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class HarvestsUpdateHarvestRenameRequestItem(
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("newName")
    val newName: String? = null,
    @JsonProperty("oldName")
    val oldName: String? = null
)
