using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PackagesUpdateRemediateRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public string? PackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediationDate")]
        public string? RemediationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediationMethodName")]
        public string? RemediationMethodName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediationSteps")]
        public string? RemediationSteps { get; set; }
    }
}
