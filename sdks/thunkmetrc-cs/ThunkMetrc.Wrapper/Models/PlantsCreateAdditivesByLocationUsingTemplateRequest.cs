using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantsCreateAdditivesByLocationUsingTemplateRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdditivesTemplateName")]
        public string? AdditivesTemplateName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationName")]
        public string? LocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Rate")]
        public string? Rate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationName")]
        public string? SublocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalAmountApplied")]
        public int? TotalAmountApplied { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalAmountUnitOfMeasure")]
        public string? TotalAmountUnitOfMeasure { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Volume")]
        public string? Volume { get; set; }
    }
}
