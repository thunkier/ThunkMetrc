package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Staged(
    @JsonProperty("CommissionedDateTime")
    val commissionedDateTime: String? = null,
    @JsonProperty("DetachedDateTime")
    val detachedDateTime: String? = null,
    @JsonProperty("FacilityId")
    val facilityId: Int? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsArchived")
    val isArchived: Boolean? = null,
    @JsonProperty("IsStaged")
    val isStaged: Boolean? = null,
    @JsonProperty("IsUsed")
    val isUsed: Boolean? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("MaxGroupSize")
    val maxGroupSize: Int? = null,
    @JsonProperty("ProductLabel")
    val productLabel: String? = null,
    @JsonProperty("QrCount")
    val qrCount: Int? = null,
    @JsonProperty("StatusName")
    val statusName: String? = null,
    @JsonProperty("TagInventoryTypeName")
    val tagInventoryTypeName: String? = null,
    @JsonProperty("TagTypeId")
    val tagTypeId: Int? = null,
    @JsonProperty("TagTypeName")
    val tagTypeName: String? = null,
    @JsonProperty("UsedDateTime")
    val usedDateTime: String? = null
)
