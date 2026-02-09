using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantsUpdatePlantsTagRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewTag")]
        public string? NewTag { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReplaceDate")]
        public string? ReplaceDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagId")]
        public int? TagId { get; set; }
    }
}
