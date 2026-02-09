using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class LabTestBatche
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemCategories")]
        public required List<LabTestsLabTestBatcheItemCategoriesItem> ItemCategories { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemCategoryCount")]
        public required int ItemCategoryCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestTypeCount")]
        public required int LabTestTypeCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestTypes")]
        public required List<LabTestsLabTestBatcheLabTestTypesItem> LabTestTypes { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresAllFromLabTestBatch")]
        public required bool RequiresAllFromLabTestBatch { get; set; }
        public class LabTestsLabTestBatcheItemCategoriesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductCategory")]
            public required LabTestsLabTestBatcheItemCategoriesItemProductCategory ProductCategory { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductCategoryId")]
            public required string ProductCategoryId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RequireLabTestBatch")]
            public required bool RequireLabTestBatch { get; set; }
            public class LabTestsLabTestBatcheItemCategoriesItemProductCategory
            {
                [global::System.Text.Json.Serialization.JsonPropertyName("CanBeDecontaminated")]
                public required bool CanBeDecontaminated { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("CanBeDestroyed")]
                public required bool CanBeDestroyed { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("CanBeRemediated")]
                public required bool CanBeRemediated { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("CanContainSeeds")]
                public required bool CanContainSeeds { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("LabTestBatchNames")]
                public required List<object> LabTestBatchNames { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
                public required string Name { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("ProductCategoryType")]
                public required int ProductCategoryType { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("QuantityType")]
                public required int QuantityType { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresAdministrationMethod")]
                public required bool RequiresAdministrationMethod { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresAllergens")]
                public required bool RequiresAllergens { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresDescription")]
                public required bool RequiresDescription { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresItemBrand")]
                public required bool RequiresItemBrand { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresLabelPhotoDescription")]
                public required bool RequiresLabelPhotoDescription { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresLabelPhotos")]
                public required int RequiresLabelPhotos { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresNumberOfDoses")]
                public required bool RequiresNumberOfDoses { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresPackagingPhotoDescription")]
                public required bool RequiresPackagingPhotoDescription { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresPackagingPhotos")]
                public required int RequiresPackagingPhotos { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresProductPDFDocuments")]
                public required int RequiresProductPDFDocuments { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresProductPhotoDescription")]
                public required bool RequiresProductPhotoDescription { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresProductPhotos")]
                public required int RequiresProductPhotos { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresPublicIngredients")]
                public required bool RequiresPublicIngredients { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresServingSize")]
                public required bool RequiresServingSize { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresStrain")]
                public required bool RequiresStrain { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresSupplyDurationDays")]
                public required bool RequiresSupplyDurationDays { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitCbdAContent")]
                public required bool RequiresUnitCbdAContent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitCbdAContentDose")]
                public required bool RequiresUnitCbdAContentDose { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitCbdAPercent")]
                public required bool RequiresUnitCbdAPercent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitCbdContent")]
                public required bool RequiresUnitCbdContent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitCbdContentDose")]
                public required bool RequiresUnitCbdContentDose { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitCbdPercent")]
                public required bool RequiresUnitCbdPercent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitThcAContent")]
                public required bool RequiresUnitThcAContent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitThcAContentDose")]
                public required bool RequiresUnitThcAContentDose { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitThcAPercent")]
                public required bool RequiresUnitThcAPercent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitThcContent")]
                public required bool RequiresUnitThcContent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitThcContentDose")]
                public required bool RequiresUnitThcContentDose { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitThcPercent")]
                public required bool RequiresUnitThcPercent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitVolume")]
                public required bool RequiresUnitVolume { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("RequiresUnitWeight")]
                public required bool RequiresUnitWeight { get; set; }
            }
        }
        public class LabTestsLabTestBatcheLabTestTypesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("AlwaysPasses")]
            public required bool AlwaysPasses { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DependencyMode")]
            public required int DependencyMode { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
            public required int Id { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("InformationalOnly")]
            public required bool InformationalOnly { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultExpirationDays")]
            public int? LabTestResultExpirationDays { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultMaximum")]
            public string? LabTestResultMaximum { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultMinimum")]
            public string? LabTestResultMinimum { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultMode")]
            public required int LabTestResultMode { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("MaxAllowedFailureCount")]
            public required int MaxAllowedFailureCount { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
            public required string Name { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RequiresTestResult")]
            public required bool RequiresTestResult { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ResearchAndDevelopment")]
            public required bool ResearchAndDevelopment { get; set; }
        }
    }
}
