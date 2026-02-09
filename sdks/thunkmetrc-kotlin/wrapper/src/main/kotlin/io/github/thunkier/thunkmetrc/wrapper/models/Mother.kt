package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Mother(
    @JsonProperty("ClonedCount")
    val clonedCount: Int? = null,
    @JsonProperty("DescendedCount")
    val descendedCount: Int? = null,
    @JsonProperty("DestroyedByUserName")
    val destroyedByUserName: String? = null,
    @JsonProperty("DestroyedDate")
    val destroyedDate: String? = null,
    @JsonProperty("DestroyedNote")
    val destroyedNote: String? = null,
    @JsonProperty("FloweringDate")
    val floweringDate: String? = null,
    @JsonProperty("GroupTagTypeMax")
    val groupTagTypeMax: Int? = null,
    @JsonProperty("GrowthPhase")
    val growthPhase: String? = null,
    @JsonProperty("HarvestCount")
    val harvestCount: Int? = null,
    @JsonProperty("HarvestId")
    val harvestId: Int? = null,
    @JsonProperty("HarvestedDate")
    val harvestedDate: String? = null,
    @JsonProperty("HarvestedUnitOfWeightAbbreviation")
    val harvestedUnitOfWeightAbbreviation: String? = null,
    @JsonProperty("HarvestedUnitOfWeightName")
    val harvestedUnitOfWeightName: String? = null,
    @JsonProperty("HarvestedWetWeight")
    val harvestedWetWeight: Double? = null,
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
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("LocationId")
    val locationId: Int? = null,
    @JsonProperty("LocationName")
    val locationName: String? = null,
    @JsonProperty("LocationTypeName")
    val locationTypeName: String? = null,
    @JsonProperty("MotherPlantDate")
    val motherPlantDate: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PlantBatchId")
    val plantBatchId: Int? = null,
    @JsonProperty("PlantBatchName")
    val plantBatchName: String? = null,
    @JsonProperty("PlantBatchTypeId")
    val plantBatchTypeId: Int? = null,
    @JsonProperty("PlantBatchTypeName")
    val plantBatchTypeName: String? = null,
    @JsonProperty("PlantedDate")
    val plantedDate: String? = null,
    @JsonProperty("State")
    val state: String? = null,
    @JsonProperty("StrainId")
    val strainId: Int? = null,
    @JsonProperty("StrainName")
    val strainName: String? = null,
    @JsonProperty("SublocationId")
    val sublocationId: Int? = null,
    @JsonProperty("SublocationName")
    val sublocationName: String? = null,
    @JsonProperty("SurvivedCount")
    val survivedCount: Int? = null,
    @JsonProperty("TagTypeMax")
    val tagTypeMax: Int? = null,
    @JsonProperty("VegetativeDate")
    val vegetativeDate: String? = null
)
