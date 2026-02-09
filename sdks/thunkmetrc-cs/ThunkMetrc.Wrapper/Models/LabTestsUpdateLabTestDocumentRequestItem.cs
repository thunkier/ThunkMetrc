using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class LabTestsUpdateLabTestDocumentRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DocumentFileBase64")]
        public string? DocumentFileBase64 { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DocumentFileName")]
        public string? DocumentFileName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultId")]
        public int? LabTestResultId { get; set; }
    }
}
