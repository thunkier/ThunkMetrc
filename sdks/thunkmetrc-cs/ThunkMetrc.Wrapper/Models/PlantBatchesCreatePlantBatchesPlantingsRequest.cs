using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesCreatePlantBatchesPlantingsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Count")]
        public int? Count { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Location")]
        public string? Location { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePlantBatches")]
        public string? SourcePlantBatches { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Strain")]
        public string? Strain { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Sublocation")]
        public string? Sublocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Type")]
        public string? Type { get; set; }
    }
}
