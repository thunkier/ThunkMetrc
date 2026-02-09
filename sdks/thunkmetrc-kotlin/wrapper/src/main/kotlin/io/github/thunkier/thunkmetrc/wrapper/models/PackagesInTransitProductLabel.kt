package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesInTransitProductLabel(
    @JsonProperty("IsActive")
    val isActive: Boolean? = null,
    @JsonProperty("IsChildFromParentWithLabel")
    val isChildFromParentWithLabel: Boolean? = null,
    @JsonProperty("LabelSource")
    val labelSource: String? = null,
    @JsonProperty("LastLabelGenerationDate")
    val lastLabelGenerationDate: String? = null,
    @JsonProperty("OriginalSourcePackageId")
    val originalSourcePackageId: Int? = null,
    @JsonProperty("OriginalSourcePackageLabel")
    val originalSourcePackageLabel: String? = null,
    @JsonProperty("QrCodeRanges")
    val qrCodeRanges: String? = null,
    @JsonProperty("QrCount")
    val qrCount: Int? = null
)
