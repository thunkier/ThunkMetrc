package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackageProductLabel(
    @JsonProperty("IsActive")
    val isActive: Boolean? = null,
    @JsonProperty("IsChildFromParentWithLabel")
    val isChildFromParentWithLabel: Boolean? = null,
    @JsonProperty("IsDiscontinued")
    val isDiscontinued: Boolean? = null,
    @JsonProperty("LastLabelGenerationDate")
    val lastLabelGenerationDate: String? = null,
    @JsonProperty("PackageId")
    val packageId: Int? = null,
    @JsonProperty("QrCount")
    val qrCount: Int? = null,
    @JsonProperty("TagId")
    val tagId: Int? = null
)
