package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Receive(
    @JsonProperty("ChildTag")
    val childTag: String? = null,
    @JsonProperty("Eaches")
    val eaches: List<String>? = null,
    @JsonProperty("LabelSource")
    val labelSource: String? = null,
    @JsonProperty("QrCount")
    val qrCount: Int? = null,
    @JsonProperty("Ranges")
    val ranges: List<List<Int>>? = null,
    @JsonProperty("RequiresVerification")
    val requiresVerification: Boolean? = null,
    @JsonProperty("SiblingTags")
    val siblingTags: List<String>? = null
)
