using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UpdatePlantsStrainRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainId")]
        public int? StrainId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainName")]
        public string? StrainName { get; set; }
    }
}
