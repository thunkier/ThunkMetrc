using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class LabTestsCreateRecordRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DocumentFileBase64")]
        public string? DocumentFileBase64 { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DocumentFileName")]
        public string? DocumentFileName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ResultDate")]
        public string? ResultDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Results")]
        public List<object>? Results { get; set; }
        public class LabTestsCreateRecordRequestResultsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("LabTestTypeName")]
            public string? LabTestTypeName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Notes")]
            public string? Notes { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Passed")]
            public bool? Passed { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
            public double? Quantity { get; set; }
        }
    }
}
