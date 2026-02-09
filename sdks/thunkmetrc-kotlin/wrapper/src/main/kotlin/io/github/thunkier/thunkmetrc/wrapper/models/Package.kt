package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Package(
    @JsonProperty("ArchivedDate")
    val archivedDate: String? = null,
    @JsonProperty("ContainsDecontaminatedProduct")
    val containsDecontaminatedProduct: Boolean? = null,
    @JsonProperty("ContainsRemediatedProduct")
    val containsRemediatedProduct: Boolean? = null,
    @JsonProperty("DecontaminationDate")
    val decontaminationDate: String? = null,
    @JsonProperty("ExpirationDate")
    val expirationDate: String? = null,
    @JsonProperty("ExternalId")
    val externalId: Int? = null,
    @JsonProperty("FinishedDate")
    val finishedDate: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("InitialLabTestingState")
    val initialLabTestingState: String? = null,
    @JsonProperty("IsDonation")
    val isDonation: Boolean? = null,
    @JsonProperty("IsDonationPersistent")
    val isDonationPersistent: Boolean? = null,
    @JsonProperty("IsFinished")
    val isFinished: Boolean? = null,
    @JsonProperty("IsOnHold")
    val isOnHold: Boolean? = null,
    @JsonProperty("IsOnHoldCombined")
    val isOnHoldCombined: Boolean? = null,
    @JsonProperty("IsOnInvestigation")
    val isOnInvestigation: Boolean? = null,
    @JsonProperty("IsOnInvestigationHold")
    val isOnInvestigationHold: Boolean? = null,
    @JsonProperty("IsOnInvestigationRecall")
    val isOnInvestigationRecall: Boolean? = null,
    @JsonProperty("IsOnRecall")
    val isOnRecall: Boolean? = null,
    @JsonProperty("IsOnRecallCombined")
    val isOnRecallCombined: Boolean? = null,
    @JsonProperty("IsOnRetailerDelivery")
    val isOnRetailerDelivery: Boolean? = null,
    @JsonProperty("IsProcessValidationTestingSample")
    val isProcessValidationTestingSample: Boolean? = null,
    @JsonProperty("IsProductionBatch")
    val isProductionBatch: Boolean? = null,
    @JsonProperty("IsTestingSample")
    val isTestingSample: Boolean? = null,
    @JsonProperty("IsTradeSample")
    val isTradeSample: Boolean? = null,
    @JsonProperty("IsTradeSamplePersistent")
    val isTradeSamplePersistent: Boolean? = null,
    @JsonProperty("Item")
    val item: PackageItem? = null,
    @JsonProperty("ItemFromFacilityLicenseNumber")
    val itemFromFacilityLicenseNumber: String? = null,
    @JsonProperty("ItemFromFacilityName")
    val itemFromFacilityName: String? = null,
    @JsonProperty("LabFacilityLicenseNumber")
    val labFacilityLicenseNumber: String? = null,
    @JsonProperty("LabFacilityName")
    val labFacilityName: String? = null,
    @JsonProperty("LabTestResultExpirationDateTime")
    val labTestResultExpirationDateTime: String? = null,
    @JsonProperty("LabTestStage")
    val labTestStage: String? = null,
    @JsonProperty("LabTestStageId")
    val labTestStageId: Int? = null,
    @JsonProperty("LabTestingPerformedDate")
    val labTestingPerformedDate: String? = null,
    @JsonProperty("LabTestingRecordedDate")
    val labTestingRecordedDate: String? = null,
    @JsonProperty("LabTestingState")
    val labTestingState: String? = null,
    @JsonProperty("LabTestingStateDate")
    val labTestingStateDate: String? = null,
    @JsonProperty("Label")
    val label: String? = null,
    @JsonProperty("LabelsLastGeneratedDateTime")
    val labelsLastGeneratedDateTime: String? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("LocationId")
    val locationId: Int? = null,
    @JsonProperty("LocationName")
    val locationName: String? = null,
    @JsonProperty("LocationTypeName")
    val locationTypeName: String? = null,
    @JsonProperty("Note")
    val note: String? = null,
    @JsonProperty("OriginalPackageQuantity")
    val originalPackageQuantity: Double? = null,
    @JsonProperty("OverallPassed")
    val overallPassed: Boolean? = null,
    @JsonProperty("PackageForProductDestruction")
    val packageForProductDestruction: String? = null,
    @JsonProperty("PackageType")
    val packageType: String? = null,
    @JsonProperty("PackagedDate")
    val packagedDate: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("ProductLabel")
    val productLabel: PackageProductLabel? = null,
    @JsonProperty("ProductRequiresDecontamination")
    val productRequiresDecontamination: Boolean? = null,
    @JsonProperty("ProductRequiresRemediation")
    val productRequiresRemediation: Boolean? = null,
    @JsonProperty("ProductionBatchNumber")
    val productionBatchNumber: String? = null,
    @JsonProperty("Quantity")
    val quantity: Double? = null,
    @JsonProperty("ReceivedDateTime")
    val receivedDateTime: String? = null,
    @JsonProperty("ReceivedFromFacilityLicenseNumber")
    val receivedFromFacilityLicenseNumber: String? = null,
    @JsonProperty("ReceivedFromFacilityName")
    val receivedFromFacilityName: String? = null,
    @JsonProperty("ReceivedFromManifestNumber")
    val receivedFromManifestNumber: String? = null,
    @JsonProperty("RemediationDate")
    val remediationDate: String? = null,
    @JsonProperty("ResultReleaseDateTime")
    val resultReleaseDateTime: String? = null,
    @JsonProperty("ResultReleased")
    val resultReleased: Boolean? = null,
    @JsonProperty("SellByDate")
    val sellByDate: String? = null,
    @JsonProperty("SourceHarvestCount")
    val sourceHarvestCount: Int? = null,
    @JsonProperty("SourceHarvestNames")
    val sourceHarvestNames: String? = null,
    @JsonProperty("SourcePackageCount")
    val sourcePackageCount: Int? = null,
    @JsonProperty("SourcePackageIsDonation")
    val sourcePackageIsDonation: Boolean? = null,
    @JsonProperty("SourcePackageIsTradeSample")
    val sourcePackageIsTradeSample: Boolean? = null,
    @JsonProperty("SourcePackageLabels")
    val sourcePackageLabels: String? = null,
    @JsonProperty("SourceProcessingJobCount")
    val sourceProcessingJobCount: Int? = null,
    @JsonProperty("SourceProductionBatchNumbers")
    val sourceProductionBatchNumbers: String? = null,
    @JsonProperty("SublocationId")
    val sublocationId: Int? = null,
    @JsonProperty("SublocationName")
    val sublocationName: String? = null,
    @JsonProperty("TestComment")
    val testComment: String? = null,
    @JsonProperty("TestInformationalOnly")
    val testInformationalOnly: Boolean? = null,
    @JsonProperty("TestPassed")
    val testPassed: Boolean? = null,
    @JsonProperty("TestPerformedDate")
    val testPerformedDate: String? = null,
    @JsonProperty("TestResultLevel")
    val testResultLevel: Int? = null,
    @JsonProperty("TestTypeName")
    val testTypeName: String? = null,
    @JsonProperty("UnitOfMeasureAbbreviation")
    val unitOfMeasureAbbreviation: String? = null,
    @JsonProperty("UnitOfMeasureName")
    val unitOfMeasureName: String? = null,
    @JsonProperty("UseByDate")
    val useByDate: String? = null
)
