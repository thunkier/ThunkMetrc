package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TagPlantAvailable(
    @JsonProperty("FacilityId")
    val facilityId: Int? = null,
    @JsonProperty("GroupTagTypeId")
    val groupTagTypeId: String? = null,
    @JsonProperty("GroupTagTypeName")
    val groupTagTypeName: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("MaxGroupSize")
    val maxGroupSize: Int? = null,
    @JsonProperty("TagInventoryTypeName")
    val tagInventoryTypeName: String? = null,
    @JsonProperty("TagTypeId")
    val tagTypeId: String? = null,
    @JsonProperty("TagTypeName")
    val tagTypeName: String? = null
)
