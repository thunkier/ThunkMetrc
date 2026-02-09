using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ProcessingJob
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CountUnitOfMeasureAbbreviation")]
        public string? CountUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CountUnitOfMeasureId")]
        public required int CountUnitOfMeasureId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CountUnitOfMeasureName")]
        public required string CountUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishNote")]
        public string? FinishNote { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishedDate")]
        public string? FinishedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsFinished")]
        public required bool IsFinished { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("JobTypeId")]
        public required int JobTypeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("JobTypeName")]
        public string? JobTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Number")]
        public required string Number { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public required List<object> Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StartDate")]
        public required string StartDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalCount")]
        public required int TotalCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalCountWaste")]
        public string? TotalCountWaste { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalQuantity")]
        public required double TotalQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalUnitOfMeasureId")]
        public required int TotalUnitOfMeasureId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalVolume")]
        public required double TotalVolume { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalVolumeWaste")]
        public string? TotalVolumeWaste { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalWeight")]
        public required double TotalWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalWeightWaste")]
        public string? TotalWeightWaste { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VolumeUnitOfMeasureAbbreviation")]
        public string? VolumeUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VolumeUnitOfMeasureId")]
        public required int VolumeUnitOfMeasureId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VolumeUnitOfMeasureName")]
        public required string VolumeUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteCountUnitOfMeasureAbbreviation")]
        public string? WasteCountUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteCountUnitOfMeasureId")]
        public int? WasteCountUnitOfMeasureId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteCountUnitOfMeasureName")]
        public string? WasteCountUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteVolumeUnitOfMeasureAbbreviation")]
        public string? WasteVolumeUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteVolumeUnitOfMeasureId")]
        public int? WasteVolumeUnitOfMeasureId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteVolumeUnitOfMeasureName")]
        public string? WasteVolumeUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeightUnitOfMeasureAbbreviation")]
        public string? WasteWeightUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeightUnitOfMeasureId")]
        public int? WasteWeightUnitOfMeasureId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeightUnitOfMeasureName")]
        public string? WasteWeightUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WeightUnitOfMeasureAbbreviation")]
        public string? WeightUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WeightUnitOfMeasureId")]
        public required int WeightUnitOfMeasureId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WeightUnitOfMeasureName")]
        public required string WeightUnitOfMeasureName { get; set; }
    }
}
