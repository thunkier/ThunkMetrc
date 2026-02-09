using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class WriteResponse
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Ids")]
        public required List<int> Ids { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Warnings")]
        public string? Warnings { get; set; }
    }
}
