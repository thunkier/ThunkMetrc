package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class FinishProcessingJobRequest(
    @JsonProperty("FinishDate")
    val finishDate: String? = null,
    @JsonProperty("FinishNote")
    val finishNote: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("TotalCountWaste")
    val totalCountWaste: String? = null,
    @JsonProperty("TotalVolumeWaste")
    val totalVolumeWaste: String? = null,
    @JsonProperty("TotalWeightWaste")
    val totalWeightWaste: Int? = null,
    @JsonProperty("WasteCountUnitOfMeasureName")
    val wasteCountUnitOfMeasureName: String? = null,
    @JsonProperty("WasteVolumeUnitOfMeasureName")
    val wasteVolumeUnitOfMeasureName: String? = null,
    @JsonProperty("WasteWeightUnitOfMeasureName")
    val wasteWeightUnitOfMeasureName: String? = null
)
