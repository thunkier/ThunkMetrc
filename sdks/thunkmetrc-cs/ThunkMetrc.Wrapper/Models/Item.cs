using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Item
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AdministrationMethod")]
        public string? AdministrationMethod { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Allergens")]
        public string? Allergens { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ApprovalStatus")]
        public required string ApprovalStatus { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ApprovalStatusDateTime")]
        public required string ApprovalStatusDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DefaultLabTestingState")]
        public required string DefaultLabTestingState { get; set; }
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
        public required string ProductCategoryType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductImages")]
        public required List<object> ProductImages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductPDFDocuments")]
        public required List<object> ProductPDFDocuments { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductPhotoDescription")]
        public string? ProductPhotoDescription { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PublicIngredients")]
        public string? PublicIngredients { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("QuantityType")]
        public required string QuantityType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ServingSize")]
        public string? ServingSize { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainId")]
        public required int StrainId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainName")]
        public required string StrainName { get; set; }
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
        public required string UnitOfMeasureName { get; set; }
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
}
