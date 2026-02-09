using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class RetailIdCreateGenerateRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public string? PackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
        public double? Quantity { get; set; }
    }
}
