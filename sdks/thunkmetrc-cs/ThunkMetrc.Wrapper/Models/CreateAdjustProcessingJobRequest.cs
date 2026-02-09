using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CreateAdjustProcessingJobRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentDate")]
        public string? AdjustmentDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentNote")]
        public string? AdjustmentNote { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentReason")]
        public string? AdjustmentReason { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CountUnitOfMeasureName")]
        public string? CountUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public List<object>? Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VolumeUnitOfMeasureName")]
        public string? VolumeUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WeightUnitOfMeasureName")]
        public string? WeightUnitOfMeasureName { get; set; }
        public class CreateAdjustProcessingJobRequestPackagesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
            public string? Label { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
            public int? Quantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasure")]
            public string? UnitOfMeasure { get; set; }
        }
    }
}
