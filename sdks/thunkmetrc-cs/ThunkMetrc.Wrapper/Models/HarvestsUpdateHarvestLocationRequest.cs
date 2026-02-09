using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class HarvestsUpdateHarvestLocationRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingLocation")]
        public string? DryingLocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingSublocation")]
        public string? DryingSublocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestName")]
        public string? HarvestName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
    }
}
