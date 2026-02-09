using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransferTemplateDeliveryPackage
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsRemediatedProduct")]
        public required bool ContainsRemediatedProduct { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDate")]
        public string? ExpirationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalId")]
        public int? ExternalId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrossUnitOfWeightName")]
        public string? GrossUnitOfWeightName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsDonation")]
        public required bool IsDonation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsTestingSample")]
        public required bool IsTestingSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsTradeSample")]
        public required bool IsTradeSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsTradeSamplePersistent")]
        public required bool IsTradeSamplePersistent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemBrandName")]
        public string? ItemBrandName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemCategoryName")]
        public required string ItemCategoryName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemId")]
        public required int ItemId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemName")]
        public required string ItemName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemServingSize")]
        public string? ItemServingSize { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemStrainName")]
        public string? ItemStrainName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemSupplyDurationDays")]
        public int? ItemSupplyDurationDays { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdAContent")]
        public double? ItemUnitCbdAContent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdAContentDose")]
        public double? ItemUnitCbdAContentDose { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdAContentDoseUnitOfMeasureName")]
        public string? ItemUnitCbdAContentDoseUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdAContentUnitOfMeasureName")]
        public string? ItemUnitCbdAContentUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdAPercent")]
        public double? ItemUnitCbdAPercent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdContent")]
        public double? ItemUnitCbdContent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdContentDose")]
        public double? ItemUnitCbdContentDose { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdContentDoseUnitOfMeasureName")]
        public string? ItemUnitCbdContentDoseUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdContentUnitOfMeasureName")]
        public string? ItemUnitCbdContentUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdPercent")]
        public double? ItemUnitCbdPercent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitQuantity")]
        public double? ItemUnitQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitQuantityUnitOfMeasureName")]
        public string? ItemUnitQuantityUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcAContent")]
        public double? ItemUnitThcAContent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcAContentDose")]
        public double? ItemUnitThcAContentDose { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcAContentDoseUnitOfMeasureName")]
        public string? ItemUnitThcAContentDoseUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcAContentUnitOfMeasureName")]
        public string? ItemUnitThcAContentUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcAPercent")]
        public double? ItemUnitThcAPercent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcContent")]
        public double? ItemUnitThcContent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcContentDose")]
        public double? ItemUnitThcContentDose { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcContentDoseUnitOfMeasureName")]
        public string? ItemUnitThcContentDoseUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcContentUnitOfMeasureName")]
        public string? ItemUnitThcContentUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcPercent")]
        public double? ItemUnitThcPercent { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitVolume")]
        public double? ItemUnitVolume { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitVolumeUnitOfMeasureName")]
        public string? ItemUnitVolumeUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitWeight")]
        public double? ItemUnitWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitWeightUnitOfMeasureName")]
        public string? ItemUnitWeightUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestingState")]
        public required string LabTestingState { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageId")]
        public required int PackageId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public required string PackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageType")]
        public required string PackageType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackagedDate")]
        public string? PackagedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductLabel")]
        public required TransfersTransferTemplateDeliveryPackageProductLabel ProductLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductRequiresRemediation")]
        public required bool ProductRequiresRemediation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductionBatchNumber")]
        public string? ProductionBatchNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedQuantity")]
        public required double ReceivedQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedUnitOfMeasureName")]
        public required string ReceivedUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediationDate")]
        public string? RemediationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SellByDate")]
        public string? SellByDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentPackageState")]
        public required string ShipmentPackageState { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShippedQuantity")]
        public required double ShippedQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShippedUnitOfMeasureName")]
        public required string ShippedUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceHarvestNames")]
        public string? SourceHarvestNames { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageIsDonation")]
        public required bool SourcePackageIsDonation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageIsTradeSample")]
        public required bool SourcePackageIsTradeSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageLabels")]
        public string? SourcePackageLabels { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceProductionBatchNumbers")]
        public string? SourceProductionBatchNumbers { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UseByDate")]
        public string? UseByDate { get; set; }
        public class TransfersTransferTemplateDeliveryPackageProductLabel
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("IsActive")]
            public required bool IsActive { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsChildFromParentWithLabel")]
            public required bool IsChildFromParentWithLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsDiscontinued")]
            public required bool IsDiscontinued { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LastLabelGenerationDate")]
            public string? LastLabelGenerationDate { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PackageId")]
            public int? PackageId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("QrCount")]
            public required int QrCount { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TagId")]
            public int? TagId { get; set; }
        }
    }
}
