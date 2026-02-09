using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class FinishPackagesRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
    }
}
