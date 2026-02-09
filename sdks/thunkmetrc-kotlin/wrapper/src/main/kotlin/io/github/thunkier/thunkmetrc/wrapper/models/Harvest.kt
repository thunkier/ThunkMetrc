package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Harvest(
    @JsonProperty("ArchivedDate")
    val archivedDate: String? = null,
    @JsonProperty("CurrentWeight")
    val currentWeight: Double? = null,
    @JsonProperty("DryingLocationId")
    val dryingLocationId: Int? = null,
    @JsonProperty("DryingLocationName")
    val dryingLocationName: String? = null,
    @JsonProperty("DryingLocationTypeName")
    val dryingLocationTypeName: String? = null,
    @JsonProperty("DryingSublocationId")
    val dryingSublocationId: Int? = null,
    @JsonProperty("DryingSublocationName")
    val dryingSublocationName: String? = null,
    @JsonProperty("FinishedDate")
    val finishedDate: String? = null,
    @JsonProperty("HarvestStartDate")
    val harvestStartDate: String? = null,
    @JsonProperty("HarvestType")
    val harvestType: String? = null,
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
    @JsonProperty("LabTestingState")
    val labTestingState: String? = null,
    @JsonProperty("LabTestingStateDate")
    val labTestingStateDate: String? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("PackageCount")
    val packageCount: Int? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PlantCount")
    val plantCount: Int? = null,
    @JsonProperty("SourceStrainCount")
    val sourceStrainCount: Int? = null,
    @JsonProperty("SourceStrainNames")
    val sourceStrainNames: String? = null,
    @JsonProperty("Strains")
    val strains: List<Any>? = null,
    @JsonProperty("TotalPackagedWeight")
    val totalPackagedWeight: Double? = null,
    @JsonProperty("TotalRestoredWeight")
    val totalRestoredWeight: Double? = null,
    @JsonProperty("TotalWasteWeight")
    val totalWasteWeight: Double? = null,
    @JsonProperty("TotalWetWeight")
    val totalWetWeight: Double? = null,
    @JsonProperty("UnitOfWeightName")
    val unitOfWeightName: String? = null
)
