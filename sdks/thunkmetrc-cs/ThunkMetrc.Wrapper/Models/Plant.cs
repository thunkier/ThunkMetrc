using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Plant
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ClonedCount")]
        public int? ClonedCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DescendedCount")]
        public int? DescendedCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DestroyedByUserName")]
        public string? DestroyedByUserName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DestroyedDate")]
        public string? DestroyedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DestroyedNote")]
        public string? DestroyedNote { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FloweringDate")]
        public required string FloweringDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GroupTagTypeMax")]
        public required int GroupTagTypeMax { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrowthPhase")]
        public required string GrowthPhase { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestCount")]
        public required int HarvestCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestId")]
        public int? HarvestId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestedDate")]
        public string? HarvestedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestedUnitOfWeightAbbreviation")]
        public string? HarvestedUnitOfWeightAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestedUnitOfWeightName")]
        public string? HarvestedUnitOfWeightName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestedWetWeight")]
        public double? HarvestedWetWeight { get; set; }
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
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public required string Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationId")]
        public required int LocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationName")]
        public required string LocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationTypeName")]
        public string? LocationTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("MotherPlantDate")]
        public string? MotherPlantDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchId")]
        public required int PlantBatchId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchName")]
        public required string PlantBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchTypeId")]
        public required int PlantBatchTypeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchTypeName")]
        public required string PlantBatchTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantedDate")]
        public required string PlantedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("State")]
        public required string State { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainId")]
        public required int StrainId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainName")]
        public required string StrainName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationId")]
        public int? SublocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationName")]
        public string? SublocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SurvivedCount")]
        public int? SurvivedCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagTypeMax")]
        public required int TagTypeMax { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VegetativeDate")]
        public required string VegetativeDate { get; set; }
    }
}
