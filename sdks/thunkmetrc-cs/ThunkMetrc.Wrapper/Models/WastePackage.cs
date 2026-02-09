using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class WastePackage
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ArchivedDate")]
        public string? ArchivedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsDecontaminatedProduct")]
        public required bool ContainsDecontaminatedProduct { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsRemediatedProduct")]
        public required bool ContainsRemediatedProduct { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DecontaminationDate")]
        public string? DecontaminationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDate")]
        public string? ExpirationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalId")]
        public int? ExternalId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishedDate")]
        public string? FinishedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("InitialLabTestingState")]
        public required string InitialLabTestingState { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsDonation")]
        public required bool IsDonation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsDonationPersistent")]
        public required bool IsDonationPersistent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsFinished")]
        public required bool IsFinished { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsFinishedGood")]
        public required bool IsFinishedGood { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnHold")]
        public required bool IsOnHold { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnHoldCombined")]
        public required bool IsOnHoldCombined { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnInvestigation")]
        public required bool IsOnInvestigation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnInvestigationHold")]
        public required bool IsOnInvestigationHold { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnInvestigationRecall")]
        public required bool IsOnInvestigationRecall { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnRecall")]
        public bool? IsOnRecall { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnRecallCombined")]
        public required bool IsOnRecallCombined { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnRetailerDelivery")]
        public required bool IsOnRetailerDelivery { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsProcessValidationTestingSample")]
        public required bool IsProcessValidationTestingSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsProductionBatch")]
        public required bool IsProductionBatch { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsTestingSample")]
        public required bool IsTestingSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsTradeSample")]
        public required bool IsTradeSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsTradeSamplePersistent")]
        public required bool IsTradeSamplePersistent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Item")]
        public required PlantsWastePackageItem Item { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemFromFacilityLicenseNumber")]
        public string? ItemFromFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemFromFacilityName")]
        public string? ItemFromFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultExpirationDateTime")]
        public string? LabTestResultExpirationDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestStage")]
        public string? LabTestStage { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestStageId")]
        public int? LabTestStageId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestingPerformedDate")]
        public string? LabTestingPerformedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestingRecordedDate")]
        public string? LabTestingRecordedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestingState")]
        public required string LabTestingState { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestingStateDate")]
        public required string LabTestingStateDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public required string Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabelsLastGeneratedDateTime")]
        public string? LabelsLastGeneratedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationId")]
        public int? LocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationName")]
        public string? LocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationTypeName")]
        public string? LocationTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Note")]
        public string? Note { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("OriginalPackageQuantity")]
        public required double OriginalPackageQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageForProductDestruction")]
        public string? PackageForProductDestruction { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageType")]
        public required string PackageType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackagedDate")]
        public required string PackagedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductLabel")]
        public required PlantsWastePackageProductLabel ProductLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductRequiresDecontamination")]
        public required bool ProductRequiresDecontamination { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductRequiresRemediation")]
        public required bool ProductRequiresRemediation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductionBatchNumber")]
        public string? ProductionBatchNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
        public required double Quantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedDateTime")]
        public string? ReceivedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedFromFacilityLicenseNumber")]
        public string? ReceivedFromFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedFromFacilityName")]
        public string? ReceivedFromFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedFromManifestNumber")]
        public string? ReceivedFromManifestNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediationDate")]
        public string? RemediationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SellByDate")]
        public string? SellByDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceHarvestCount")]
        public required int SourceHarvestCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceHarvestNames")]
        public string? SourceHarvestNames { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageCount")]
        public required int SourcePackageCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageIsDonation")]
        public required bool SourcePackageIsDonation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageIsTradeSample")]
        public required bool SourcePackageIsTradeSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageLabels")]
        public string? SourcePackageLabels { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceProcessingJobCount")]
        public required int SourceProcessingJobCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceProductionBatchNumbers")]
        public string? SourceProductionBatchNumbers { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationId")]
        public int? SublocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationName")]
        public string? SublocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasureAbbreviation")]
        public required string UnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasureName")]
        public required string UnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UseByDate")]
        public string? UseByDate { get; set; }
        public class PlantsWastePackageItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("AdministrationMethod")]
            public string? AdministrationMethod { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Allergens")]
            public string? Allergens { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ApprovalStatus")]
            public required int ApprovalStatus { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ApprovalStatusDateTime")]
            public required string ApprovalStatusDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DefaultLabTestingState")]
            public required int DefaultLabTestingState { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Description")]
            public string? Description { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("GlobalProductName")]
            public string? GlobalProductName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("GlobalProductNumber")]
            public string? GlobalProductNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("HasExpirationDate")]
            public required bool HasExpirationDate { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("HasSellByDate")]
            public required bool HasSellByDate { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("HasUseByDate")]
            public required bool HasUseByDate { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
            public required int Id { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsExpirationDateRequired")]
            public required bool IsExpirationDateRequired { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsSellByDateRequired")]
            public required bool IsSellByDateRequired { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsUseByDateRequired")]
            public required bool IsUseByDateRequired { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsUsed")]
            public required bool IsUsed { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemBrandId")]
            public required int ItemBrandId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemBrandName")]
            public string? ItemBrandName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LabTestBatchNames")]
            public required List<object> LabTestBatchNames { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LabelImages")]
            public required List<object> LabelImages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LabelPhotoDescription")]
            public string? LabelPhotoDescription { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
            public required string Name { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("NumberOfDoses")]
            public string? NumberOfDoses { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PackagingImages")]
            public required List<object> PackagingImages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PackagingPhotoDescription")]
            public string? PackagingPhotoDescription { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProcessingJobCategoryName")]
            public string? ProcessingJobCategoryName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProcessingJobTypeName")]
            public string? ProcessingJobTypeName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductBrandName")]
            public string? ProductBrandName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductCategoryName")]
            public required string ProductCategoryName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductCategoryType")]
            public required int ProductCategoryType { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductImages")]
            public required List<object> ProductImages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductPDFDocuments")]
            public required List<object> ProductPDFDocuments { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductPhotoDescription")]
            public string? ProductPhotoDescription { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PublicIngredients")]
            public string? PublicIngredients { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("QuantityType")]
            public required int QuantityType { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ServingSize")]
            public string? ServingSize { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("StrainId")]
            public int? StrainId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("StrainName")]
            public string? StrainName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdAContent")]
            public double? UnitCbdAContent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdAContentDose")]
            public double? UnitCbdAContentDose { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdAContentDoseUnitOfMeasureName")]
            public string? UnitCbdAContentDoseUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdAContentUnitOfMeasureName")]
            public string? UnitCbdAContentUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdAPercent")]
            public double? UnitCbdAPercent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdContent")]
            public double? UnitCbdContent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdContentDose")]
            public double? UnitCbdContentDose { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdContentDoseUnitOfMeasureName")]
            public string? UnitCbdContentDoseUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdContentUnitOfMeasureName")]
            public string? UnitCbdContentUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitCbdPercent")]
            public double? UnitCbdPercent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasureName")]
            public string? UnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitQuantity")]
            public double? UnitQuantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitQuantityUnitOfMeasureName")]
            public string? UnitQuantityUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcAContent")]
            public double? UnitThcAContent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcAContentDose")]
            public double? UnitThcAContentDose { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcAContentDoseUnitOfMeasureId")]
            public int? UnitThcAContentDoseUnitOfMeasureId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcAContentDoseUnitOfMeasureName")]
            public string? UnitThcAContentDoseUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcAContentUnitOfMeasureName")]
            public string? UnitThcAContentUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcAPercent")]
            public double? UnitThcAPercent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcContent")]
            public double? UnitThcContent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcContentDose")]
            public double? UnitThcContentDose { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcContentDoseUnitOfMeasureId")]
            public int? UnitThcContentDoseUnitOfMeasureId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcContentDoseUnitOfMeasureName")]
            public string? UnitThcContentDoseUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcContentUnitOfMeasureName")]
            public string? UnitThcContentUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcPercent")]
            public double? UnitThcPercent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitVolume")]
            public double? UnitVolume { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitVolumeUnitOfMeasureName")]
            public string? UnitVolumeUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitWeight")]
            public double? UnitWeight { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitWeightUnitOfMeasureName")]
            public string? UnitWeightUnitOfMeasureName { get; set; }
        }
        public class PlantsWastePackageProductLabel
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("IsActive")]
            public required bool IsActive { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsChildFromParentWithLabel")]
            public required bool IsChildFromParentWithLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LabelSource")]
            public string? LabelSource { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LastLabelGenerationDate")]
            public string? LastLabelGenerationDate { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("OriginalSourcePackageId")]
            public int? OriginalSourcePackageId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("OriginalSourcePackageLabel")]
            public string? OriginalSourcePackageLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("QrCodeRanges")]
            public string? QrCodeRanges { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("QrCount")]
            public required int QrCount { get; set; }
        }
    }
}
