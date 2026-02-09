package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackageSummary {
    @JsonProperty("ArchivedDate")
    public String archivedDate;
    @JsonProperty("ContainsDecontaminatedProduct")
    public Boolean containsDecontaminatedProduct;
    @JsonProperty("ContainsRemediatedProduct")
    public Boolean containsRemediatedProduct;
    @JsonProperty("DecontaminationDate")
    public String decontaminationDate;
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("ExternalId")
    public Integer externalId;
    @JsonProperty("FinishedDate")
    public String finishedDate;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("InitialLabTestingState")
    public String initialLabTestingState;
    @JsonProperty("IsDonation")
    public Boolean isDonation;
    @JsonProperty("IsDonationPersistent")
    public Boolean isDonationPersistent;
    @JsonProperty("IsFinished")
    public Boolean isFinished;
    @JsonProperty("IsOnHold")
    public Boolean isOnHold;
    @JsonProperty("IsOnHoldCombined")
    public Boolean isOnHoldCombined;
    @JsonProperty("IsOnInvestigation")
    public Boolean isOnInvestigation;
    @JsonProperty("IsOnInvestigationHold")
    public Boolean isOnInvestigationHold;
    @JsonProperty("IsOnInvestigationRecall")
    public Boolean isOnInvestigationRecall;
    @JsonProperty("IsOnRecall")
    public Boolean isOnRecall;
    @JsonProperty("IsOnRecallCombined")
    public Boolean isOnRecallCombined;
    @JsonProperty("IsOnRetailerDelivery")
    public Boolean isOnRetailerDelivery;
    @JsonProperty("IsProcessValidationTestingSample")
    public Boolean isProcessValidationTestingSample;
    @JsonProperty("IsProductionBatch")
    public Boolean isProductionBatch;
    @JsonProperty("IsTestingSample")
    public Boolean isTestingSample;
    @JsonProperty("IsTradeSample")
    public Boolean isTradeSample;
    @JsonProperty("IsTradeSamplePersistent")
    public Boolean isTradeSamplePersistent;
    @JsonProperty("Item")
    public PackagesPackageSummaryItem item;
    @JsonProperty("ItemFromFacilityLicenseNumber")
    public String itemFromFacilityLicenseNumber;
    @JsonProperty("ItemFromFacilityName")
    public String itemFromFacilityName;
    @JsonProperty("LabTestResultExpirationDateTime")
    public String labTestResultExpirationDateTime;
    @JsonProperty("LabTestStage")
    public String labTestStage;
    @JsonProperty("LabTestStageId")
    public Integer labTestStageId;
    @JsonProperty("LabTestingPerformedDate")
    public String labTestingPerformedDate;
    @JsonProperty("LabTestingRecordedDate")
    public String labTestingRecordedDate;
    @JsonProperty("LabTestingState")
    public String labTestingState;
    @JsonProperty("LabTestingStateDate")
    public String labTestingStateDate;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("LabelsLastGeneratedDateTime")
    public String labelsLastGeneratedDateTime;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("LocationId")
    public Integer locationId;
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("LocationTypeName")
    public String locationTypeName;
    @JsonProperty("Note")
    public String note;
    @JsonProperty("OriginalPackageQuantity")
    public Double originalPackageQuantity;
    @JsonProperty("PackageForProductDestruction")
    public String packageForProductDestruction;
    @JsonProperty("PackageType")
    public String packageType;
    @JsonProperty("PackagedDate")
    public String packagedDate;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("ProductLabel")
    public PackagesPackageSummaryProductLabel productLabel;
    @JsonProperty("ProductRequiresDecontamination")
    public Boolean productRequiresDecontamination;
    @JsonProperty("ProductRequiresRemediation")
    public Boolean productRequiresRemediation;
    @JsonProperty("ProductionBatchNumber")
    public String productionBatchNumber;
    @JsonProperty("Quantity")
    public Double quantity;
    @JsonProperty("ReceivedDateTime")
    public String receivedDateTime;
    @JsonProperty("ReceivedFromFacilityLicenseNumber")
    public String receivedFromFacilityLicenseNumber;
    @JsonProperty("ReceivedFromFacilityName")
    public String receivedFromFacilityName;
    @JsonProperty("ReceivedFromManifestNumber")
    public String receivedFromManifestNumber;
    @JsonProperty("RemediationDate")
    public String remediationDate;
    @JsonProperty("SellByDate")
    public String sellByDate;
    @JsonProperty("SourceHarvestCount")
    public Integer sourceHarvestCount;
    @JsonProperty("SourceHarvestNames")
    public String sourceHarvestNames;
    @JsonProperty("SourcePackageCount")
    public Integer sourcePackageCount;
    @JsonProperty("SourcePackageIsDonation")
    public Boolean sourcePackageIsDonation;
    @JsonProperty("SourcePackageIsTradeSample")
    public Boolean sourcePackageIsTradeSample;
    @JsonProperty("SourcePackageLabels")
    public String sourcePackageLabels;
    @JsonProperty("SourceProcessingJobCount")
    public Integer sourceProcessingJobCount;
    @JsonProperty("SourceProductionBatchNumbers")
    public String sourceProductionBatchNumbers;
    @JsonProperty("SublocationId")
    public Integer sublocationId;
    @JsonProperty("SublocationName")
    public String sublocationName;
    @JsonProperty("UnitOfMeasureAbbreviation")
    public String unitOfMeasureAbbreviation;
    @JsonProperty("UnitOfMeasureName")
    public String unitOfMeasureName;
    @JsonProperty("UseByDate")
    public String useByDate;
    public static class PackagesPackageSummaryItem {
    @JsonProperty("AdministrationMethod")
    public String administrationMethod;
    @JsonProperty("Allergens")
    public String allergens;
    @JsonProperty("ApprovalStatus")
    public Integer approvalStatus;
    @JsonProperty("ApprovalStatusDateTime")
    public String approvalStatusDateTime;
    @JsonProperty("DefaultLabTestingState")
    public Integer defaultLabTestingState;
    @JsonProperty("Description")
    public String description;
    @JsonProperty("GlobalProductName")
    public String globalProductName;
    @JsonProperty("GlobalProductNumber")
    public String globalProductNumber;
    @JsonProperty("HasExpirationDate")
    public Boolean hasExpirationDate;
    @JsonProperty("HasSellByDate")
    public Boolean hasSellByDate;
    @JsonProperty("HasUseByDate")
    public Boolean hasUseByDate;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IsExpirationDateRequired")
    public Boolean isExpirationDateRequired;
    @JsonProperty("IsSellByDateRequired")
    public Boolean isSellByDateRequired;
    @JsonProperty("IsUseByDateRequired")
    public Boolean isUseByDateRequired;
    @JsonProperty("IsUsed")
    public Boolean isUsed;
    @JsonProperty("ItemBrandId")
    public Integer itemBrandId;
    @JsonProperty("ItemBrandName")
    public String itemBrandName;
    @JsonProperty("LabTestBatchNames")
    public List<Object> labTestBatchNames;
    @JsonProperty("LabelImages")
    public List<Object> labelImages;
    @JsonProperty("LabelPhotoDescription")
    public String labelPhotoDescription;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("NumberOfDoses")
    public String numberOfDoses;
    @JsonProperty("PackagingImages")
    public List<Object> packagingImages;
    @JsonProperty("PackagingPhotoDescription")
    public String packagingPhotoDescription;
    @JsonProperty("ProcessingJobCategoryName")
    public String processingJobCategoryName;
    @JsonProperty("ProcessingJobTypeName")
    public String processingJobTypeName;
    @JsonProperty("ProductBrandName")
    public String productBrandName;
    @JsonProperty("ProductCategoryName")
    public String productCategoryName;
    @JsonProperty("ProductCategoryType")
    public Integer productCategoryType;
    @JsonProperty("ProductImages")
    public List<Object> productImages;
    @JsonProperty("ProductPDFDocuments")
    public List<Object> productPDFDocuments;
    @JsonProperty("ProductPhotoDescription")
    public String productPhotoDescription;
    @JsonProperty("PublicIngredients")
    public String publicIngredients;
    @JsonProperty("QuantityType")
    public Integer quantityType;
    @JsonProperty("ServingSize")
    public String servingSize;
    @JsonProperty("StrainId")
    public Integer strainId;
    @JsonProperty("StrainName")
    public String strainName;
    @JsonProperty("UnitCbdAContent")
    public Double unitCbdAContent;
    @JsonProperty("UnitCbdAContentDose")
    public Double unitCbdAContentDose;
    @JsonProperty("UnitCbdAContentDoseUnitOfMeasureName")
    public String unitCbdAContentDoseUnitOfMeasureName;
    @JsonProperty("UnitCbdAContentUnitOfMeasureName")
    public String unitCbdAContentUnitOfMeasureName;
    @JsonProperty("UnitCbdAPercent")
    public Double unitCbdAPercent;
    @JsonProperty("UnitCbdContent")
    public Double unitCbdContent;
    @JsonProperty("UnitCbdContentDose")
    public Double unitCbdContentDose;
    @JsonProperty("UnitCbdContentDoseUnitOfMeasureName")
    public String unitCbdContentDoseUnitOfMeasureName;
    @JsonProperty("UnitCbdContentUnitOfMeasureName")
    public String unitCbdContentUnitOfMeasureName;
    @JsonProperty("UnitCbdPercent")
    public Double unitCbdPercent;
    @JsonProperty("UnitOfMeasureName")
    public String unitOfMeasureName;
    @JsonProperty("UnitQuantity")
    public Double unitQuantity;
    @JsonProperty("UnitQuantityUnitOfMeasureName")
    public String unitQuantityUnitOfMeasureName;
    @JsonProperty("UnitThcAContent")
    public Double unitThcAContent;
    @JsonProperty("UnitThcAContentDose")
    public Double unitThcAContentDose;
    @JsonProperty("UnitThcAContentDoseUnitOfMeasureId")
    public Integer unitThcAContentDoseUnitOfMeasureId;
    @JsonProperty("UnitThcAContentDoseUnitOfMeasureName")
    public String unitThcAContentDoseUnitOfMeasureName;
    @JsonProperty("UnitThcAContentUnitOfMeasureName")
    public String unitThcAContentUnitOfMeasureName;
    @JsonProperty("UnitThcAPercent")
    public Double unitThcAPercent;
    @JsonProperty("UnitThcContent")
    public Double unitThcContent;
    @JsonProperty("UnitThcContentDose")
    public Double unitThcContentDose;
    @JsonProperty("UnitThcContentDoseUnitOfMeasureId")
    public Integer unitThcContentDoseUnitOfMeasureId;
    @JsonProperty("UnitThcContentDoseUnitOfMeasureName")
    public String unitThcContentDoseUnitOfMeasureName;
    @JsonProperty("UnitThcContentUnitOfMeasureName")
    public String unitThcContentUnitOfMeasureName;
    @JsonProperty("UnitThcPercent")
    public Double unitThcPercent;
    @JsonProperty("UnitVolume")
    public Double unitVolume;
    @JsonProperty("UnitVolumeUnitOfMeasureName")
    public String unitVolumeUnitOfMeasureName;
    @JsonProperty("UnitWeight")
    public Double unitWeight;
    @JsonProperty("UnitWeightUnitOfMeasureName")
    public String unitWeightUnitOfMeasureName;
}
    public static class PackagesPackageSummaryProductLabel {
    @JsonProperty("IsActive")
    public Boolean isActive;
    @JsonProperty("IsChildFromParentWithLabel")
    public Boolean isChildFromParentWithLabel;
    @JsonProperty("IsDiscontinued")
    public Boolean isDiscontinued;
    @JsonProperty("LastLabelGenerationDate")
    public String lastLabelGenerationDate;
    @JsonProperty("PackageId")
    public Integer packageId;
    @JsonProperty("QrCount")
    public Integer qrCount;
    @JsonProperty("TagId")
    public Integer tagId;
}
}
