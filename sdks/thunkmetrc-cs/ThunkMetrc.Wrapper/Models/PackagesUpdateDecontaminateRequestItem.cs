using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PackagesUpdateDecontaminateRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DecontaminationDate")]
        public string? DecontaminationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DecontaminationMethodName")]
        public string? DecontaminationMethodName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DecontaminationSteps")]
        public string? DecontaminationSteps { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public string? PackageLabel { get; set; }
    }
}
