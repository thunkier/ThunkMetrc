package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ProcessingJob(
    @JsonProperty("CountUnitOfMeasureAbbreviation")
    val countUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("CountUnitOfMeasureId")
    val countUnitOfMeasureId: Int? = null,
    @JsonProperty("CountUnitOfMeasureName")
    val countUnitOfMeasureName: String? = null,
    @JsonProperty("FinishNote")
    val finishNote: String? = null,
    @JsonProperty("FinishedDate")
    val finishedDate: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsFinished")
    val isFinished: Boolean? = null,
    @JsonProperty("JobTypeId")
    val jobTypeId: Int? = null,
    @JsonProperty("JobTypeName")
    val jobTypeName: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("Number")
    val number: String? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("StartDate")
    val startDate: String? = null,
    @JsonProperty("TotalCount")
    val totalCount: Int? = null,
    @JsonProperty("TotalCountWaste")
    val totalCountWaste: String? = null,
    @JsonProperty("TotalQuantity")
    val totalQuantity: Double? = null,
    @JsonProperty("TotalUnitOfMeasureId")
    val totalUnitOfMeasureId: Int? = null,
    @JsonProperty("TotalVolume")
    val totalVolume: Double? = null,
    @JsonProperty("TotalVolumeWaste")
    val totalVolumeWaste: String? = null,
    @JsonProperty("TotalWeight")
    val totalWeight: Double? = null,
    @JsonProperty("TotalWeightWaste")
    val totalWeightWaste: String? = null,
    @JsonProperty("VolumeUnitOfMeasureAbbreviation")
    val volumeUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("VolumeUnitOfMeasureId")
    val volumeUnitOfMeasureId: Int? = null,
    @JsonProperty("VolumeUnitOfMeasureName")
    val volumeUnitOfMeasureName: String? = null,
    @JsonProperty("WasteCountUnitOfMeasureAbbreviation")
    val wasteCountUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("WasteCountUnitOfMeasureId")
    val wasteCountUnitOfMeasureId: Int? = null,
    @JsonProperty("WasteCountUnitOfMeasureName")
    val wasteCountUnitOfMeasureName: String? = null,
    @JsonProperty("WasteVolumeUnitOfMeasureAbbreviation")
    val wasteVolumeUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("WasteVolumeUnitOfMeasureId")
    val wasteVolumeUnitOfMeasureId: Int? = null,
    @JsonProperty("WasteVolumeUnitOfMeasureName")
    val wasteVolumeUnitOfMeasureName: String? = null,
    @JsonProperty("WasteWeightUnitOfMeasureAbbreviation")
    val wasteWeightUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("WasteWeightUnitOfMeasureId")
    val wasteWeightUnitOfMeasureId: Int? = null,
    @JsonProperty("WasteWeightUnitOfMeasureName")
    val wasteWeightUnitOfMeasureName: String? = null,
    @JsonProperty("WeightUnitOfMeasureAbbreviation")
    val weightUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("WeightUnitOfMeasureId")
    val weightUnitOfMeasureId: Int? = null,
    @JsonProperty("WeightUnitOfMeasureName")
    val weightUnitOfMeasureName: String? = null
)
