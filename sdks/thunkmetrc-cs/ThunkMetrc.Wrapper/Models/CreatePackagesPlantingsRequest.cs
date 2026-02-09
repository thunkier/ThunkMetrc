using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CreatePackagesPlantingsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationName")]
        public string? LocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageAdjustmentAmount")]
        public int? PackageAdjustmentAmount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageAdjustmentUnitOfMeasureName")]
        public string? PackageAdjustmentUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public string? PackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchName")]
        public string? PlantBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchType")]
        public string? PlantBatchType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantCount")]
        public int? PlantCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantedDate")]
        public string? PlantedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainName")]
        public string? StrainName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationName")]
        public string? SublocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnpackagedDate")]
        public string? UnpackagedDate { get; set; }
    }
}
