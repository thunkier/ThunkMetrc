using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ProcessingJobFinishProcessingJobRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishDate")]
        public string? FinishDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishNote")]
        public string? FinishNote { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalCountWaste")]
        public string? TotalCountWaste { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalVolumeWaste")]
        public string? TotalVolumeWaste { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalWeightWaste")]
        public int? TotalWeightWaste { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteCountUnitOfMeasureName")]
        public string? WasteCountUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteVolumeUnitOfMeasureName")]
        public string? WasteVolumeUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeightUnitOfMeasureName")]
        public string? WasteWeightUnitOfMeasureName { get; set; }
    }
}
