using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class RetailIdCreateAssociateRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public string? PackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("QrUrls")]
        public List<object>? QrUrls { get; set; }
    }
}
