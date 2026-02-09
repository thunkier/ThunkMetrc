package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ProcessingJobCreateAdjustProcessingJobRequestItem(
    @JsonProperty("AdjustmentDate")
    val adjustmentDate: String? = null,
    @JsonProperty("AdjustmentNote")
    val adjustmentNote: String? = null,
    @JsonProperty("AdjustmentReason")
    val adjustmentReason: String? = null,
    @JsonProperty("CountUnitOfMeasureName")
    val countUnitOfMeasureName: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("VolumeUnitOfMeasureName")
    val volumeUnitOfMeasureName: String? = null,
    @JsonProperty("WeightUnitOfMeasureName")
    val weightUnitOfMeasureName: String? = null
)
