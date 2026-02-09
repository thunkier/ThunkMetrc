using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Patient
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("HasSalesLimitExemption")]
        public required bool HasSalesLimitExemption { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LicenseEffectiveEndDate")]
        public required string LicenseEffectiveEndDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LicenseEffectiveStartDate")]
        public required string LicenseEffectiveStartDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LicenseNumber")]
        public required string LicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("OtherFacilitiesCount")]
        public required int OtherFacilitiesCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientId")]
        public required int PatientId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecommendedPlants")]
        public required int RecommendedPlants { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecommendedSmokableQuantity")]
        public required double RecommendedSmokableQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RegistrationDate")]
        public required string RegistrationDate { get; set; }
    }
}
