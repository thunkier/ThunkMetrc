using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ProcessingJobStartProcessingJobRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CountUnitOfMeasure")]
        public string? CountUnitOfMeasure { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("JobName")]
        public string? JobName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("JobType")]
        public string? JobType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public List<object>? Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StartDate")]
        public string? StartDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VolumeUnitOfMeasure")]
        public string? VolumeUnitOfMeasure { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WeightUnitOfMeasure")]
        public string? WeightUnitOfMeasure { get; set; }
        public class ProcessingJobStartProcessingJobRequestPackagesItem
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
