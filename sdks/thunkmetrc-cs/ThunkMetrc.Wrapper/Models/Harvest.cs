using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Harvest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ArchivedDate")]
        public string? ArchivedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CurrentWeight")]
        public required double CurrentWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingLocationId")]
        public required int DryingLocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingLocationName")]
        public required string DryingLocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingLocationTypeName")]
        public string? DryingLocationTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingSublocationId")]
        public required int DryingSublocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingSublocationName")]
        public string? DryingSublocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishedDate")]
        public string? FinishedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestStartDate")]
        public required string HarvestStartDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestType")]
        public required string HarvestType { get; set; }
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
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestingState")]
        public string? LabTestingState { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestingStateDate")]
        public string? LabTestingStateDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageCount")]
        public required int PackageCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantCount")]
        public required int PlantCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceStrainCount")]
        public required int SourceStrainCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceStrainNames")]
        public string? SourceStrainNames { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Strains")]
        public required List<object> Strains { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPackagedWeight")]
        public required double TotalPackagedWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalRestoredWeight")]
        public required double TotalRestoredWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalWasteWeight")]
        public required double TotalWasteWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalWetWeight")]
        public required double TotalWetWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfWeightName")]
        public required string UnitOfWeightName { get; set; }
    }
}
