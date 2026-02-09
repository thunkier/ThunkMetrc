using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantsCreatePlantsPlantingsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationName")]
        public string? LocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchName")]
        public string? PlantBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchType")]
        public string? PlantBatchType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantCount")]
        public int? PlantCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantLabel")]
        public string? PlantLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainName")]
        public string? StrainName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationName")]
        public string? SublocationName { get; set; }
    }
}
