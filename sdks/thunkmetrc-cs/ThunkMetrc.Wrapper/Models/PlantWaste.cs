using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantWaste
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("GrowthPhase")]
        public required int GrowthPhase { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public required string Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationId")]
        public required int LocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationName")]
        public required string LocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantId")]
        public required int PlantId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantWasteId")]
        public required int PlantWasteId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StateName")]
        public required string StateName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationId")]
        public required int SublocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationName")]
        public string? SublocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteAmount")]
        public required int WasteAmount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteUnitOfMeasureAbbreviation")]
        public required string WasteUnitOfMeasureAbbreviation { get; set; }
    }
}
