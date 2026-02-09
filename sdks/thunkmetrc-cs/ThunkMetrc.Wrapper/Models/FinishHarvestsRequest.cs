using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class FinishHarvestsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
    }
}
