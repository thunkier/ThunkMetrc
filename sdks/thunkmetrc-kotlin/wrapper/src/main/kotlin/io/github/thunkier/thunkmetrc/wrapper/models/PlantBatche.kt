package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PlantBatche(
    @JsonProperty("DestroyedCount")
    val destroyedCount: Int? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsOnHold")
    val isOnHold: Boolean? = null,
    @JsonProperty("IsOnInvestigation")
    val isOnInvestigation: Boolean? = null,
    @JsonProperty("IsOnInvestigationHold")
    val isOnInvestigationHold: Boolean? = null,
    @JsonProperty("IsOnInvestigationRecall")
    val isOnInvestigationRecall: Boolean? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("LocationId")
    val locationId: Int? = null,
    @JsonProperty("LocationName")
    val locationName: String? = null,
    @JsonProperty("LocationTypeName")
    val locationTypeName: String? = null,
    @JsonProperty("MultiPlantBatch")
    val multiPlantBatch: Boolean? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("PackagedCount")
    val packagedCount: Int? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PlantBatchTypeId")
    val plantBatchTypeId: Int? = null,
    @JsonProperty("PlantBatchTypeName")
    val plantBatchTypeName: String? = null,
    @JsonProperty("PlantedDate")
    val plantedDate: String? = null,
    @JsonProperty("SourcePackageId")
    val sourcePackageId: Int? = null,
    @JsonProperty("SourcePackageLabel")
    val sourcePackageLabel: String? = null,
    @JsonProperty("SourcePlantBatchIds")
    val sourcePlantBatchIds: List<Any>? = null,
    @JsonProperty("SourcePlantBatchNames")
    val sourcePlantBatchNames: String? = null,
    @JsonProperty("SourcePlantId")
    val sourcePlantId: Int? = null,
    @JsonProperty("SourcePlantLabel")
    val sourcePlantLabel: String? = null,
    @JsonProperty("StrainId")
    val strainId: Int? = null,
    @JsonProperty("StrainName")
    val strainName: String? = null,
    @JsonProperty("SublocationId")
    val sublocationId: Int? = null,
    @JsonProperty("SublocationName")
    val sublocationName: String? = null,
    @JsonProperty("TrackedCount")
    val trackedCount: Int? = null,
    @JsonProperty("UntrackedCount")
    val untrackedCount: Int? = null
)
