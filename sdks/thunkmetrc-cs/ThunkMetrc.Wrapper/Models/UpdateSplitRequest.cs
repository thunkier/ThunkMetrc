using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UpdateSplitRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantCount")]
        public int? PlantCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePlantLabel")]
        public string? SourcePlantLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SplitDate")]
        public string? SplitDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainLabel")]
        public string? StrainLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagLabel")]
        public string? TagLabel { get; set; }
    }
}
