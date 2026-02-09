using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesCreateGrowthPhaseRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Count")]
        public int? Count { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrowthDate")]
        public string? GrowthDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrowthPhase")]
        public string? GrowthPhase { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewLocation")]
        public string? NewLocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewSublocation")]
        public string? NewSublocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StartingTag")]
        public string? StartingTag { get; set; }
    }
}
