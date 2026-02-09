using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UpdateMergeRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("MergeDate")]
        public string? MergeDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePlantGroupLabel")]
        public string? SourcePlantGroupLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TargetPlantGroupLabel")]
        public string? TargetPlantGroupLabel { get; set; }
    }
}
