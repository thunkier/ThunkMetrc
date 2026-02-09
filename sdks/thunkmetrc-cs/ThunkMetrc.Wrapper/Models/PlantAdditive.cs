using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantAdditive
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AdditiveTypeName")]
        public string? AdditiveTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AmountUnitOfMeasure")]
        public required string AmountUnitOfMeasure { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ApplicationDevice")]
        public required string ApplicationDevice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EpaRegistrationNumber")]
        public string? EpaRegistrationNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Note")]
        public string? Note { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchId")]
        public int? PlantBatchId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchName")]
        public string? PlantBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantCount")]
        public required int PlantCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductSupplier")]
        public required string ProductSupplier { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductTradeName")]
        public required string ProductTradeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Rate")]
        public string? Rate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RestrictiveEntryIntervalQuantityDescription")]
        public string? RestrictiveEntryIntervalQuantityDescription { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RestrictiveEntryIntervalTimeDescription")]
        public string? RestrictiveEntryIntervalTimeDescription { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalAmountApplied")]
        public required int TotalAmountApplied { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Volume")]
        public double? Volume { get; set; }
    }
}
