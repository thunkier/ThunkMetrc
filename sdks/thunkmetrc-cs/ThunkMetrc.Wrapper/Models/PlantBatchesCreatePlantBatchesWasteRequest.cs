using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesCreatePlantBatchesWasteRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("MixedMaterial")]
        public string? MixedMaterial { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Note")]
        public string? Note { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchName")]
        public string? PlantBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReasonName")]
        public string? ReasonName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasureName")]
        public string? UnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteDate")]
        public string? WasteDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteMethodName")]
        public string? WasteMethodName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeight")]
        public double? WasteWeight { get; set; }
    }
}
