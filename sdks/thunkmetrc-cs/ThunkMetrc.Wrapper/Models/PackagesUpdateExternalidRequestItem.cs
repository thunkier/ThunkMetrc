using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PackagesUpdateExternalidRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalId")]
        public string? ExternalId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public string? PackageLabel { get; set; }
    }
}
