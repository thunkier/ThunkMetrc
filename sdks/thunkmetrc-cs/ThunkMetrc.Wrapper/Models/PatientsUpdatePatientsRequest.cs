using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PatientsUpdatePatientsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ConcentrateOuncesAllowed")]
        public int? ConcentrateOuncesAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FlowerOuncesAllowed")]
        public int? FlowerOuncesAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HasSalesLimitExemption")]
        public bool? HasSalesLimitExemption { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("InfusedOuncesAllowed")]
        public int? InfusedOuncesAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LicenseEffectiveEndDate")]
        public string? LicenseEffectiveEndDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LicenseEffectiveStartDate")]
        public string? LicenseEffectiveStartDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LicenseNumber")]
        public string? LicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("MaxConcentrateThcPercentAllowed")]
        public int? MaxConcentrateThcPercentAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("MaxFlowerThcPercentAllowed")]
        public int? MaxFlowerThcPercentAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewLicenseNumber")]
        public string? NewLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecommendedPlants")]
        public int? RecommendedPlants { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecommendedSmokableQuantity")]
        public int? RecommendedSmokableQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ThcOuncesAllowed")]
        public int? ThcOuncesAllowed { get; set; }
    }
}
