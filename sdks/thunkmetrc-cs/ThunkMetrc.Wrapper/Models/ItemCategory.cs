using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ItemCategory
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
        public required string ProductCategoryType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("QuantityType")]
        public required string QuantityType { get; set; }
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
