using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UpdateAdjustPlantsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustCount")]
        public int? AdjustCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustReason")]
        public string? AdjustReason { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdjustmentDate")]
        public string? AdjustmentDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReasonNote")]
        public string? ReasonNote { get; set; }
    }
}
