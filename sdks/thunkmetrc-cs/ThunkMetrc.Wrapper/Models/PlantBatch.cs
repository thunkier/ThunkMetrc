using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatch
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DestroyedCount")]
        public required int DestroyedCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnHold")]
        public required bool IsOnHold { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnInvestigation")]
        public required bool IsOnInvestigation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnInvestigationHold")]
        public required bool IsOnInvestigationHold { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOnInvestigationRecall")]
        public required bool IsOnInvestigationRecall { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationId")]
        public int? LocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationName")]
        public string? LocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationTypeName")]
        public string? LocationTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("MultiPlantBatch")]
        public required bool MultiPlantBatch { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackagedCount")]
        public required int PackagedCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchTypeId")]
        public required int PlantBatchTypeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchTypeName")]
        public required string PlantBatchTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantedDate")]
        public required string PlantedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageId")]
        public int? SourcePackageId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageLabel")]
        public string? SourcePackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePlantBatchIds")]
        public required List<object> SourcePlantBatchIds { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePlantBatchNames")]
        public string? SourcePlantBatchNames { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePlantId")]
        public int? SourcePlantId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePlantLabel")]
        public string? SourcePlantLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainId")]
        public required int StrainId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainName")]
        public required string StrainName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationId")]
        public int? SublocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationName")]
        public string? SublocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TrackedCount")]
        public required int TrackedCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UntrackedCount")]
        public required int UntrackedCount { get; set; }
    }
}
