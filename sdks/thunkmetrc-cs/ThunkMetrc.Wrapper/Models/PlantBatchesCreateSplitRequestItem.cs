using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesCreateSplitRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Count")]
        public int? Count { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GroupName")]
        public string? GroupName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Location")]
        public string? Location { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatch")]
        public string? PlantBatch { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Strain")]
        public string? Strain { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Sublocation")]
        public string? Sublocation { get; set; }
    }
}
