using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesCreateAdjustPlantBatchesRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentDate")]
        public string? AdjustmentDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentReason")]
        public string? AdjustmentReason { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchName")]
        public string? PlantBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
        public int? Quantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReasonNote")]
        public string? ReasonNote { get; set; }
    }
}
