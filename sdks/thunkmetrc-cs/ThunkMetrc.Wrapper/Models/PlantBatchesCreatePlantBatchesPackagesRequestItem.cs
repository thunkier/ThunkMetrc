using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesCreatePlantBatchesPackagesRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Count")]
        public int? Count { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDate")]
        public string? ExpirationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsDonation")]
        public bool? IsDonation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsTradeSample")]
        public bool? IsTradeSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Item")]
        public string? Item { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Location")]
        public string? Location { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Note")]
        public string? Note { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatch")]
        public string? PlantBatch { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SellByDate")]
        public string? SellByDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Sublocation")]
        public string? Sublocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Tag")]
        public string? Tag { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UseByDate")]
        public string? UseByDate { get; set; }
    }
}
