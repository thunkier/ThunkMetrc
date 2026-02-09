using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Adjustment
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustedByUserName")]
        public required string AdjustedByUserName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentDate")]
        public required string AdjustmentDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentNote")]
        public required string AdjustmentNote { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentQuantity")]
        public required double AdjustmentQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentReasonName")]
        public required string AdjustmentReasonName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentUnitOfMeasureAbbreviation")]
        public required string AdjustmentUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentUnitOfMeasureName")]
        public required string AdjustmentUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemCategoryName")]
        public required string ItemCategoryName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemName")]
        public required string ItemName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageId")]
        public required int PackageId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabTestResultExpirationDateTime")]
        public string? PackageLabTestResultExpirationDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public required string PackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackagedDate")]
        public required string PackagedDate { get; set; }
    }
}
