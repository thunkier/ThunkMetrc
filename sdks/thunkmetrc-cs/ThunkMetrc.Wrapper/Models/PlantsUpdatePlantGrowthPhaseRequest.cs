using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantsUpdatePlantGrowthPhaseRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("GrowthDate")]
        public string? GrowthDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrowthPhase")]
        public string? GrowthPhase { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewLocation")]
        public string? NewLocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewSublocation")]
        public string? NewSublocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewTag")]
        public string? NewTag { get; set; }
    }
}
